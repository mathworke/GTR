package ui

import (
	helper "GTR/assets"

	"cogentcore.org/core/core"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/units"
)

var (
	MODULE = make([]core.ChooserItem, 0)

/* unused variable
moduleBind        *core.Chooser
dateFirmwareBind  *core.TextField
moduleVersionBind *core.TextField
useProjectBind    *core.TextField
dateTestingBind   *core.TextField
*/
)

type ModuleInformation struct {
	Module                      string
	ModuleVersion, DateFirmware string
	DateTesting                 string
	UseProject                  string
}

func (m *ModuleInformation) InitUI(groupTabs *core.Tabs, xmlConf *helper.Modules, logger *helper.Logger) {
	m.loadData(xmlConf)

	tab := groupTabs.NewTab("Модуль", icons.Info)
	tab.Styler(func(s *styles.Style) {

	})

	/* init components

	Initialization of necessary fields for their further rendering and information collection
	Инициализация необходимых полей для их последующей визуализации и сбора информации
	*/
	logger.LogIngo("Start initialization module information components\n")

	moduleChooser := core.NewChooser(tab).SetPlaceholder("Выберите модуль").SetItems(MODULE...).SetType(core.ChooserOutlined)
	dateTestingTextField := core.NewTextField(tab).SetPlaceholder("Дата тестирования")
	dateFirmwareTextField := core.NewTextField(tab).SetPlaceholder("Дата выхода прошивки")
	moduleVersionTextField := core.NewTextField(tab).SetPlaceholder("Версия модуля")
	useProjectTextField := core.NewTextField(tab).SetPlaceholder("Использованный проект")

	logger.LogIngo("Module information components initialization\n")

	// style components
	moduleChooser.Styler(func(s *styles.Style) {

	})
	dateFirmwareTextField.Styler(func(s *styles.Style) {
		s.Gap.Y = units.Px(200)
	})
	moduleVersionTextField.Styler(func(s *styles.Style) {
		s.Gap = units.XY{X: units.Px(50), Y: units.Px(50)}
	})

	/* bind data

	Binding of fields to a variable for systematic retrieval of information from a component
	Привязка полей к переменной для систематизированного получения информации с компонента
	*/
	logger.LogIngo("Start bind data module information\n")

	_ = core.Bind(&m.Module, moduleChooser)                 // moduleBind
	_ = core.Bind(&m.DateTesting, dateTestingTextField)     // dateTestingBind
	_ = core.Bind(&m.DateFirmware, dateFirmwareTextField)   // dateFirmwareBind
	_ = core.Bind(&m.ModuleVersion, moduleVersionTextField) // moduleVersionBind
	_ = core.Bind(&m.UseProject, useProjectTextField)       // useProjectBind

	logger.LogIngo("Module information data binded\n")

}

func (m *ModuleInformation) loadData(xmlConf *helper.Modules) {
	for i := 0; i < len(xmlConf.Modules); i++ {
		MODULE = append(MODULE, core.ChooserItem{Value: xmlConf.Modules[i].Type, Tooltip: xmlConf.Modules[i].Description})
	}
}

// maybe have most perfect algorithms... maybe using interface
