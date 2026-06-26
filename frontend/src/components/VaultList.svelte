<script>
let { vaults = [] } = $props();
</script>

{#if vaults.length === 0}
  <p>未发现 Vault</p>
{:else}
  <ul>
    {#each vaults as vault}
      <li>
        <strong>{vault.name}</strong>
        <span>{vault.path}</span>
      </li>
    {/each}
  </ul>
{/if}
