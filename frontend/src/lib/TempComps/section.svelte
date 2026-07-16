<script lang="ts" module>
  import type { Snippet } from 'svelte';

  /** Section 的属性。 */
  export interface Props {
    children: Snippet;
  }
</script>

<script lang="ts">
  let { children }: Props = $props();
</script>

<section class="section">
  {@render children()}
</section>

<style>
  .section {
    display: grid; /* Tailwind: grid */
    gap: var(--space-4); /* Tailwind 近似写法: gap-4 */
  }
</style>
