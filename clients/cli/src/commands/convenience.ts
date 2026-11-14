import type { Command } from "./types.js";
import mode from "./mode.js";
import { KEYBOARD_HELP } from "../components/keyboardHelp.js";
import { CLI_VERSION } from "../utils/version.js";

export const convenienceCommands: Command[] = [
  {
    name: "plan",
    description: "進入選用的計畫模式（等同 /mode plan）",
    execute(args, context) {
      if (args.trim()) { context.addSystemEvent("用法：/plan（接著送出你的規劃訊息）"); return; }
      return mode.execute("plan", context);
    },
  },
  {
    name: "shortcuts",
    description: "顯示鍵盤快速鍵",
    execute(_args, context) { context.addSystemEvent(KEYBOARD_HELP.join("\n")); },
  },
  {
    name: "version",
    aliases: ["about"],
    description: "顯示執行中的 Decepticon 版本與 CLI 執行環境",
    execute(_args, context) {
      context.addSystemEvent(`Decepticon ${CLI_VERSION}\nCLI runtime: Node ${process.version} · ${process.platform}/${process.arch}`);
    },
  },
  {
    name: "pwd",
    aliases: ["cwd"],
    description: "顯示目標工作區與 CLI 執行目錄",
    execute(_args, context) {
      context.addSystemEvent(`Target workspace: ${process.env.DECEPTICON_WORKSPACE_PATH || "/workspace"}\nCLI process directory: ${process.cwd()}`);
    },
  },
];
