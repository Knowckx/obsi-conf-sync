<script>
import { Dialogs } from '@wailsio/runtime';
import { VaultService } from '../../bindings/obsi-conf-sync/go_src/inner/svc';
import VaultList from '../components/VaultList.svelte';

let vaults = $state([]);
let root = $state('');
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

  root = selected;
  scanning = true;
  try {
    vaults = await VaultService.ScanVaults(selected);
  } catch (err) {
    error = err?.message ?? String(err);
    vaults = [];
  } finally {
    scanning = false;
  }
};
</script>

<section class="scan-view">
  <div class="toolbar">
    <button onclick={chooseAndScan} disabled={scanning}>
      {scanning ? '扫描中' : '选择目录并扫描'}
    </button>
    <input value={root} readonly placeholder="未选择目录" />
  </div>

  {#if error}
    <p class="error">{error}</p>
  {/if}

  <VaultList {vaults} />
</section>

<style>
  .scan-view {
    display: grid;
    gap: 16px;
    padding: 24px;
  }

  .toolbar {
    display: grid;
    grid-template-columns: max-content minmax(0, 1fr);
    gap: 12px;
    align-items: center;
  }

  button,
  input {
    height: 36px;
  }

  button {
    padding: 0 14px;
  }

  input {
    min-width: 0;
    padding: 0 10px;
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
