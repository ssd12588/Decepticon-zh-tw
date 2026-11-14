import type { SessionDoctor } from "./sessionDoctor.js";
import type { BudgetUsage } from "../components/BudgetFooter.js";

export interface DoctorReportInput {
  readonly live?: SessionDoctor;
  readonly version: string;
  readonly engagement?: string;
  readonly target?: string;
  readonly targetType?: string;
  readonly authorizationConfirmed: boolean;
  readonly workflowId?: string;
  readonly workflowVersion?: string;
  readonly planningWorkflowId?: string;
  readonly planningWorkflowVersion?: string;
  readonly assistantId: string;
  readonly runState: string;
  readonly modelOverride?: string;
  readonly threadId?: string;
  readonly connection: string;
  readonly budgetUsage: BudgetUsage | null;
}

function workflowPin(id?: string, version?: string): string {
  return id ? `${id}@${version || "latest"}` : "預設（未鎖定）";
}

/** Render the read-only in-session diagnostics report as a single block. */
export function formatDoctorReport(input: DoctorReportInput): string {
  const budget = input.budgetUsage
    ? `budget ${input.budgetUsage.pct}% (${input.budgetUsage.spent} / ${input.budgetUsage.cap} USD)`
    : "本次執行尚無 budget_warning 事件";

  return [
    `Decepticon ${input.version} · doctor (in-session, read-only)`,
    "",
    `Engagement: ${input.engagement || "未選擇"}`,
    `Scope/ROE:  target ${input.target || "none"} (${input.target ? input.targetType || "direct_scope" : "未設定範圍"})`
      + ` · authorization ${input.authorizationConfirmed ? "已確認" : "未確認"}`,
    `Workflow selection: execution ${workflowPin(input.workflowId, input.workflowVersion)}`
      + ` · planning ${workflowPin(input.planningWorkflowId, input.planningWorkflowVersion)}`,
    ...(input.live ? [
      `Machine RoE: ${input.live.scope}`,
      `Checkpoint: ${input.live.checkpoint}`,
      `OPPLAN revision: ${input.live.opplanRevision}`,
      `Pinned workflows: ${input.live.workflowPins.join(" · ") || "未記錄"}`,
      `Pending: ${input.live.pending}`,
      `Plugin bundles: ${input.live.bundles}`,
      `Runtime overrides: ${input.live.controls || "預設值或未記錄"}`,
    ] : []),
    `Agent:      ${input.assistantId} · run ${input.runState}`,
    `Model:      ${input.modelOverride || "自動供應商路由（本次工作階段未覆寫）"}`,
    `Session:    ${input.threadId || "尚未開始"} · LangGraph ${input.connection}`,
    `Budget:     ${budget}`,
    "",
    "主機端執行期細節（middleware 插槽矩陣、KG、Skillogy、MCP 驗證）：",
    `decepticon doctor --deep${input.engagement ? ` --engagement ${input.engagement}` : ""}`,
  ].join("\n");
}
