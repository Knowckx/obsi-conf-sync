import { VaultService } from '@bindings/obsi-conf-sync/go_src/inner/svc';
import {
  SyncPlanAction,
  SyncPlan as SyncPlanModel,
  SyncRequest as SyncRequestModel,
} from '@bindings/obsi-conf-sync/go_src/inner/svc/models';

export { SyncPlanAction };

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

export type SyncPlanItem = {
  path: string;
  action: SyncPlanAction;
};

export type TargetSyncPlan = {
  vaultPath: string;
  items: SyncPlanItem[];
};

export type SyncPlan = {
  mainVaultPath: string;
  targets: TargetSyncPlan[];
};

export type TargetSyncResult = {
  vaultPath: string;
  created: string[];
  overwrote: string[];
  errors: string[];
};

export type SyncResult = {
  targets: TargetSyncResult[];
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

export const executeSyncPlan = (plan: SyncPlan): Promise<SyncResult> => {
  return VaultService.ExecuteSyncPlan(new SyncPlanModel(plan));
};
