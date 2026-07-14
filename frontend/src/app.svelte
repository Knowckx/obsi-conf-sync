<script lang="ts">
import { onMount } from 'svelte';
import { Window as WailsWindow } from '@wailsio/runtime';
import SyncWizard from './pages/sync_wizard.svelte';

onMount(async () => {
  if (!import.meta.env.DEV) {
    return;
  }

  const before = await WailsWindow.GetZoom();
  await WailsWindow.ZoomReset();
  const after = await WailsWindow.GetZoom();

  console.table([{ before, after, innerWidth, outerWidth, devicePixelRatio }]);
});
</script>

<main>
  <SyncWizard />
</main>

<style>
  main {
    min-height: 100vh;
  }
</style>
