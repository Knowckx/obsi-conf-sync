<script lang="ts">
import type { SyncPlan } from '@/lib/api/vault_service';

type Props = {
  plan?: SyncPlan | null;
  loading?: boolean;
  error?: string;
};

let { plan = null, loading = false, error = '' }: Props = $props();
</script>

<div class="step-content">
  <div class="header">
    <div>
      <h2>同步计划</h2>
      <p>确认各目标库将要复制和跳过的配置项。</p>
    </div>
  </div>

  {#if loading}
    <p class="muted">正在生成同步计划…</p>
  {:else if error}
    <p class="status-error">{error}</p>
  {:else if !plan || plan.targets.length === 0}
    <p class="muted">暂无同步目标</p>
  {:else}
    <div class="sync-source">
      <section class="source-card">
        <span class="source-role">主库</span>
        <span class="source-path" title={plan.mainVaultPath}>{plan.mainVaultPath}</span>
      </section>
      <span class="sync-arrow" aria-hidden="true">↓</span>
    </div>

    <div class="target-list">
      {#each plan.targets as target (target.vaultPath)}
        <section class="target-card">
          <div class="target-header">
            <div class="target-context">
              <span class="target-role">目标库</span>
              <h3 title={target.vaultPath}>{target.vaultPath}</h3>
            </div>
            <div class="target-summary">
              <span><strong>{target.create.length}</strong> 项新增</span>
              <span><strong>{target.overwrite.length}</strong> 项覆盖</span>
            </div>
          </div>

          <div class="plan-columns">
            <div class="plan-group create-group">
              <div class="group-title">
                <span class="group-marker" aria-hidden="true"></span>
                <h4>新增配置</h4>
                <span class="count">{target.create.length}</span>
              </div>
              {#if target.create.length === 0}
                <p class="muted empty-state">没有需要新增的配置</p>
              {:else}
                <ul class="path-list">
                  {#each target.create as path}
                    <li>{path}</li>
                  {/each}
                </ul>
              {/if}
            </div>

            <div class="plan-group overwrite-group">
              <div class="group-title">
                <span class="group-marker" aria-hidden="true"></span>
                <h4>覆盖现有配置</h4>
                <span class="count">{target.overwrite.length}</span>
              </div>
              {#if target.overwrite.length === 0}
                <p class="muted empty-state">没有需要覆盖的配置</p>
              {:else}
                <ul class="path-list">
                  {#each target.overwrite as path}
                    <li>{path}</li>
                  {/each}
                </ul>
              {/if}
            </div>
          </div>
        </section>
      {/each}
    </div>
  {/if}
</div>

<style>
  .header,
  .target-list,
  .target-card,
  .plan-group {
    display: grid;
    gap: var(--space-3);
  }

  .sync-source {
    display: grid;
    justify-items: start;
    gap: var(--space-2);
  }

  .source-card {
    display: flex;
    gap: var(--space-2);
    align-items: center;
    max-width: 100%;
    padding: var(--space-3) var(--space-4);
    border: 1px solid var(--color-primary);
    border-radius: var(--radius-control);
    background: var(--color-primary-bg);
  }

  .source-role {
    flex: none;
    padding: 3px var(--space-2);
    border-radius: var(--radius-round);
    background: var(--color-primary);
    color: var(--color-surface);
    font-size: var(--font-size-sm);
    font-weight: 600;
  }

  .source-path {
    min-width: 0;
    overflow: hidden;
    color: var(--color-primary-text);
    font-weight: 500;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  .sync-arrow {
    padding-left: var(--space-4);
    color: var(--color-text-muted);
    font-size: 18px;
    line-height: 1;
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

  .target-card {
    overflow: hidden;
    border: 1px solid var(--color-border);
    border-radius: var(--radius-control);
    background: var(--color-surface);
  }

  .target-header {
    display: flex;
    justify-content: space-between;
    gap: var(--space-4);
    align-items: center;
    padding: var(--space-4);
    border-bottom: 1px solid var(--color-border-subtle);
    background: var(--color-surface-muted);
  }

  .target-context {
    display: flex;
    gap: var(--space-2);
    align-items: center;
    min-width: 0;
  }

  .target-role,
  .count {
    flex: none;
    padding: 3px var(--space-2);
    border-radius: var(--radius-round);
    font-size: var(--font-size-sm);
  }

  .target-role {
    background: var(--color-primary-bg);
    color: var(--color-primary-text);
    font-weight: 500;
  }

  h3 {
    min-width: 0;
    font-size: 15px;
    overflow-wrap: anywhere;
  }

  .target-summary {
    display: flex;
    gap: var(--space-2);
    flex: none;
    color: var(--color-text-muted);
    font-size: var(--font-size-sm);
  }

  .target-summary span {
    padding: 5px var(--space-2);
    border: 1px solid var(--color-border-subtle);
    border-radius: var(--radius-control);
    background: var(--color-surface);
  }

  .target-summary strong {
    margin-right: var(--space-1);
    color: var(--color-text);
  }

  .plan-columns {
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }

  .plan-group {
    align-content: start;
    padding: var(--space-4);
  }

  .plan-group + .plan-group {
    border-left: 1px solid var(--color-border-subtle);
  }

  .group-title {
    display: flex;
    gap: var(--space-2);
    align-items: center;
  }

  .group-marker {
    width: 8px;
    height: 8px;
    border-radius: var(--radius-round);
  }

  .create-group .group-marker {
    background: var(--color-primary);
  }

  .overwrite-group .group-marker {
    background: var(--color-danger);
  }

  .count {
    background: var(--color-surface-muted);
    color: var(--color-text-muted);
    font-weight: 400;
  }

  .path-list {
    display: grid;
    gap: var(--space-2);
    margin: 0;
    padding: 0;
    list-style: none;
  }

  .path-list li {
    padding: 9px var(--space-3);
    border: 1px solid var(--color-border-subtle);
    border-radius: var(--radius-control);
    background: var(--color-surface-muted);
    color: var(--color-text-subtle);
    overflow-wrap: anywhere;
  }

  .create-group .path-list li {
    border-color: var(--color-border);
    background: var(--color-primary-bg);
    color: var(--color-primary-text);
  }

  .overwrite-group .path-list li {
    color: var(--color-text);
  }

  .empty-state {
    padding: var(--space-3);
    border: 1px dashed var(--color-border);
    border-radius: var(--radius-control);
    text-align: center;
  }

  .muted {
    color: var(--color-text-muted);
  }

  @media (max-width: 760px) {
    .target-header {
      display: grid;
    }

    .target-summary {
      flex-wrap: wrap;
    }

    .plan-columns {
      grid-template-columns: 1fr;
    }

    .plan-group + .plan-group {
      border-top: 1px solid var(--color-border-subtle);
      border-left: 0;
    }
  }
</style>
