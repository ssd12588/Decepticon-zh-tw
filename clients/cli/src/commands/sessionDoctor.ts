import { describeRunControls } from "../utils/runControls.js";
import { assertRedApproved } from "./planMode.js";
import { readThreadSnapshot } from "../utils/threadStore.js";

export interface SessionDoctor {
  readonly controls?: string;
  readonly scope: string;
  readonly checkpoint: string;
  readonly opplanRevision: string;
  readonly workflowPins: readonly string[];
  readonly pending: string;
  readonly bundles: string;
}
function record(value: unknown): Record<string, unknown> {
  return value !== null && typeof value === "object" && !Array.isArray(value) ? value as Record<string, unknown> : {};
}
function pin(value: unknown): string | null {
  try {
    const snapshot = record(typeof value === "string" ? JSON.parse(value) : value);
    return typeof snapshot.id === "string" && typeof snapshot.version === "string" &&
      typeof snapshot.sha256 === "string" && /^[a-f0-9]{64}$/.test(snapshot.sha256)
      ? `${snapshot.id}@${snapshot.version} sha256:${snapshot.sha256}` : null;
  } catch { return null; }
}

export async function readSessionDoctor(threadId: string | null): Promise<SessionDoctor> {
  let scope = "機器範圍無法使用";
  try { await assertRedApproved(); scope = "機器範圍就緒（執行期會強制檢查目標與授權時段）"; }
  catch { scope = "機器範圍無法使用或無效；紅隊模式仍受限制"; }
  const [stateResult, bundlesResult] = await Promise.allSettled([
    threadId ? readThreadSnapshot(threadId) : Promise.resolve(null),
    fetch(`${process.env.DECEPTICON_API_URL || "http://localhost:2024"}/_decepticon/bundles`, { signal: AbortSignal.timeout(10_000) }).then(async response => {
      if (!response.ok) throw new Error(`HTTP ${response.status}`);
      const data = record(await response.json());
      if (!Array.isArray(data.bundles)) throw new Error("外掛組合回應格式無效");
      return data.bundles.flatMap((value: unknown) => {
        const item = record(value);
        return typeof item.name === "string" && typeof item.enabled === "boolean" ? [`${item.name}: ${item.enabled ? "enabled" : "disabled"}`] : [];
      }).slice(0, 64).join(", ") || "未註冊任何項目";
    }),
  ]);
  const state = stateResult.status === "fulfilled" ? record(stateResult.value) : {};
  const values = record(state.values);
  const pins = [pin(values.workflow_snapshot), pin(values.planning_workflow_snapshot)].filter((value): value is string => value !== null);
  const checkpoint = record(state.checkpoint);
  const tasks = Array.isArray(state.tasks) ? state.tasks : [];
  const next = Array.isArray(state.next) ? state.next.filter((value): value is string => typeof value === "string").slice(0, 32) : [];
  const unavailable = stateResult.status === "rejected" ? "無法使用（讀取 checkpoint 失敗，或工作階段範圍不一致）" : "尚未記錄";
  return {
    scope,
    controls: describeRunControls(values.run_controls),
    checkpoint: typeof checkpoint.checkpoint_id === "string" ? checkpoint.checkpoint_id : unavailable,
    opplanRevision: typeof values.plan_revision === "number" ? String(values.plan_revision) : unavailable,
    workflowPins: pins,
    pending: stateResult.status === "rejected" ? unavailable : `${tasks.length} task(s); next ${next.join(", ") || "none"}`,
    bundles: bundlesResult.status === "fulfilled" ? bundlesResult.value : "無法使用（讀取執行期外掛組合失敗）",
  };
}
