package console

import "testing"

func TestConsole(t *testing.T) {
	console := NewColorConsole()
	console.Info("info", "info")
	console.InfoF("%v", "info")

	console.Debug("debug", "debug")
	console.DebugF("%v", "debug")

	console.Warning("warning", "warning")
	console.WarningF("%v", "warning")

	console.Error("error")
	console.ErrorF("%v", "error")
}
