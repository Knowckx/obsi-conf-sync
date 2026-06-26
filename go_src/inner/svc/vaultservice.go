package svc

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"obsi-conf-sync/go_src/conf"

	"github.com/cockroachdb/errors"
)

// VaultService 向前端暴露 vault 扫描接口。
type VaultService struct{}

// VaultInfo 表示扫描发现的 Obsidian vault。
type VaultInfo struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

// String 返回 vault 的简短展示文本。
func (t VaultInfo) String() string {
	out := fmt.Sprintf("Name:%s", t.Name)
	return out
}

// ScanVaults 扫描 root 并返回包含 .obsidian/ 的目录。
func (s *VaultService) ScanVaults(root string) ([]VaultInfo, error) {
	rootPath, err := precheckScanRoot(root)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	vaults, err := scanVaultRoot(rootPath)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return vaults, nil
}

// precheckScanRoot 检查扫描根目录并返回绝对路径。
func precheckScanRoot(root string) (string, error) {
	if strings.TrimSpace(root) == "" {
		return "", errors.New("root 不能为空")
	}

	rootPath, err := filepath.Abs(root)
	if err != nil {
		return "", err
	}

	info, err := os.Stat(rootPath)
	if err != nil {
		return "", err
	}
	if !info.IsDir() {
		return "", errors.Errorf("root 不是目录: %s", rootPath)
	}
	return rootPath, nil
}

// scanVaultRoot 扫描 root 本身或 root 的最多 2 层子目录。
func scanVaultRoot(rootPath string) ([]VaultInfo, error) {
	isVault, err := isVaultDir(rootPath)
	if err != nil {
		return nil, err
	}
	if isVault {
		return []VaultInfo{{
			Path: rootPath,
			Name: filepath.Base(rootPath),
		}}, nil
	}

	vaults := make([]VaultInfo, 0)
	err = scanVaultChildren(rootPath, 2, &vaults)
	if err != nil {
		return nil, err
	}
	return vaults, nil
}

// scanVaultChildren 扫描子目录，remainDepth 表示还允许向下检查的层数。
func scanVaultChildren(parentPath string, remainDepth int, vaults *[]VaultInfo) error {
	if remainDepth <= 0 {
		return nil
	}

	entries, err := os.ReadDir(parentPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		if shouldSkipScanDir(entry.Name()) {
			continue
		}

		childPath := filepath.Join(parentPath, entry.Name())
		isVault, err := isVaultDir(childPath)
		if err != nil {
			return err
		}
		if isVault {
			*vaults = append(*vaults, VaultInfo{
				Path: childPath,
				Name: entry.Name(),
			})
			continue
		}

		if err := scanVaultChildren(childPath, remainDepth-1, vaults); err != nil {
			return err
		}
	}
	return nil
}

// isVaultDir 判断目录是否为 Obsidian vault。
func isVaultDir(dir string) (bool, error) {
	obsidianDir := filepath.Join(dir, ".obsidian")
	info, err := os.Stat(obsidianDir)
	if err == nil {
		return info.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// shouldSkipScanDir 判断目录是否应跳过扫描。
func shouldSkipScanDir(name string) bool {
	if _, ok := conf.SkipVaultScanDirNames[name]; ok {
		return true
	}
	// 目录名以 . 开头，就认为是隐藏目录，跳过扫描
	return strings.HasPrefix(name, ".")
}
