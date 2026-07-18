import { VaultService } from '@bindings/obsi-conf-sync/go_src/inner/svc';
import { SyncRequest as SyncRequestModel } from '@bindings/obsi-conf-sync/go_src/inner/svc/models';

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

export type SyncRequest = {
  mainVaultPath: string;
  targetVaultPaths: string[];
  selectedPaths: string[];
};

export type TargetSyncPlan = {
  vaultPath: string;
  create: string[];
  overwrite: string[];
};

export type SyncPlan = {
  mainVaultPath: string;
  targets: TargetSyncPlan[];
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

export const buildSyncPlan = (request: SyncRequest): Promise<SyncPlan> => {
  return VaultService.BuildSyncPlan(new SyncRequestModel(request));
};
