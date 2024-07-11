package ui

import (
	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/styles"
)

var (
	MODULE = []core.ChooserItem{
		{Value: "MK-501-022", Tooltip: "Модуль центрального процессора"},
		{Value: "MK-502-142", Tooltip: "Модуль центрального процессора"},
		{Value: "MK-503-120", Tooltip: "Модуль центрального процессора"},
		{Value: "MK-545-002", Tooltip: "Модуль коммуникационный Powerlink"},
		{Value: "MK-531-032A", Tooltip: "Модуль дискретного вывода"},
		{Value: "MK-532-032A", Tooltip: "Модуль дискретного вывода с Namur"},
		{Value: "MK-521-032A", Tooltip: "Модуль дискретного ввода"},
		{Value: "MK-523-032A", Tooltip: "Модуль дискретного ввода с Namur"},
		{Value: "MK-514-008A", Tooltip: "Модуль аналогового вывода"},
		{Value: "MK-513-016A", Tooltip: "Модуль аналогового ввода"},
		{Value: "MK-516-008A", Tooltip: "Модуль аналогового ввода"},
		{Value: "MK-574-008A", Tooltip: "Модуль аналогового вывода с HART"},
		{Value: "MK-576-008A", Tooltip: "Модуль аналогового ввода с HART"},
		{Value: "MK-576-016A", Tooltip: "Модуль аналогового ввода с HART"},
		{Value: "MK-541-002", Tooltip: "Модуль коммуникационный RS-485"},
	}
	moduleBind        *core.Chooser
	dateFirmwareBind  *core.TextField
	moduleVersionBind *core.TextField
	mainTaskBind      *core.TextField
)

type ModuleInformation struct {
	Module                      string
	moduleVersion, dateFirmware string
	mainTask                    string
}

func (this *ModuleInformation) InitUI(groupTabs *core.Tabs) {
	tab := groupTabs.NewTab("Информация о модуле")
	tab.Styler(func(s *styles.Style) {

	})

	// init components
	moduleChooser := core.NewChooser(tab).SetPlaceholder("Select a fruit").SetItems(MODULE...).SetType(core.ChooserOutlined)
	dateFirmwareTextField := core.NewTextField(tab).SetPlaceholder("Дата выхода прошивки")
	moduleVersionTextField := core.NewTextField(tab).SetPlaceholder("Версия модуля")
	mainTaskTextFeild := core.NewTextField(tab).SetPlaceholder("Основание для тестирования (№ задачи redmine)")

	// bind data
	moduleBind = core.Bind(&this.Module, moduleChooser)
	dateFirmwareBind = core.Bind(&this.Module, dateFirmwareTextField)
	moduleVersionBind = core.Bind(&this.Module, moduleVersionTextField)
	mainTaskBind = core.Bind(&this.Module, mainTaskTextFeild)

	// set unfocus func
	moduleBind.OnFocusLost(focusLost)
}

func focusLost(e events.Event) {
	moduleBind.Update()
}
