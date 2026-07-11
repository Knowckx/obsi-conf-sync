<script>
import { Button, Card, ContentShell, PanelBg } from 'infa-s5';
import ScanVaults from './ScanVaults.svelte';
import SelectScopeStep from './steps/SelectScopeStep.svelte';
import SelectVaultsStep from './steps/SelectVaultsStep.svelte';
import StepNav from '@/lib/components/StepNav.svelte';

const steps = [
  { key: 'scan', label: '扫描库' },
  { key: 'vaults', label: '选择库' },
  { key: 'scope', label: '同步范围' },
  { key: 'plan', label: '同步计划' },
  { key: 'result', label: '执行结果' },
];

let stepIndex = $state(0);
let root = $state('');
let vaults = $state([]);
let mainVault = $state(null);
let targetVaults = $state([]);
let configItems = $state([]);
let selectedPaths = $state([]);

let currentStep = $derived(steps[stepIndex]);
let canBack = $derived(stepIndex > 0);
let canNext = $derived(getCanNext());

const setScannedVaults = (selectedRoot, foundVaults) => {
  root = selectedRoot;
  const vaultMap = new Map(vaults.map((vault) => [vault.path, vault]));
  for (const vault of foundVaults) {
    vaultMap.set(vault.path, vault);
  }
  vaults = [...vaultMap.values()];
};

const setMainVault = (vault) => {
  mainVault = vault;
  targetVaults = vaults.filter((item) => item.path !== vault.path);
  configItems = [];
  selectedPaths = [];
};

const toggleTargetVault = (vault) => {
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

function getCanNext() {
  if (currentStep.key === 'scan') {
    return vaults.length > 0;
  }

  if (currentStep.key === 'vaults') {
    return mainVault && targetVaults.length > 0;
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
        <div class="summary">
          <div>
            <span>主库</span>
            <strong>{mainVault?.name ?? '未选择'}</strong>
          </div>
          <div>
            <span>从库</span>
            <strong>{targetVaults.length}</strong>
          </div>
          <div>
            <span>已发现</span>
            <strong>{vaults.length}</strong>
          </div>
          <div>
            <span>同步项</span>
            <strong>{selectedPaths.length}</strong>
          </div>
        </div>

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
    gap: 16px;
  }

  .summary {
    display: grid;
    grid-template-columns: repeat(4, minmax(0, 1fr));
    gap: 12px;
    margin-bottom: 20px;
  }

  .summary div {
    display: grid;
    gap: 4px;
    padding: 10px 12px;
    border: 1px solid #e4e7ec;
    border-radius: 6px;
  }

  .summary span {
    color: #667085;
    font-size: 13px;
  }

  .summary strong {
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  .step-body {
    min-height: 360px;
  }

  .pending-step {
    display: grid;
    gap: 8px;
  }

  .pending-step h2,
  .pending-step p {
    margin: 0;
  }

  .footer {
    display: flex;
    justify-content: flex-end;
    gap: 10px;
    margin-top: 20px;
  }

  @media (max-width: 760px) {
    .layout,
    .summary {
      grid-template-columns: 1fr;
    }
  }
</style>
