import { machineScopeTargets } from "@decepticon/streaming";
import { lstat, open } from "node:fs/promises";
import path from "node:path";

function workspaceRoot(): string | null {
  if (!process.env.DECEPTICON_ENGAGEMENT) return null;
  return process.env.DECEPTICON_WORKSPACE_PATH || "/workspace";
}

// Red is gated on the engagement's enforced machine scope: creating the
// engagement with targets (or importing plan/roe.json) is the authorization.
// There is no separate approval marker — the RoE middleware enforces the
// scope on every tool call.
export async function assertRedApproved(): Promise<void> {
  const root = workspaceRoot();
  if (!root) throw new Error("紅隊模式需要一個已授權目標的授權測試任務工作區。");
  const roePath = path.join(root, "plan", "roe.json");
  const plan = await lstat(path.join(root, "plan"));
  if (!plan.isDirectory() || plan.isSymbolicLink()) {
    throw new Error("拒絕符號連結或非目錄的計畫路徑。");
  }
  // Open first, then verify the opened inode against the path: a swap between
  // the check and the read cannot redirect the content that gets parsed, and
  // the symlink refusal compares the path itself against the opened identity.
  const handle = await open(roePath, "r");
  let roe: unknown;
  try {
    const opened = await handle.stat();
    if (!opened.isFile()) throw new Error("拒絕非一般檔案的範圍路徑。");
    const current = await lstat(roePath);
    if (current.isSymbolicLink() || current.dev !== opened.dev || current.ino !== opened.ino) {
      throw new Error("拒絕符號連結或非一般檔案的範圍路徑。");
    }
    try {
      roe = JSON.parse(await handle.readFile("utf-8"));
    } catch {
      throw new Error("plan/roe.json 不是有效的範圍文件。請用 /scope import <file> --confirm 匯入範圍。");
    }
  } finally {
    await handle.close();
  }
  if (!machineScopeTargets(roe).length) {
    throw new Error("此任務的範圍缺少強制執行用的 machine 區塊。請在建立任務時指定目標，或用 /scope import <file> --confirm 匯入更新後的 plan/roe.json。");
  }
}
