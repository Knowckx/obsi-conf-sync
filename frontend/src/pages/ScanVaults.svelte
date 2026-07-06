<script>
import { Dialogs } from '@wailsio/runtime';
import { Button, Input } from 'infa-s5';
import { VaultService } from '../../bindings/obsi-conf-sync/go_src/inner/svc';
import VaultList from '@/lib/components/VaultList.svelte';

let { root = '', vaults = [], onScanned = () => {} } = $props();
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
    const foundVaults = await VaultService.ScanVaults(selected);
    onScanned(selected, foundVaults);
  } catch (err) {
    error = err?.message ?? String(err);
    onScanned(selected, []);
  } finally {
    scanning = false;
  }
};
</script>

<div class="scan-vaults">
  <div class="toolbar">
    <Button onclick={chooseAndScan} disabled={scanning}>
      {scanning ? '扫描中' : '选择目录'}
    </Button>
    <Input value={root} readonly placeholder="未选择目录" />
  </div>

  {#if error}
    <p class="error">{error}</p>
  {/if}

  <VaultList {vaults} />
</div>

<style>
  .scan-vaults {
    display: grid;
    gap: 16px;
  }

  .toolbar {
    display: grid;
    grid-template-columns: max-content minmax(0, 1fr);
    gap: 12px;
    align-items: center;
  }

  .error {
    color: #b42318;
    margin: 0;
  }

  @media (max-width: 640px) {
    .toolbar {
      grid-template-columns: 1fr;
    }
  }
</style>
