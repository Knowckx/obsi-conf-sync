package svc

import (
	"os"
	"path/filepath"

	"github.com/cockroachdb/errors"
)

// DevService 向前端暴露开发预设接口。
type DevService struct{}

// ResetTestCases 重建开发预设使用的测试目录。
func (s *DevService) ResetTestCases() error {
	sourcePath, err := precheckScanRoot("test_cases")
	if err != nil {
		return errors.WithStack(err)
	}
	targetPath, err := absoluteDirectoryPath("temp/test_cases_1")
	if err != nil {
		return errors.WithStack(err)
	}
	if err := os.RemoveAll(targetPath); err != nil {
		return errors.WithStack(err)
	}
	if err := copyDirectory(sourcePath, targetPath); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// RemoveDirectory 删除开发预设生成的测试目录。
func (s *DevService) RemoveDirectory() error {
	path, err := absoluteDirectoryPath("temp/test_cases_1")
	if err != nil {
		return errors.WithStack(err)
	}
	if err := os.RemoveAll(path); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// copyDirectory 复制目录内容并保留文件、目录权限。
func copyDirectory(sourcePath string, targetPath string) error {
	return filepath.WalkDir(sourcePath, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		relPath, err := filepath.Rel(sourcePath, path)
		if err != nil {
			return err
		}
		destinationPath := filepath.Join(targetPath, relPath)
		if entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				return err
			}
			return os.MkdirAll(destinationPath, info.Mode())
		}
		info, err := entry.Info()
		if err != nil {
			return err
		}
		return copySyncFile(path, destinationPath, info.Mode())
	})
}
