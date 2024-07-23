package main

import (
	"GTR/assets"
	"GTR/ui"

	"cogentcore.org/core/core"
)

func main() {
	logger := assets.NewLogger(true)

	mainWindow := core.NewBody("GTR")

	groupTabs := core.NewTabs(mainWindow)

	// init ui
	module := &ui.ModuleInformation{}
	module.InitUI(groupTabs, *logger)

	mainWindow.RunMainWindow()
}
