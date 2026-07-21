<script lang="ts">
import {
  SyncPlanAction,
  type SyncPlan,
  type SyncPlanItem,
} from '@/lib/api/vault_service';

type Props = {
  plan?: SyncPlan | null;
  loading?: boolean;
  error?: string;
};

let { plan = null, loading = false, error = '' }: Props = $props();

const countAction = (items: SyncPlanItem[], action: SyncPlanAction) => {
  return items.filter((item) => item.action === action).length;
};
</script>

<div class="step-content">
  <div class="header">
    <div>
      <h2>同步计划</h2>
      <p>确认各目标库将要新增或覆盖的配置项。</p>
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
              <span><strong>{countAction(target.items, SyncPlanAction.SyncPlanActionCreate)}</strong> 项新增</span>
              <span><strong>{countAction(target.items, SyncPlanAction.SyncPlanActionOverwrite)}</strong> 项覆盖</span>
            </div>
          </div>

          <ul class="plan-list">
            {#each target.items as item (item.path)}
              <li>
                <span class="item-path">{item.path}</span>
                <span
                  class="action-badge"
                  class:create-action={item.action === SyncPlanAction.SyncPlanActionCreate}
                  class:overwrite-action={item.action === SyncPlanAction.SyncPlanActionOverwrite}
                >
                  {item.action === SyncPlanAction.SyncPlanActionCreate ? '新增' : '覆盖'}
                </span>
              </li>
            {/each}
          </ul>
        </section>
      {/each}
    </div>
  {/if}
</div>

<style>
  .header,
  .target-list,
  .target-card {
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
  .action-badge {
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

  .plan-list {
    display: grid;
    gap: var(--space-2);
    margin: 0;
    padding: 0 var(--space-4) var(--space-4);
    list-style: none;
  }

  .plan-list li {
    display: flex;
    justify-content: space-between;
    gap: var(--space-3);
    align-items: center;
    padding: 9px var(--space-3);
    border: 1px solid var(--color-border-subtle);
    border-radius: var(--radius-control);
    background: var(--color-surface-muted);
    color: var(--color-text-subtle);
  }

  .item-path {
    min-width: 0;
    overflow-wrap: anywhere;
  }

  .create-action {
    background: var(--color-success-bg);
    color: var(--color-success);
    font-weight: 500;
  }

  .overwrite-action {
    background: var(--color-surface);
    color: var(--color-danger);
    font-weight: 500;
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

    .plan-list li {
      align-items: flex-start;
    }
  }
</style>
