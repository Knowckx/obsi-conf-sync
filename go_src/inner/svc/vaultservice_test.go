package svc_test

import (
	"testing"

	"obsi-conf-sync/go_src/inner/svc"
)

func Test_ScanVaults(t *testing.T) {
	testPath := `D:\AsukaFiles\AsuObsidianStore`

	service := &svc.VaultService{}
	result, err := service.ScanVaults(testPath)
	if err != nil {
		t.Fatalf("ScanVaults 失败: %v", err)
	}
	t.Logf("result:\n%+v", result)
}
