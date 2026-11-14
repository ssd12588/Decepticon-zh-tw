import { randomUUID } from "node:crypto";
import { link, lstat, mkdir, open, unlink } from "node:fs/promises";
import path from "node:path";
import { MAX_ROE_BYTES, validateRoe } from "./roeValidation.js";
import type { Command } from "./types.js";

async function missing(filename: string): Promise<void> {
  try {
    await lstat(filename);
  } catch (error) {
    if (error instanceof Error && "code" in error && error.code === "ENOENT") return;
    throw error;
  }
  throw new Error("範圍已存在。匯入不會覆寫 plan/roe.json。");
}

export const scope: Command = {
  name: "scope",
  description: "匯入操作員提供的授權測試範圍（僅限新建）",
  argumentHint: "import <file> --confirm",
  async execute(args, ctx) {
    try {
      const match = /^import\s+(.+?)\s+--confirm$/.exec(args.trim());
      if (!match?.[1]) throw new Error("用法：/scope import <file> --confirm");
      if (!process.env.DECEPTICON_ENGAGEMENT) throw new Error("匯入範圍需要在授權測試任務的工作區內執行。");
      if (!ctx.mutateWorkspace) throw new Error("匯入範圍需要處於安全暫停狀態的工作區界線。");
      const root = process.env.DECEPTICON_WORKSPACE_PATH || "/workspace";
      const candidate = path.resolve(root, match[1]);
      const handle = await open(candidate, "r");
      let raw: Buffer;
      try {
        const info = await handle.stat();
        if (!info.isFile() || info.size > MAX_ROE_BYTES) throw new Error("範圍必須是一般檔案，且不超過 1 MiB。");
        raw = Buffer.alloc(MAX_ROE_BYTES + 1);
        let length = 0;
        while (length < raw.length) {
          const { bytesRead } = await handle.read(raw, length, raw.length - length, null);
          if (!bytesRead) break;
          length += bytesRead;
        }
        raw = raw.subarray(0, length);
      } finally {
        await handle.close();
      }
      await validateRoe(raw);
      const installed = await ctx.mutateWorkspace(async () => {
        const plan = path.join(root, "plan");
        await mkdir(plan, { recursive: true });
        const directory = await lstat(plan);
        if (directory.isSymbolicLink() || !directory.isDirectory()) throw new Error("拒絕符號連結或非目錄的計畫路徑。");
        const destination = path.join(plan, "roe.json");
        await missing(destination);
        const temporary = path.join(plan, `.roe-import-${randomUUID()}`);
        const output = await open(temporary, "wx", 0o600);
        try {
          await output.writeFile(raw);
          await output.sync();
          await link(temporary, destination);
        } finally {
          await output.close();
          await unlink(temporary);
        }
      });
      if (installed) ctx.addSystemEvent("範圍已匯入 plan/roe.json。紅隊模式會強制執行此範圍。");
    } catch (error) {
      ctx.addSystemEvent(error instanceof Error ? error.message : String(error));
    }
  },
};
