package svc

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/cockroachdb/errors"
)

// ExecuteSyncPlan 执行覆盖同步；已有同步任务运行时拒绝新的调用。
func (t *VaultService) ExecuteSyncPlan(plan SyncPlan) (SyncResult, error) {
	if !t.isSyncing.CompareAndSwap(false, true) {
		return SyncResult{}, errors.New("同步正在执行，请等待当前任务完成")
	}
	defer t.isSyncing.Store(false)

	mainVaultPath, targetVaultPaths, err := precheckSyncPlan(plan)
	if err != nil {
		return SyncResult{}, errors.WithStack(err)
	}

	result := SyncResult{Targets: make([]TargetSyncResult, 0, len(plan.Targets))}
	for i, target := range plan.Targets {
		targetResult := TargetSyncResult{
			VaultPath: targetVaultPaths[i],
			Created:   make([]string, 0, len(target.Items)),
			Overwrote: make([]string, 0, len(target.Items)),
			Errors:    make([]string, 0),
		}

		for _, item := range target.Items {
			src := filepath.Join(mainVaultPath, ".obsidian", filepath.FromSlash(item.Path))
			dst := filepath.Join(targetVaultPaths[i], ".obsidian", filepath.FromSlash(item.Path))
			if err := copySyncItem(src, dst); err != nil {
				targetResult.Errors = append(targetResult.Errors, errors.Wrapf(err, "同步配置失败: %s", item.Path).Error())
				continue
			}
			if item.Action == SyncPlanActionCreate {
				targetResult.Created = append(targetResult.Created, item.Path)
			} else {
				targetResult.Overwrote = append(targetResult.Overwrote, item.Path)
			}
		}
		result.Targets = append(result.Targets, targetResult)
	}
	return result, nil
}

// precheckSyncPlan 校验执行计划并返回规范化的 vault 路径。
func precheckSyncPlan(plan SyncPlan) (string, []string, error) {
	mainVaultPath, err := precheckVaultPath(plan.MainVaultPath)
	if err != nil {
		return "", nil, err
	}

	targetVaultPaths := make([]string, 0, len(plan.Targets))
	for _, target := range plan.Targets {
		targetVaultPath, err := precheckVaultPath(target.VaultPath)
		if err != nil {
			return "", nil, err
		}
		if samePath(mainVaultPath, targetVaultPath) {
			return "", nil, errors.Errorf("主库不能作为从库: %s", targetVaultPath)
		}
		for _, item := range target.Items {
			if item.Action != SyncPlanActionCreate && item.Action != SyncPlanActionOverwrite {
				return "", nil, errors.Errorf("无效的同步动作: %s", item.Action)
			}
			if err := precheckSyncPath(item.Path); err != nil {
				return "", nil, err
			}
		}
		targetVaultPaths = append(targetVaultPaths, targetVaultPath)
	}
	return mainVaultPath, targetVaultPaths, nil
}

// precheckSyncPath 检查同步路径是否为 .obsidian 下的相对路径。
func precheckSyncPath(path string) error {
	path = filepath.Clean(filepath.FromSlash(normalizeSyncPath(path)))
	if path == "." || filepath.IsAbs(path) || path == ".." || strings.HasPrefix(path, ".."+string(filepath.Separator)) {
		return errors.Errorf("无效的同步配置路径: %s", path)
	}
	return nil
}

// copySyncItem 将单个文件或目录覆盖复制到目标路径。
func copySyncItem(src string, dst string) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return copySyncFile(src, dst, info.Mode())
	}

	return filepath.WalkDir(src, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		targetPath := filepath.Join(dst, relPath)
		if entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				return err
			}
			return os.MkdirAll(targetPath, info.Mode())
		}
		info, err := entry.Info()
		if err != nil {
			return err
		}
		return copySyncFile(path, targetPath, info.Mode())
	})
}

// copySyncFile 覆盖复制文件并保留权限位。
func copySyncFile(src string, dst string, mode fs.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}

	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.OpenFile(dst, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, mode)
	if err != nil {
		return err
	}
	if _, err := io.Copy(dstFile, srcFile); err != nil {
		dstFile.Close()
		return err
	}
	return dstFile.Close()
}
