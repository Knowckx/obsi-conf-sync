<script lang="ts">
type Step = {
  key: string;
  label: string;
};

type Props = {
  steps?: Step[];
  currentKey?: string;
};

let { steps = [], currentKey = '' }: Props = $props();
</script>

<nav class="step-nav" aria-label="同步步骤">
  {#each steps as step, index}
    <div class:active={step.key === currentKey}>
      <span>{index + 1}</span>
      <strong>{step.label}</strong>
    </div>
  {/each}
</nav>

<style>
  .step-nav {
    display: grid;
    align-content: start;
    gap: var(--space-2);
  }

  .step-nav div {
    display: grid;
    grid-template-columns: 28px minmax(0, 1fr);
    gap: var(--space-2);
    align-items: center;
    padding: var(--space-2);
    border: 1px solid var(--color-border-subtle);
    border-radius: var(--radius-control);
    color: var(--color-text-subtle);
  }

  .step-nav div.active {
    border-color: var(--color-primary);
    color: var(--color-primary-text);
    background: var(--color-primary-bg);
  }

  span {
    display: grid;
    place-items: center;
    width: 28px;
    height: 28px;
    border-radius: var(--radius-round);
    background: var(--color-surface-muted);
    font-size: var(--font-size-sm);
  }

  .active span {
    color: var(--color-surface);
    background: var(--color-primary);
  }

  strong {
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }
</style>
