package ui

import (
	l "GTR/assets"

	"cogentcore.org/core/core"
	// "cogentcore.org/core/events"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/units"
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
	ModuleVersion, DateFirmware string
	MainTask                    string
}

func (m *ModuleInformation) InitUI(groupTabs *core.Tabs, logger *l.Logger) {
	tab := groupTabs.NewTab("Module")
	tab.Styler(func(s *styles.Style) {

	})

	/* init components

	Initialization of necessary fields for their further rendering and information collection
	Инициализация необходимых полей для их последующей визуализации и сбора информации
	*/
	logger.LogIngo(" Start initialization module informtion components\n")

	moduleChooser := core.NewChooser(tab).SetPlaceholder("Выберите модуль").SetItems(MODULE...).SetType(core.ChooserOutlined)
	dateFirmwareTextField := core.NewTextField(tab).SetPlaceholder("Дата выхода прошивки")
	moduleVersionTextField := core.NewTextField(tab).SetPlaceholder("Версия модуля")
	mainTaskTextFeild := core.NewTextField(tab).SetPlaceholder("Основание для тестирования (№ задачи redmine)")

	logger.LogIngo(" Module information components initialization\n")

	// style components
	moduleChooser.Styler(func(s *styles.Style) {

	})
	dateFirmwareTextField.Styler(func(s *styles.Style) {
		s.Gap.Y = units.Px(200)
	})
	moduleVersionTextField.Styler(func(s *styles.Style) {
		s.Gap = units.XY{X: units.Px(50), Y: units.Px(50)}
	})
	mainTaskTextFeild.Styler(func(s *styles.Style) {
		s.Min = units.XY{Y: units.Px(25)}
	})

	/* bind data

	Binding of fields to a variable for systematic retrieval of information from a component
	Привязка полей к переменной для систематизированного получения информации с компонента
	*/
	logger.LogIngo(" Start bind data module information\n")

	moduleBind = core.Bind(&m.Module, moduleChooser)
	dateFirmwareBind = core.Bind(&m.DateFirmware, dateFirmwareTextField)
	moduleVersionBind = core.Bind(&m.ModuleVersion, moduleVersionTextField)
	mainTaskBind = core.Bind(&m.MainTask, mainTaskTextFeild)

	logger.LogIngo(" Module information data binded\n")

}

// maybe have most perfect algorithms... maybe using interface
