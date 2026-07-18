package svc

// SyncRequest 描述一次同步计划请求。
type SyncRequest struct {
	MainVaultPath string `json:"mainVaultPath"`
	// 从库
	TargetVaultPaths []string `json:"targetVaultPaths"`
	// 选择要同步的配置
	SelectedPaths []string `json:"selectedPaths"`
}

// SyncPlan 描述一次同步计划。
type SyncPlan struct {
	MainVaultPath string           `json:"mainVaultPath"`
	Targets       []TargetSyncPlan `json:"targets"`
}

// TargetSyncPlan 描述单个目标库的同步计划。
type TargetSyncPlan struct {
	// 当前目标库的绝对路径
	VaultPath string `json:"vaultPath"`
	// 目标库中尚不存在的配置路径
	Create []string `json:"create"`
	// 目标库中已经存在的配置路径
	Overwrite []string `json:"overwrite"`
}
