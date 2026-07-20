<script lang="ts">
import type { SyncResult } from '@/lib/api/vault_service';

type Props = {
  result?: SyncResult | null;
};

let { result = null }: Props = $props();
let totalCreated = $derived(result?.targets.reduce((total, target) => total + target.created.length, 0) ?? 0);
let totalOverwrote = $derived(
  result?.targets.reduce((total, target) => total + target.overwrote.length, 0) ?? 0,
);
let totalErrors = $derived(result?.targets.reduce((total, target) => total + target.errors.length, 0) ?? 0);
</script>

<div class="step-content">
  <div class="header">
    <div>
      <h2>{totalErrors > 0 ? '同步完成，但有失败项' : '同步完成'}</h2>
      <p>以下是本次同步的实际执行结果。</p>
    </div>
  </div>

  {#if !result}
    <p class="muted">暂无执行结果</p>
  {:else}
    <div class="summary">
      <div class="summary-item created-summary">
        <strong>{totalCreated}</strong>
        <span>新增配置</span>
      </div>
      <div class="summary-item overwrite-summary">
        <strong>{totalOverwrote}</strong>
        <span>覆盖配置</span>
      </div>
      <div class:error-summary={totalErrors > 0} class="summary-item">
        <strong>{totalErrors}</strong>
        <span>失败项</span>
      </div>
    </div>

    <div class="target-list">
      {#each result.targets as target (target.vaultPath)}
        <section class="target-card">
          <h3 title={target.vaultPath}>{target.vaultPath}</h3>

          <div class="result-columns">
            <div class="result-group created-group">
              <h4>新增 <span>{target.created.length}</span></h4>
              {#if target.created.length === 0}
                <p class="muted empty-state">暂无新增项</p>
              {:else}
                <ul>
                  {#each target.created as path}
                    <li class="result-item">
                      <span class="status-icon" aria-hidden="true">✅</span>
                      <span>{path}</span>
                    </li>
                  {/each}
                </ul>
              {/if}
            </div>

            <div class="result-group overwrite-group">
              <h4>覆盖 <span>{target.overwrote.length}</span></h4>
              {#if target.overwrote.length === 0}
                <p class="muted empty-state">暂无覆盖项</p>
              {:else}
                <ul>
                  {#each target.overwrote as path}
                    <li class="result-item">
                      <span class="status-icon" aria-hidden="true">✅</span>
                      <span>{path}</span>
                    </li>
                  {/each}
                </ul>
              {/if}
            </div>
          </div>

          {#if target.errors.length > 0}
            <div class="errors">
              <h4>失败项 <span>{target.errors.length}</span></h4>
              <ul>
                {#each target.errors as error}
                  <li class="result-item">
                    <span class="status-icon" aria-hidden="true">❌</span>
                    <span>{error}</span>
                  </li>
                {/each}
              </ul>
            </div>
          {/if}
        </section>
      {/each}
    </div>
  {/if}
</div>

<style>
  .header,
  .target-list,
  .target-card,
  .result-group,
  .errors {
    display: grid;
    gap: var(--space-3);
  }

  h2,
  h3,
  h4,
  p {
    margin: 0;
  }

  p {
    color: var(--color-text-muted);
  }

  .summary {
    display: grid;
    grid-template-columns: repeat(3, minmax(0, 1fr));
    gap: var(--space-3);
  }

  .summary-item {
    display: grid;
    gap: var(--space-1);
    padding: var(--space-3);
    border: 1px solid var(--color-border);
    border-radius: var(--radius-control);
    background: var(--color-surface-muted);
  }

  .summary-item strong {
    font-size: 22px;
    line-height: 1;
  }

  .summary-item span {
    color: var(--color-text-muted);
    font-size: var(--font-size-sm);
  }

  .created-summary strong,
  .created-group h4 {
    color: var(--color-primary-text);
  }

  .overwrite-summary strong,
  .overwrite-group h4 {
    color: var(--color-text);
  }

  .error-summary strong,
  .errors h4 {
    color: var(--color-danger);
  }

  .target-card {
    overflow: hidden;
    border: 1px solid var(--color-border);
    border-radius: var(--radius-control);
    background: var(--color-surface);
  }

  h3 {
    padding: var(--space-4);
    border-bottom: 1px solid var(--color-border-subtle);
    background: var(--color-surface-muted);
    overflow-wrap: anywhere;
  }

  .result-columns {
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }

  .result-group {
    align-content: start;
    padding: var(--space-4);
  }

  .result-group + .result-group {
    border-left: 1px solid var(--color-border-subtle);
  }

  h4 {
    display: flex;
    gap: var(--space-2);
    align-items: center;
  }

  h4 span {
    color: var(--color-text-muted);
    font-size: var(--font-size-sm);
    font-weight: 400;
  }

  ul {
    display: grid;
    gap: var(--space-2);
    margin: 0;
    padding: 0;
    list-style: none;
  }

  li {
    padding: 9px var(--space-3);
    border: 1px solid var(--color-border-subtle);
    border-radius: var(--radius-control);
    background: var(--color-surface-muted);
    color: var(--color-text-subtle);
    overflow-wrap: anywhere;
  }

  .result-item {
    display: flex;
    gap: var(--space-2);
    align-items: flex-start;
  }

  .status-icon {
    flex: 0 0 auto;
    line-height: inherit;
  }

  .created-group li {
    background: var(--color-primary-bg);
    color: var(--color-primary-text);
  }

  .errors {
    padding: var(--space-4);
    border-top: 1px solid var(--color-border-subtle);
  }

  .errors li {
    border-color: color-mix(in srgb, var(--color-danger) 30%, var(--color-border));
    background: color-mix(in srgb, var(--color-danger) 6%, var(--color-surface));
    color: var(--color-danger);
  }

  .empty-state,
  .muted {
    color: var(--color-text-muted);
  }

  @media (max-width: 640px) {
    .summary,
    .result-columns {
      grid-template-columns: 1fr;
    }

    .result-group + .result-group {
      border-top: 1px solid var(--color-border-subtle);
      border-left: 0;
    }
  }
</style>
