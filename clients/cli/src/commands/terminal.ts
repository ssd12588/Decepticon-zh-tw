import type { Command } from "./types.js";
export const terminalCommands: Command[] = ["tui", "mouse"].map(name => ({
  name, description: name === "tui" ? "選擇全螢幕操作或終端機原生捲動" : "啟用滑鼠操作，或保留終端機原生選取",
  argumentHint: name === "tui" ? "[fullscreen | inline]" : "[on | off]",
  execute: (args, context) => context.terminalCommand?.(name as "tui" | "mouse", args.trim()),
}));
terminalCommands.push({ name: "answer", description: "回到待回答的問題", execute: (_args, context) => context.addSystemEvent("沒有待回答的問題。送出訊息以繼續，或用 /resume 選擇已儲存的工作階段。") });
