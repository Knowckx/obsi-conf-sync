<script>
let { vaults = [] } = $props();
</script>

{#if vaults.length === 0}
  <p>未发现 Vault</p>
{:else}
  <ul class="vault-list">
    {#each vaults as vault}
      <li>
        <strong>{vault.name}</strong>
        <span>{vault.path}</span>
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
    gap: 4px;
    padding: 10px 12px;
    border: 1px solid #d0d5dd;
    border-radius: 6px;
  }

  span {
    color: #475467;
    overflow-wrap: anywhere;
  }
</style>
