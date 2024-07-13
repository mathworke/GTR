package main

import (
	"GTR/ui"
	"cogentcore.org/core/core"
)

func main() {
	mainWindow := core.NewBody("GTR")
	groupTabs := core.NewTabs(mainWindow)

	// init ui
	module := ui.ModuleInformation{}
	module.InitUI(groupTabs)

	mainWindow.RunMainWindow()
}
