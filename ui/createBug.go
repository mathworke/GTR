package ui

import (
	helper "github.com/mathworke/GTR/assets"

	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/icons"
)

var (
	tabBug  *core.Frame
	bugList *core.Table
)

type CreateBug struct {
	Bugs []bugField
}

type bugField struct {
	Номер          string
	Название       string
	Исправленность bool
	Серьёзность    string
}

func (c *CreateBug) InitUI(groupTabs *core.Tabs, logger *helper.Logger) {
	logger.LogIngo("Start initialization createBug components\n")
	tabBug = groupTabs.NewTab("Дефекты", icons.Error)

	/* init components

	Initialization of necessary fields for their further rendering and information collection
	Инициализация необходимых полей для их последующей визуализации и сбора информации
	*/
	btnBugFrame := core.NewFrame(tabBug)
	addBug := core.NewButton(btnBugFrame).SetText("Добавить дефект")
	addBug.SetTooltip("Добавить дополнительное поле для ввода дефекта.")

	resetBug := core.NewButton(btnBugFrame).SetIcon(icons.Reset)
	resetBug.SetTooltip("Сбросить поля дефектов.")

	bugs := make([]*bugField, 0)
	bugList = core.NewTable(tabBug).SetSlice(&bugs)

	logger.LogIngo("components initialozation")

	// click button command
	addBug.OnClick(func(e events.Event) {
		logger.LogIngo("func createBug.addBug.OnClick() starting")
		addBugField(logger)
	})

	resetBug.OnClick(func(e events.Event) {
		logger.LogIngo("func createBug.reset,OnClick() starting")
		cur := bugList.Slice.(*[]*bugField)
		*cur = make([]*bugField, 0)
		bugList.Update()
	})
}

func addBugField(logger *helper.Logger) {
	logger.LogIngo("createBug.addBugField() start")
	cur := bugList.Slice.(*[]*bugField)
	adderBug := newBugField()
	*cur = append(*cur, adderBug)
	bugList.Update()
}

func newBugField() *bugField {
	return &bugField{}
}
