<script lang="ts" module>
  import type { Snippet } from 'svelte';

  /** Page 的属性。 */
  export interface Props {
    children: Snippet;
  }
</script>

<script lang="ts">
  let { children }: Props = $props();
</script>

<main class="page">
  {@render children()}
</main>

<style>
  .page {
    box-sizing: border-box; /* Tailwind: box-border */
    width: 100%; /* Tailwind: w-full */
    max-width: 72rem; /* Tailwind: max-w-6xl */
    margin-inline: auto; /* Tailwind: mx-auto */
    padding: 2.5rem 1.5rem; /* Tailwind: px-6 py-10 */
  }
</style>
