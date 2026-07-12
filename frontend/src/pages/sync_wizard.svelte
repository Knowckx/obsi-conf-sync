<script lang="ts">
import { onMount } from 'svelte';
import { Button, Card, ContentShell, PanelBg } from 'infa-s5';
import { scanVaults, type ConfigItem, type VaultInfo } from '@/lib/api/vault_service';
import ScanVaults from './scan_vaults.svelte';
import SelectScopeStep from './steps/select_scope_step.svelte';
import SelectVaultsStep from './steps/select_vaults_step.svelte';
import StepNav from '@/lib/components/step_nav.svelte';

type StepKey = 'scan' | 'vaults' | 'scope' | 'plan' | 'result';

type WizardStep = {
  key: StepKey;
  label: string;
};

const steps: WizardStep[] = [
  { key: 'scan', label: '扫描库' },
  { key: 'vaults', label: '选择库' },
  { key: 'scope', label: '同步范围' },
  { key: 'plan', label: '同步计划' },
  { key: 'result', label: '执行结果' },
];

let stepIndex = $state(0);
let root = $state('');
let vaults = $state<VaultInfo[]>([]);
let mainVault = $state<VaultInfo | null>(null);
let targetVaults = $state<VaultInfo[]>([]);
let configItems = $state<ConfigItem[]>([]);
let selectedPaths = $state<string[]>([]);
let startupError = $state('');

let currentStep = $derived(steps[stepIndex]!);
let canBack = $derived(stepIndex > 0);
let canNext = $derived(getCanNext());

const setScannedVaults = (selectedRoot: string, foundVaults: VaultInfo[]) => {
  root = selectedRoot;
  const vaultMap = new Map(vaults.map((vault) => [vault.path, vault]));
  for (const vault of foundVaults) {
    vaultMap.set(vault.path, vault);
  }
  vaults = [...vaultMap.values()];
};

const setMainVault = (vault: VaultInfo) => {
  mainVault = vault;
  targetVaults = vaults.filter((item) => item.path !== vault.path);
  configItems = [];
  selectedPaths = [];
};

// 开发环境按本机预设自动进入同步范围。
const applyDevPreset = async () => {
  if (!import.meta.env.DEV || import.meta.env.VITE_DEV_AUTO_ENTER_M3 !== 'true') {
    return;
  }

  const devRoot = import.meta.env.VITE_DEV_VAULT_ROOT;
  const devMainVault = import.meta.env.VITE_DEV_MAIN_VAULT;
  if (!devRoot || !devMainVault) {
    startupError = '开发启动预设缺少 vault 根目录或主库路径';
    return;
  }

  try {
    const foundVaults = await scanVaults(devRoot);
    setScannedVaults(devRoot, foundVaults);

    const selectedMainVault = foundVaults.find((vault) => vault.path === devMainVault);
    if (!selectedMainVault) {
      throw new Error(`开发预设主库未在扫描结果中：${devMainVault}`);
    }

    setMainVault(selectedMainVault);
    if (targetVaults.length === 0) {
      throw new Error('开发预设没有可用的从库');
    }

    stepIndex = 2;
  } catch (err) {
    startupError = err instanceof Error ? err.message : String(err);
    stepIndex = 0;
  }
};

onMount(() => {
  void applyDevPreset();
});

const toggleTargetVault = (vault: VaultInfo) => {
  if (mainVault?.path === vault.path) {
    return;
  }

  const exists = targetVaults.some((item) => item.path === vault.path);
  targetVaults = exists
    ? targetVaults.filter((item) => item.path !== vault.path)
    : [...targetVaults, vault];
};

const goBack = () => {
  if (canBack) {
    stepIndex -= 1;
  }
};

const goNext = () => {
  if (canNext && stepIndex < steps.length - 1) {
    stepIndex += 1;
  }
};

function getCanNext(): boolean {
  if (currentStep.key === 'scan') {
    return vaults.length > 0;
  }

  if (currentStep.key === 'vaults') {
    return Boolean(mainVault && targetVaults.length > 0);
  }

  if (currentStep.key === 'scope') {
    return selectedPaths.length > 0;
  }

  return false;
}
</script>

<PanelBg>
  <ContentShell maxWidth="max-w-6xl">
    <div class="layout">
      <StepNav {steps} currentKey={currentStep.key} />

      <Card>
        {#if startupError}
          <p class="status-error startup-error">{startupError}</p>
        {/if}

        <section class="step-body">
          {#if currentStep.key === 'scan'}
            <ScanVaults {root} {vaults} onScanned={setScannedVaults} />
          {:else if currentStep.key === 'vaults'}
            <SelectVaultsStep
              {vaults}
              {mainVault}
              {targetVaults}
              onMainChange={setMainVault}
              onTargetToggle={toggleTargetVault}
            />
          {:else if currentStep.key === 'scope'}
            <SelectScopeStep
              {mainVault}
              {configItems}
              {selectedPaths}
              onConfigItemsChange={(items) => (configItems = items)}
              onSelectedPathsChange={(paths) => (selectedPaths = paths)}
            />
          {:else}
            <div class="pending-step">
              <h2>{currentStep.label}</h2>
              <p>等待对应后端接口补齐。</p>
            </div>
          {/if}
        </section>

        <div class="footer">
          <Button onclick={goBack} disabled={!canBack}>上一步</Button>
          <Button onclick={goNext} disabled={!canNext}>下一步</Button>
        </div>
      </Card>
    </div>
  </ContentShell>
</PanelBg>

<style>
  .layout {
    display: grid;
    grid-template-columns: 160px minmax(0, 1fr);
    gap: var(--space-4);
  }

  .step-body {
    min-height: 360px;
  }

  .startup-error {
    margin-bottom: var(--space-4);
  }

  .pending-step {
    display: grid;
    gap: var(--space-2);
  }

  .pending-step h2,
  .pending-step p {
    margin: 0;
  }

  .footer {
    display: flex;
    justify-content: flex-end;
    gap: 10px;
    margin-top: var(--space-5);
  }

  @media (max-width: 760px) {
    .layout {
      grid-template-columns: 1fr;
    }
  }
</style>
