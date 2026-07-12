import { VaultService } from '@bindings/obsi-conf-sync/go_src/inner/svc';

export type VaultInfo = {
  path: string;
  name: string;
};

export type ConfigItem = {
  path: string;
  name: string;
  version: string;
  isDir: boolean;
  description: string;
  defaultSelected: boolean;
  isPlugin: boolean;
};

export const scanVaults = (root: string): Promise<VaultInfo[]> => {
  return VaultService.ScanVaults(root);
};

export const listConfigItems = (vaultPath: string): Promise<ConfigItem[]> => {
  return VaultService.ListConfigItems(vaultPath);
};

export const openVaultConfigDir = (vaultPath: string): Promise<void> => {
  return VaultService.OpenVaultConfigDir(vaultPath);
};
