package ui

import (
	l "GTR/assets" // logger
	"fmt"

	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/styles"
)

var (
	testCaseTable *core.Table
	tab           *core.Frame
)

const (
	CURRENT_MODULE = "Текущие тест-кейсы для модуля: "
)

type SelectTestCase struct {
	Module           string
	CompleteTestCase [][]string
}

type testCaseStruct struct { //types:add

	// test case status
	Done bool `width:"25"`

	// title test case
	TestCase string `width:"25"`

	// a tester
	Tester string `width:"25"`

	// comment for test-case
	Comment string `width:"25"`
}

func (s *SelectTestCase) InitUI(groupTabs *core.Tabs, logger *l.Logger, args ...interface{}) {
	tab = groupTabs.NewTab("Test-case")
	tab.Styler(func(s *styles.Style) {

	})

	
	/* init components

	Initialization of necessary fields for their further rendering and information collection
	Инициализация необходимых полей для их последующей визуализации и сбора информации
	*/
	logger.LogIngo(" Start initialization components\n")

	module := core.NewText(tab).SetType(core.TextDisplayMedium)
	currentModule := core.NewText(tab).SetType(core.TextBodySmall)
	refreshTestCase := core.NewButton(tab).SetText("Обновить...")
	tbl := make([]*testCaseStruct, 0)
	testCaseTable = core.NewTable(tab).SetSlice(&tbl)

	logger.LogIngo(" Components initialization\n")


	/* bind data

	Binding of fields to a variable for systematic retrieval of information from a component
	Привязка полей к переменной для систематизированного получения информации с компонента
	*/
	moduleBind := &args[0].(*ModuleInformation).Module
	core.Bind(moduleBind, module)


	// click button command
	refreshTestCase.OnClick(func(e events.Event) {
		s.Module = module.Text
		currentModule.SetText(CURRENT_MODULE + s.Module).Update()
		s.RenderTestCase(logger)
	})
}

func getTestCase() {

}

/* render test case table
*/
func (s *SelectTestCase) RenderTestCase(logger *l.Logger) {
	logger.LogIngo(" rendering test-case...")

	tbl := make([]*testCaseStruct, 50)
	for i := range tbl {
		ts := &testCaseStruct{TestCase: fmt.Sprintf("1-12%v", i)}
		tbl[i] = ts
	}
	testCaseTable.SetSlice(&tbl)

	logger.LogIngo(" rendered test-case is complete")
}
