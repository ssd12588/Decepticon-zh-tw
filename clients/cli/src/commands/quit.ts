import type { Command } from "./types.js";

const quit: Command = {
  name: "quit",
  description: "離開 Decepticon CLI",
  aliases: ["exit"],
  execute(_args, ctx) {
    ctx.exit();
  },
};

export default quit;
