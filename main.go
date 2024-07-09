package main

import "cogentcore.org/core/core"

func main() {
	mainWindow := core.NewBody("GTR")
	groupTabs := core.NewTabs(mainWindow)

	mainWindow.RunMainWindow()
}
