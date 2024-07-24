package ui

import (
	l "GTR/assets" // logger

	"cogentcore.org/core/core"
	// "cogentcore.org/core/events" 
	"cogentcore.org/core/styles"
)

var ()

type SelectTestCase struct {
	Module           string
	CompleteTestCase [][]string
}

func (s *SelectTestCase) InitUI(groupTabs *core.Tabs, logger *l.Logger, args ...interface{}) {
	tab := groupTabs.NewTab("Test-case")
	tab.Styler(func(s *styles.Style) {

	})

	/* init components

	Initialization of necessary fields for their further rendering and information collection
	Инициализация необходимых полей для их последующей визуализации и сбора информации
	*/
	logger.LogIngo(" Start initialization components\n")
	_ = core.NewText(tab).SetType(core.TextDisplayLarge).SetText("asd")
	logger.LogIngo(" Components initialization\n")

	/* bind data

	Binding of fields to a variable for systematic retrieval of information from a component
	Привязка полей к переменной для систематизированного получения информации с компонента
	*/
	// moduleBind = core.Bind(&m.Module, moduleChooser)

	/* set unfocus func

	!!!WFT!!! replace on most ef algorithms

	*/


}

func (s *SelectTestCase) renderTestCase() {
	println(0x1A8)
}
