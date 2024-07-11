package main

import (
	"GTR/ui"
	"cogentcore.org/core/core"
	"fmt"
)

func main() {
	mainWindow := core.NewBody("GTR")
	groupTabs := core.NewTabs(mainWindow)

	// init ui
	module := ui.ModuleInformation{}
	module.InitUI(groupTabs)

	go func() {
		for {
			fmt.Println(module.Module)
		}
	}()

	mainWindow.RunMainWindow()
}
