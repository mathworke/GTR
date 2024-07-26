package main

import (
	"GTR/assets"
	"GTR/ui"
	"flag"

	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/icons"
)

const (
	PATH_CONFIG = "config/module.xml"
	PATH_PERSON = "config/person.xml"

	TITLE = "GTR - Генератор отчётов тестирования"
)

var debug = flag.Bool("debug", false, "run in debug mode")

func main() {
	flag.Parse()
	logger := assets.NewLogger(*debug)

	conf := new(assets.Modules)
	conf.GetConfiguration(PATH_CONFIG, logger)
	person := new(assets.Person)
	person.GetPerson(PATH_PERSON, logger)

	// init root
	mainWindow := core.NewBody(TITLE)

	groupTabs := core.NewTabs(mainWindow)

	// init ui
	module := &ui.ModuleInformation{}
	testCase := &ui.SelectTestCase{}
	task := &ui.Task{}

	module.InitUI(groupTabs, conf, logger)
	task.InitUI(groupTabs, logger)
	testCase.InitUI(groupTabs, conf, logger, module)

	// configuration button
	settings := core.NewFrame(mainWindow)
	core.NewButton(settings).SetIcon(icons.Settings).SetTooltip("Обновить конфигурационный файл модулей").OnClick(func(e events.Event) {
		conf.GetConfiguration(PATH_CONFIG, logger)
		logger.LogPrint("Update configuration file for module")
	})
	_ = core.NewSwitches(settings).SetType(core.SwitchChip).SetStrings("Релиз")

	generate := core.NewButton(settings).SetText("Создать отчёт")
	generate.OnClick(func(e events.Event) {
		comment := core.NewBody().AddTitle("Комментарий к тестированию").AddText("Коммантарий")
		commentTextFiels := core.NewTextField(comment)
		comment.AddBottomBar(func(parent core.Widget) {
			comment.AddCancel(parent)
			comment.AddOK(parent).OnClick(func(e events.Event) {
				core.MessageSnackbar(generate, commentTextFiels.Text())
			})
		})
		comment.RunDialog(generate)
	})

	// core.NewButton(settings).SetIcon(icons.Start).SetText(
	// 	"Создать отчёт",
	// ).SetTooltip(
	// 	"Сгенерировать отчёт по тестированию выбранно модуля.",
	// ).OnClick(func(e events.Event) {
	// 	logger.LogPrint("GENERATE REPORT")
	// })

	// start
	mainWindow.RunMainWindow()
}
