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
    gap: 8px;
  }

  .step-nav div {
    display: grid;
    grid-template-columns: 28px minmax(0, 1fr);
    gap: 8px;
    align-items: center;
    padding: 8px;
    border: 1px solid #e4e7ec;
    border-radius: 6px;
    color: #475467;
  }

  .step-nav div.active {
    border-color: #7f56d9;
    color: #53389e;
    background: #f9f5ff;
  }

  span {
    display: grid;
    place-items: center;
    width: 28px;
    height: 28px;
    border-radius: 999px;
    background: #f2f4f7;
    font-size: 13px;
  }

  .active span {
    color: #fff;
    background: #7f56d9;
  }

  strong {
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }
</style>
