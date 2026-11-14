import type { Command } from "./types.js";

const clear: Command = {
  name: "clear",
  description: "開始全新對話（保留已儲存的工作階段）",
  aliases: ["new"],
  execute(args, ctx) {
    if (args.trim()) { ctx.addSystemEvent("用法：/clear 或 /new（不接受參數）"); return; }
    ctx.clearEvents();
  },
};

export default clear;
