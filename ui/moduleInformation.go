package ui

import (
	"cogentcore.org/core/core"
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
)

func ModuleInformatio(tabs *core.Tabs) {
	tab := tabs.NewTab("Информация о модуле")
	tab.Styler(func(s *styles.Style) {

	})

	module := core.NewChooser(tab).SetPlaceholder("Select a fruit").SetItems(MODULE...).SetType(core.ChooserOutlined)
	dateFirmware := core.NewTextField(tab).SetPlaceholder("Дата выхода прошивки")
	moduleVersion := core.NewTextField(tab).SetPlaceholder("Версия модуля")
	mainTask := core.NewTextField(tab).SetPlaceholder("Основание для тестирования (№ задачи redmine)")

}
