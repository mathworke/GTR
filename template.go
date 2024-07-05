package main

import (
	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/styles"
	"fmt"
)

func Main() {
	mainWindow := core.NewBody("MAIN")
	ts := core.NewTabs(mainWindow)

	home(ts)
	task(ts)

	mainWindow.RunMainWindow()
}

var str string
var test *core.TextField

func home(ts *core.Tabs) {
	tab := ts.NewTab("Главная", icons.Home)
	tab.Styler(func(s *styles.Style) {
		s.CenterAll()
	})

	core.NewTextField(tab).SetPlaceholder("Название модуля")
	core.NewTextField(tab).SetPlaceholder("Дата тестирования")

	test = core.Bind(&str, core.NewTextField(tab))
	test.SetPlaceholder("test bind data").OnFocus(foc)
	test.OnFocusLost(unfoc)

	core.NewTextField(tab).SetType(core.TextFieldOutlined).SetPlaceholder("Название модуля").AddClearButton()
	core.NewTextField(tab).SetType(core.TextFieldOutlined).SetPlaceholder("Дата тестирования").AddClearButton()
}

func foc(e events.Event) {
	str = "on foc"
	test.Update()
}

func unfoc(e events.Event) {
	str = test.Text()
	test.Update()
}

func task(ts *core.Tabs) {
	tab := ts.NewTab("Задачи", icons.Task)
	tab.Styler(func(s *styles.Style) {
		s.CenterAll()
	})

	core.NewButton(tab).OnClick(func(e events.Event) {
		fmt.Println(str)
	})
}
