package svc

// VaultService exposes vault scan APIs to the frontend.
type VaultService struct{}

// VaultInfo describes an Obsidian vault found by scanning.
type VaultInfo struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

// ScanVaults scans root and returns directories containing .obsidian/.
func (s *VaultService) ScanVaults(root string) ([]VaultInfo, error) {
	return []VaultInfo{}, nil
}
