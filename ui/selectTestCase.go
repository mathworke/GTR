package ui

import (
	l "GTR/assets"

	"cogentcore.org/core/core"
	// "cogentcore.org/core/events"
	"cogentcore.org/core/styles"
)

var (
	
)

type SelectTestCase struct {
	Module           string
	CompleteTestCase [][]string
}

func (s *SelectTestCase) InitUI(groupTabs *core.Tabs, logger *l.Logger) {
	tab := groupTabs.NewTab("Информация о модуле")
	tab.Styler(func(s *styles.Style) {

	})

	/* init components

	Initialization of necessary fields for their further rendering and information collection
	Инициализация необходимых полей для их последующей визуализации и сбора информации
	*/
	logger.LogIngo(" Start initialization components\n")
	moduleChooser := core.NewChooser(tab).SetPlaceholder("Выберите модуль").SetItems(MODULE...).SetType(core.ChooserOutlined)
	dateFirmwareTextField := core.NewTextField(tab).SetPlaceholder("Дата выхода прошивки")
	moduleVersionTextField := core.NewTextField(tab).SetPlaceholder("Версия модуля")
	mainTaskTextFeild := core.NewTextField(tab).SetPlaceholder("Основание для тестирования (№ задачи redmine)")
	logger.LogIngo(" Components initialization\n")

	/* bind data

	Binding of fields to a variable for systematic retrieval of information from a component
	Привязка полей к переменной для систематизированного получения информации с компонента
	*/
	moduleBind = core.Bind(&m.Module, moduleChooser)
	dateFirmwareBind = core.Bind(&m.Module, dateFirmwareTextField)
	moduleVersionBind = core.Bind(&m.Module, moduleVersionTextField)
	mainTaskBind = core.Bind(&m.Module, mainTaskTextFeild)

	/* set unfocus func

	!!!WFT!!! replace on most ef algorithms

	*/
	moduleBind.OnFocusLost(focusLostModule)
	dateFirmwareBind.OnFocusLost(focusLostDateFirmware)
	moduleVersionBind.OnFocusLost(focusLostModuleVersion)
	mainTaskBind.OnFocusLost(focusLostMainTask)
}

func (s *SelectTestCase) RenderTestCase() {

}
