import type { Command } from "./types.js";

const definitions = [
  { name: "status", description: "顯示模式、工作階段、連線與 Blue 監控狀態" },
  { name: "doctor", description: "檢視本次工作階段的範圍、RoE、工作流程鎖定、模型與預算（唯讀）" },
  { name: "sessions", description: "尋找此授權測試任務中的對話", argumentHint: "[all | legacy]" },
  { name: "queue", description: "Edit or remove waiting messages", argumentHint: "[clear]" },
  { name: "alerts", description: "檢視 Blue 事件與調查細節", argumentHint: "[incident-id]" },
  { name: "export", description: "匯出目前對話（不覆寫既有檔案）", argumentHint: "[markdown | json] [path]" },
] as const;

// Synchronous argument completion for the fixed-choice local commands.
// The complete contract is sync (commands/types.ts), so dynamic lists
// (incident ids, session names) are not completable here.
const firstTokenCompletions: Partial<Record<(typeof definitions)[number]["name"], (args: string) => readonly string[]>> = {
  sessions: (args) => (args.includes(" ") ? [] : ["all", "legacy"].filter((choice) => choice.startsWith(args))),
  queue: (args) => (args.includes(" ") ? [] : ["clear"].filter((choice) => choice.startsWith(args))),
  export: (args) => (args.includes(" ") ? [] : ["markdown", "json"].filter((choice) => choice.startsWith(args))),
};

export const localCommands: Command[] = definitions.map((definition) => ({
  ...definition,
  ...(firstTokenCompletions[definition.name]
    ? { complete: (args: string) => firstTokenCompletions[definition.name]!(args) }
    : {}),
  execute: (args, context) => context.runLocalCommand?.(definition.name, args),
}));