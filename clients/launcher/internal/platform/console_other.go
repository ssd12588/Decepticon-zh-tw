//go:build !windows

package platform

// EnsureUTF8Console is a no-op outside Windows: macOS and Linux terminals are
// already UTF-8, so no code-page switch is needed.
func EnsureUTF8Console() {}
