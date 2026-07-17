<script lang="ts" module>
  import type { Snippet } from 'svelte';

  /** AppLayout 的属性。 */
  export interface Props {
    children: Snippet;
  }
</script>

<script lang="ts">
  let { children }: Props = $props();
</script>

<div class="app-layout">
  {@render children()}
</div>

<style>
  .app-layout {
    min-height: 100vh; /* Tailwind: min-h-screen */
    background: var(--color-surface-muted); /* Tailwind 近似写法: bg-slate-50 */
    color: var(--color-text); /* Tailwind 近似写法: text-slate-900 */
    display: flex; /* Tailwind: flex */
    flex-direction: column; /* Tailwind: flex-col */
  }
</style>
