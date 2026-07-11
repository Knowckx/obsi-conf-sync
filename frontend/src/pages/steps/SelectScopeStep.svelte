<script>
import { Button } from 'infa-s5';
import { VaultService } from '../../../bindings/obsi-conf-sync/go_src/inner/svc';

let {
  mainVault = null,
  configItems = [],
  selectedPaths = [],
  onConfigItemsChange = () => {},
  onSelectedPathsChange = () => {},
} = $props();

let error = $state('');
let loading = $state(false);
let loadedVaultPath = $state('');

$effect(() => {
  if (!mainVault || configItems.length > 0 || loadedVaultPath === mainVault.path) {
    return;
  }
  loadConfigItems();
});

const loadConfigItems = async () => {
  if (!mainVault) {
    return;
  }

  error = '';
  loading = true;
  loadedVaultPath = mainVault.path;

  try {
    const items = await VaultService.ListConfigItems(mainVault.path);
    onConfigItemsChange(items);
    if (selectedPaths.length === 0) {
      onSelectedPathsChange(items.filter(isDefaultSelected).map((item) => item.path));
    }
  } catch (err) {
    error = err?.message ?? String(err);
  } finally {
    loading = false;
  }
};

const togglePath = (path) => {
  const exists = selectedPaths.includes(path);
  onSelectedPathsChange(
    exists ? selectedPaths.filter((item) => item !== path) : [...selectedPaths, path],
  );
};

const selectAll = () => {
  onSelectedPathsChange(configItems.map((item) => item.path));
};

const selectDefault = () => {
  onSelectedPathsChange(configItems.filter(isDefaultSelected).map((item) => item.path));
};

const openConfigDir = async () => {
  if (!mainVault) {
    return;
  }

  error = '';
  try {
    await VaultService.OpenVaultConfigDir(mainVault.path);
  } catch (err) {
    error = err?.message ?? String(err);
  }
};

function isDefaultSelected(item) {
  return item.path.includes('/') || !item.path.startsWith('workspace');
}
</script>

<div class="select-scope">
  <div class="header">
    <div>
      <h2>选择同步范围</h2>
      <p>{mainVault?.name ?? ''} 的 .obsidian 配置项</p>
    </div>
    <div class="actions">
      <Button onclick={openConfigDir} disabled={!mainVault}>打开配置文件夹</Button>
      <Button onclick={selectDefault} disabled={loading || configItems.length === 0}>默认</Button>
      <Button onclick={selectAll} disabled={loading || configItems.length === 0}>全选</Button>
    </div>
  </div>

  {#if loading}
    <p class="muted">加载中</p>
  {:else if error}
    <p class="error">{error}</p>
  {:else if configItems.length === 0}
    <p class="muted">未发现配置项</p>
  {:else}
    <ul class="config-list">
      {#each configItems as item}
        <li>
          <label>
            <input
              type="checkbox"
              checked={selectedPaths.includes(item.path)}
              onchange={() => togglePath(item.path)}
            />
            <span>{item.path}</span>
          </label>
          {#if !isDefaultSelected(item)}
            <em>默认跳过</em>
          {/if}
        </li>
      {/each}
    </ul>
  {/if}
</div>

<style>
  .select-scope {
    display: grid;
    gap: 16px;
  }

  .header {
    display: flex;
    justify-content: space-between;
    gap: 16px;
  }

  .actions {
    display: flex;
    gap: 8px;
  }

  h2,
  p {
    margin: 0;
  }

  p {
    color: #667085;
  }

  .config-list {
    display: grid;
    gap: 8px;
    list-style: none;
    margin: 0;
    padding: 0;
  }

  li {
    display: grid;
    grid-template-columns: minmax(0, 1fr) max-content;
    gap: 12px;
    align-items: center;
    padding: 10px 12px;
    border: 1px solid #d0d5dd;
    border-radius: 6px;
  }

  label {
    display: flex;
    gap: 10px;
    align-items: center;
    min-width: 0;
  }

  input {
    width: 16px;
    height: 16px;
  }

  span {
    overflow-wrap: anywhere;
  }

  em {
    color: #667085;
    font-style: normal;
    font-size: 13px;
  }

  .muted {
    color: #667085;
  }

  .error {
    color: #b42318;
  }

  @media (max-width: 640px) {
    .header,
    li {
      grid-template-columns: 1fr;
    }

    .header {
      display: grid;
    }
  }
</style>
