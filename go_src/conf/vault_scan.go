package conf

// SkipVaultScanDirNames 是 vault 扫描时固定跳过的目录名。
var SkipVaultScanDirNames = map[string]struct{}{
	".git":         {},
	"node_modules": {},
}
