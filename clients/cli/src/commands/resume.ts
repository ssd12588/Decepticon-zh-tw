import type { Command } from "./types.js";

const resume: Command = {
  name: "resume",
  description: "選擇已儲存的工作階段並接續對話",
  aliases: ["r", "continue"],
  argumentHint: "[thread-id]",
  execute(args, ctx) {
    if (args.trim()) ctx.restoreSession?.(args.trim());
    else return ctx.runLocalCommand?.("sessions", "current");
  },
};

export default resume;
