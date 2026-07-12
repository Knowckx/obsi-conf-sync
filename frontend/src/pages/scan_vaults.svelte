<script lang="ts">
import { Dialogs } from '@wailsio/runtime';
import { Button, Input } from 'infa-s5';
import { scanVaults, type VaultInfo } from '@/lib/api/vault_service';
import VaultList from '@/lib/components/vault_list.svelte';

type Props = {
  root?: string;
  vaults?: VaultInfo[];
  onScanned?: (root: string, vaults: VaultInfo[]) => void;
};

let { root = '', vaults = [], onScanned = () => {} }: Props = $props();
let error = $state('');
let scanning = $state(false);

const chooseAndScan = async () => {
  error = '';
  const selected = await Dialogs.OpenFile({
    Title: '选择 Obsidian 目录',
    ButtonText: '选择',
    CanChooseDirectories: true,
    CanChooseFiles: false,
  });

  if (!selected) {
    return;
  }

  scanning = true;
  try {
    const foundVaults = await scanVaults(selected);
    onScanned(selected, foundVaults);
  } catch (err) {
    error = getErrMsg(err);
    onScanned(selected, []);
  } finally {
    scanning = false;
  }
};

const getErrMsg = (err: unknown): string => {
  return err instanceof Error ? err.message : String(err);
};
</script>

<div class="step-content">
  <div class="toolbar">
    <Button onclick={chooseAndScan} disabled={scanning}>
      {scanning ? '扫描中' : '选择目录'}
    </Button>
    <Input value={root} readonly placeholder="未选择目录" />
  </div>

  {#if error}
    <p class="status-error">{error}</p>
  {/if}

  <VaultList {vaults} />
</div>

<style>
  .toolbar {
    display: grid;
    grid-template-columns: max-content minmax(0, 1fr);
    gap: var(--space-3);
    align-items: center;
  }

  @media (max-width: 640px) {
    .toolbar {
      grid-template-columns: 1fr;
    }
  }
</style>
