<script lang="ts">
import type { VaultInfo } from '@/lib/api/vault_service';
import VaultList from '@/lib/components/VaultList.svelte';

type Props = {
  vaults?: VaultInfo[];
  mainVault?: VaultInfo | null;
  targetVaults?: VaultInfo[];
  onMainChange?: (vault: VaultInfo) => void;
  onTargetToggle?: (vault: VaultInfo) => void;
};

let {
  vaults = [],
  mainVault = null,
  targetVaults = [],
  onMainChange = () => {},
  onTargetToggle = () => {},
}: Props = $props();
</script>

<div class="select-vaults">
  <div>
    <h2>选择主库和从库</h2>
    <p>主库配置会覆盖到选中的从库。</p>
  </div>

  <VaultList
    {vaults}
    {mainVault}
    {targetVaults}
    mode="select"
    onMainChange={onMainChange}
    onTargetToggle={onTargetToggle}
  />
</div>

<style>
  .select-vaults {
    display: grid;
    gap: 16px;
  }

  h2,
  p {
    margin: 0;
  }

  p {
    color: #667085;
  }
</style>
