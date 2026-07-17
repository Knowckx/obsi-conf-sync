<script lang="ts" module>
  import type { Snippet } from 'svelte';

  /** Card 的属性。 */
  export interface Props {
    children: Snippet;
  }
</script>

<script lang="ts">
  let { children }: Props = $props();
</script>

<div class="card">
  {@render children()}
</div>

<style>
  .card {
    display: flex; /* Tailwind: flex */
    flex-direction: column; /* Tailwind: flex-col */
    gap: var(--space-4); /* Tailwind 近似写法: gap-4 */
    padding: var(--space-5); /* Tailwind 近似写法: p-5 */
    background: var(--color-surface); /* Tailwind 近似写法: bg-white */
    border: 1px solid var(--color-border-subtle); /* Tailwind 近似写法: border border-slate-200 */
    border-radius: var(--radius-control); /* Tailwind 近似写法: rounded-lg */
    box-shadow: 0 1px 2px rgb(0 0 0 / 5%); /* Tailwind: shadow-sm */
  }
</style>
