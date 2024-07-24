package main

import (
	"GTR/assets"
	"GTR/ui"
	"flag"

	"cogentcore.org/core/core"
)

var debug = flag.Bool("debug", false, "run in debug mode")

func main() {
	flag.Parse()
	logger := assets.NewLogger(*debug)

	mainWindow := core.NewBody("GTR")

	groupTabs := core.NewTabs(mainWindow)

	// init ui
	module := &ui.ModuleInformation{}
	testCase := &ui.SelectTestCase{}

	module.InitUI(groupTabs, logger)
	testCase.InitUI(groupTabs, logger, module)

	mainWindow.RunMainWindow()
}
