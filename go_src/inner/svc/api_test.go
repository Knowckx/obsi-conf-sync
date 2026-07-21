package svc_test

import (
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"testing"

	"obsi-conf-sync/go_src/inner/svc"
)

var coreSyncSelectedPaths = []string{
	"app.json",
	"community-plugins.json",
	"snippets/",
	"themes/",
	"plugins/open-in-new-tab/",
	"plugins/open-tab-settings/",
}

// Test_CoreSyncRegression 验证同步 MVP 的核心后端流程和固定预期结果。
func Test_CoreSyncRegression(t *testing.T) {
	fixturesRoot := coreTestCasesRoot(t)
	runRoot := t.TempDir()
	for _, name := range []string{"vault1", "vault2", "vault3"} {
		copyTree(t, filepath.Join(fixturesRoot, name), filepath.Join(runRoot, name))
	}

	mainVault := filepath.Join(runRoot, "vault1")
	targetVault2 := filepath.Join(runRoot, "vault2")
	targetVault3 := filepath.Join(runRoot, "vault3")
	service := &svc.VaultService{}

	items, err := service.ListConfigItems(mainVault)
	if err != nil {
		t.Fatalf("ListConfigItems 失败: %v", err)
	}
	assertContainsPaths(t, configItemPaths(items), coreSyncSelectedPaths)

	plan, err := service.BuildSyncPlan(svc.SyncRequest{
		MainVaultPath:    mainVault,
		TargetVaultPaths: []string{targetVault2, targetVault3},
		SelectedPaths:    coreSyncSelectedPaths,
	})
	if err != nil {
		t.Fatalf("BuildSyncPlan 失败: %v", err)
	}
	if len(plan.Targets) != 2 {
		t.Fatalf("目标库数量不符: want 2, got %d", len(plan.Targets))
	}

	assertTargetPlan(t, plan.Targets[0], []svc.SyncPlanAction{
		svc.SyncPlanActionOverwrite,
		svc.SyncPlanActionCreate,
		svc.SyncPlanActionCreate,
		svc.SyncPlanActionCreate,
		svc.SyncPlanActionCreate,
		svc.SyncPlanActionCreate,
	})
	assertTargetPlan(t, plan.Targets[1], []svc.SyncPlanAction{
		svc.SyncPlanActionOverwrite,
		svc.SyncPlanActionCreate,
		svc.SyncPlanActionOverwrite,
		svc.SyncPlanActionOverwrite,
		svc.SyncPlanActionOverwrite,
		svc.SyncPlanActionCreate,
	})

	result, err := service.ExecuteSyncPlan(plan)
	if err != nil {
		t.Fatalf("ExecuteSyncPlan 失败: %v", err)
	}
	if len(result.Targets) != 2 {
		t.Fatalf("执行结果目标库数量不符: want 2, got %d", len(result.Targets))
	}

	vault2Result := result.Targets[0]
	assertPaths(t, "vault2.created", vault2Result.Created, []string{
		"community-plugins.json", "snippets/", "themes/",
		"plugins/open-in-new-tab/", "plugins/open-tab-settings/",
	})
	assertPaths(t, "vault2.overwrote", vault2Result.Overwrote, []string{"app.json"})
	assertPaths(t, "vault2.errors", vault2Result.Errors, nil)

	vault3Result := result.Targets[1]
	assertPaths(t, "vault3.created", vault3Result.Created, []string{
		"community-plugins.json", "plugins/open-tab-settings/",
	})
	assertPaths(t, "vault3.overwrote", vault3Result.Overwrote, []string{
		"app.json", "snippets/", "plugins/open-in-new-tab/",
	})
	if len(vault3Result.Errors) != 1 || !strings.Contains(vault3Result.Errors[0], "themes") {
		t.Fatalf("vault3.errors 不符合预期: %v", vault3Result.Errors)
	}

	assertFileEqual(t, filepath.Join(mainVault, ".obsidian", "snippets", "table-spacing-fix.css"), filepath.Join(targetVault3, ".obsidian", "snippets", "table-spacing-fix.css"))
	assertFileExists(t, filepath.Join(targetVault3, ".obsidian", "snippets", "vscode_light.css"))
	assertFileExists(t, filepath.Join(targetVault3, ".obsidian", "snippets", "target-only.css"))
	assertRegularFile(t, filepath.Join(targetVault3, ".obsidian", "themes"))
}

func coreTestCasesRoot(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("无法定位核心测试文件")
	}
	return filepath.Join(filepath.Dir(file), "..", "..", "..", "test_cases")
}

func copyTree(t *testing.T, source string, target string) {
	t.Helper()
	err := filepath.WalkDir(source, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		relPath, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}
		targetPath := filepath.Join(target, relPath)
		if entry.IsDir() {
			return os.MkdirAll(targetPath, 0o755)
		}
		info, err := entry.Info()
		if err != nil {
			return err
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if err := os.MkdirAll(filepath.Dir(targetPath), 0o755); err != nil {
			return err
		}
		return os.WriteFile(targetPath, data, info.Mode().Perm())
	})
	if err != nil {
		t.Fatalf("复制测试 vault 失败: %v", err)
	}
}

func configItemPaths(items []svc.ConfigItem) []string {
	paths := make([]string, 0, len(items))
	for _, item := range items {
		paths = append(paths, item.Path)
	}
	return paths
}

func assertContainsPaths(t *testing.T, got []string, want []string) {
	t.Helper()
	gotSet := make(map[string]struct{}, len(got))
	for _, path := range got {
		gotSet[path] = struct{}{}
	}
	for _, path := range want {
		if _, ok := gotSet[path]; !ok {
			t.Errorf("缺少配置路径: %s\n实际路径: %v", path, got)
		}
	}
}

func assertTargetPlan(t *testing.T, target svc.TargetSyncPlan, wantActions []svc.SyncPlanAction) {
	t.Helper()
	if len(target.Items) != len(coreSyncSelectedPaths) {
		t.Fatalf("%s 同步项数量不符: want %d, got %d", target.VaultPath, len(coreSyncSelectedPaths), len(target.Items))
	}
	for i, item := range target.Items {
		if item.Path != coreSyncSelectedPaths[i] || item.Action != wantActions[i] {
			t.Errorf("%s 第 %d 项不符: want %s/%s, got %s/%s", target.VaultPath, i, coreSyncSelectedPaths[i], wantActions[i], item.Path, item.Action)
		}
	}
}

func assertPaths(t *testing.T, name string, got []string, want []string) {
	t.Helper()
	gotCopy := append([]string(nil), got...)
	wantCopy := append([]string(nil), want...)
	sort.Strings(gotCopy)
	sort.Strings(wantCopy)
	if strings.Join(gotCopy, "\n") != strings.Join(wantCopy, "\n") {
		t.Errorf("%s 不符:\nwant: %v\ngot:  %v", name, wantCopy, gotCopy)
	}
}

func assertFileEqual(t *testing.T, wantPath string, gotPath string) {
	t.Helper()
	want, err := os.ReadFile(wantPath)
	if err != nil {
		t.Fatalf("读取期望文件失败: %v", err)
	}
	got, err := os.ReadFile(gotPath)
	if err != nil {
		t.Fatalf("读取实际文件失败: %v", err)
	}
	if string(got) != string(want) {
		t.Errorf("文件内容不符: %s", gotPath)
	}
}

func assertFileExists(t *testing.T, path string) {
	t.Helper()
	if _, err := os.Stat(path); err != nil {
		t.Errorf("文件不存在: %s: %v", path, err)
	}
}

func assertRegularFile(t *testing.T, path string) {
	t.Helper()
	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("目标路径不存在: %s: %v", path, err)
	}
	if info.IsDir() {
		t.Errorf("目标路径应为文件: %s", path)
	}
}
