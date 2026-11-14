package main

import (
	"github.com/BitterSecurity/Decepticon/clients/launcher/cmd"
	"github.com/BitterSecurity/Decepticon/clients/launcher/internal/platform"
)

func main() {
	// UTF-8 console before any output: without it, zh-TW Windows (CP950) renders
	// the Traditional Chinese UI text as mojibake.
	platform.EnsureUTF8Console()
	cmd.Execute()
}
