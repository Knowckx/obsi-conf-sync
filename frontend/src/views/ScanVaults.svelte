<script>
import VaultList from '../components/VaultList.svelte';

let vaults = $state([]);
let root = $state('');
let error = $state('');
</script>

<section>
  <h1>扫描 Vault</h1>
  <input bind:value={root} placeholder="选择目录后显示路径" />
  {#if error}
    <p>{error}</p>
  {/if}
  <VaultList {vaults} />
</section>
