//go:build windows

package platform

import "golang.org/x/sys/windows"

// EnsureUTF8Console switches the attached Windows console to UTF-8 (code page
// 65001) at startup. Traditional Chinese UI text is otherwise mangled by the
// legacy CP950/Big5 code page that zh-TW Windows installs default to: Go always
// writes UTF-8 bytes, while the console decodes them using its active code page.
//
// Errors are ignored on purpose — when stdout is redirected there is no console
// to reconfigure and the call simply fails without affecting the command.
func EnsureUTF8Console() {
	_ = windows.SetConsoleOutputCP(65001)
	_ = windows.SetConsoleCP(65001)
}
