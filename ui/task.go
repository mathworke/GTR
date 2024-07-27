package ui

import (
	helper "GTR/assets"

	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/units"
)

var (
	tabTask      *core.Frame
	mainTaskBind *core.TextField
	changedList  *core.Table
)

type Task struct {
	MainTask string
	Changed  []string
}

type Change struct {
	Bug string
}

func (t *Task) InitUI(groupTabs *core.Tabs, logger *helper.Logger) {
	logger.LogIngo("Start initialization task components\n")
	tabTask = groupTabs.NewTab("Задача", icons.Task)

	/* init components

	Initialization of necessary fields for their further rendering and information collection
	Инициализация необходимых полей для их последующей визуализации и сбора информации
	*/
	t.Changed = make([]string, 0)

	mainTaskTextField := core.NewTextField(tabTask).SetPlaceholder("Основание для тестирования (№ задачи на Redmine)")
	btnFrame := core.NewFrame(tabTask)
	addChanged := core.NewButton(btnFrame).SetText("Добавить изменение")
	addChanged.SetTooltip("Добавить дополнительное поле для ввода изменения.")

	resetTask := core.NewButton(btnFrame).SetIcon(icons.Reset)
	resetTask.SetTooltip("Сбросить поля изменений.")

	sl := make([]*Change, 0)
	changedList = core.NewTable(tabTask).SetSlice(&sl)

	logger.LogIngo("components initialozation")

	/* styling components

	 */
	mainTaskTextField.Styler(func(s *styles.Style) {
		s.Min = units.XY{Y: units.Px(25)}
	})

	/* bind data

	Binding of fields to a variable for systematic retrieval of information from a component
	Привязка полей к переменной для систематизированного получения информации с компонента
	*/
	mainTaskBind = core.Bind(&t.MainTask, mainTaskTextField)

	logger.LogIngo("data binded")

	// click button command
	addChanged.OnClick(func(e events.Event) {
		logger.LogIngo("func task.addChanged.OnClick() starting")
		addChangedField(logger)
		t.FillChanged(logger)
	})

	resetTask.OnClick(func(e events.Event) {
		logger.LogIngo("func task.reset.OnClick() starting")
		cur := changedList.Slice.(*[]*Change)
		*cur = make([]*Change, 0)
		changedList.Update()
	})
}

func (t *Task) FillChanged(logger *helper.Logger) {
	logger.LogIngo("starting FillChanged")
	t.Changed = make([]string, 0)
	for _, v := range *changedList.Slice.(*[]*Change) {
		t.Changed = append(t.Changed, v.Bug)
	}
}

func addChangedField(logger *helper.Logger) {
	logger.LogIngo("task.addChangedField() start")
	cur := changedList.Slice.(*[]*Change)
	tmp := Change{Bug: ""}
	*cur = append(*cur, &tmp)
	changedList.Update()
}
