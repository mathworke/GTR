package ui

import (
	helper "GTR/assets" // logger
	"errors"

	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/styles"
)

var (
	testCaseTable *core.Table
	tab           *core.Frame
)

const (
	CURRENT_MODULE = "Тест-кейсы загружены для модуля: "
)

type SelectTestCase struct {
	Module           string
	CompleteTestCase [][]testCaseStruct // 0 - test passed ; 1 - test failed
}

type testCaseStruct struct {
	Done     bool   `width:"25"`
	TestCase string `width:"25"`
	Tester   string `width:"25"`
	Comment  string `width:"25"`
}

func (s *SelectTestCase) InitUI(groupTabs *core.Tabs, xmlConf *helper.Modules, logger *helper.Logger, args ...interface{}) {
	tab = groupTabs.NewTab("Test-case", icons.Settings)
	tab.Styler(func(s *styles.Style) {

	})

	/* init components

	Initialization of necessary fields for their further rendering and information collection
	Инициализация необходимых полей для их последующей визуализации и сбора информации
	*/
	logger.LogIngo("Start initialization components\n")

	module := core.NewText(tab).SetType(core.TextDisplayMedium)
	currentModule := core.NewText(tab).SetType(core.TextBodySmall)
	refreshTestCase := core.NewButton(tab).SetText("Обновить...")
	tbl := make([]*testCaseStruct, 0)
	testCaseTable = core.NewTable(tab).SetSlice(&tbl)

	/* bind data

	Binding of fields to a variable for systematic retrieval of information from a component
	Привязка полей к переменной для систематизированного получения информации с компонента
	*/
	moduleBind := &args[0].(*ModuleInformation).Module
	core.Bind(moduleBind, module)

	logger.LogIngo("bind data complete")

	// click button command
	refreshTestCase.OnClick(func(e events.Event) {
		s.Module = module.Text
		currentModule.SetText(CURRENT_MODULE + s.Module).Update()
		s.RenderTestCase(xmlConf, logger)
	})

	logger.LogIngo("Components initialization\n")
}

/* render test case table
 */
func (s *SelectTestCase) RenderTestCase(xmlConf *helper.Modules, logger *helper.Logger) {
	logger.LogIngo("rendering test-case...")

	id, err := getIdModule(s.Module, xmlConf)
	if err != nil {
		logger.LogError(err.Error())
		core.MessageSnackbar(tab, "Выбранный модуль не добавлен в конфигурацию.")
		clearSlice()
		return
	}

	if len(xmlConf.Modules[id].Tests.TestCase) == 0 {
		logger.LogError("not add test-case for %v", s.Module)
		core.MessageSnackbar(tab, "Для выбранного модуля не добавлены тест-кейсы в конфигурационный файл.")
		clearSlice()
		return
	}

	tbl := make([]*testCaseStruct, len(xmlConf.Modules[id].Tests.TestCase))

	for i := range tbl {
		logger.LogIngo("create test-case #%v for %v", xmlConf.Modules[id].Tests.TestCase[i].Id, xmlConf.Modules[id].Type)
		ts := newTestCase(xmlConf.Modules[id].Tests.TestCase[i].Id)
		tbl[i] = ts
	}

	testCaseTable.SetSlice(&tbl).UpdateSliceSize()
	logger.LogIngo("rendered test-case is complete")
}

func clearSlice() {
	tbl := make([]*testCaseStruct, 0)
	testCaseTable.SetSlice(&tbl).UpdateSliceSize()
}

func newTestCase(name string) (testCase *testCaseStruct) {
	return &testCaseStruct{
		Done:     false,
		TestCase: name,
		Tester:   "",
		Comment:  "",
	}
}

func getIdModule(module string, xmlConf *helper.Modules) (int, error) {
	for index, item := range xmlConf.Modules {
		if item.Type == module {
			return index, nil
		}
	}

	return -1, errors.New("not found source module")
}

func (t *SelectTestCase) GetTestCase(logger *helper.Logger) {
	logger.LogIngo("start GetTestCase()")
	s := testCaseTable.Slice.(*[]*testCaseStruct)
	t.CompleteTestCase = make([][]testCaseStruct, 2)
	t.CompleteTestCase[0] = make([]testCaseStruct, 0)
	t.CompleteTestCase[1] = make([]testCaseStruct, 0)
	for _, v := range *s {
		logger.LogIngo("processing test-case %v...", v)
		if v.Done {
			t.CompleteTestCase[0] = append(t.CompleteTestCase[0], *v)
		} else {
			t.CompleteTestCase[1] = append(t.CompleteTestCase[1], *v)
		}
	}
	logger.LogIngo("end GetTestCase()")
}
