<script lang="ts">
import type { VaultInfo } from '@/lib/api/vault_service';

type Props = {
  vaults?: VaultInfo[];
  mode?: 'read' | 'select';
  mainVault?: VaultInfo | null;
  targetVaults?: VaultInfo[];
  onMainChange?: (vault: VaultInfo) => void;
  onTargetToggle?: (vault: VaultInfo) => void;
};

let {
  vaults = [],
  mode = 'read',
  mainVault = null,
  targetVaults = [],
  onMainChange = () => {},
  onTargetToggle = () => {},
}: Props = $props();

const isMainVault = (vault: VaultInfo) => mainVault?.path === vault.path;
const isTargetVault = (vault: VaultInfo) => targetVaults.some((item) => item.path === vault.path);
</script>

{#if vaults.length === 0}
  <p>未发现 Vault</p>
{:else}
  <ul class="vault-list">
    {#each vaults as vault}
      <li>
        <div class="vault-info">
          <strong>{vault.name}</strong>
          <span>{vault.path}</span>
        </div>

        {#if mode === 'select'}
          <div class="actions">
            <button class:active={isMainVault(vault)} onclick={() => onMainChange(vault)}>
              主库
            </button>
            <button
              class:active={isTargetVault(vault)}
              disabled={isMainVault(vault)}
              onclick={() => onTargetToggle(vault)}
            >
              从库
            </button>
          </div>
        {/if}
      </li>
    {/each}
  </ul>
{/if}

<style>
  p {
    margin: 0;
  }

  .vault-list {
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

  .vault-info {
    display: grid;
    gap: 4px;
    min-width: 0;
  }

  span {
    color: #475467;
    overflow-wrap: anywhere;
  }

  .actions {
    display: flex;
    gap: 8px;
  }

  button {
    min-width: 52px;
    height: 32px;
    border: 1px solid #d0d5dd;
    border-radius: 6px;
    background: #fff;
    color: #344054;
    cursor: pointer;
  }

  button.active {
    border-color: #7f56d9;
    background: #f9f5ff;
    color: #53389e;
  }

  button:disabled {
    color: #98a2b3;
    cursor: not-allowed;
  }

  @media (max-width: 640px) {
    li {
      grid-template-columns: 1fr;
    }
  }
</style>
