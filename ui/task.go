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

	reset := core.NewButton(btnFrame).SetIcon(icons.Reset)
	reset.SetTooltip("Сбросить поля дефектов")

	sl := make([]*Change, 0)
	// tmp := Change{Bug: ""}
	// sl[0] = &tmp
	changedList = core.NewTable(tabTask).SetSlice(&sl)

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

	// click button command
	addChanged.OnClick(func(e events.Event) {
		addChangedField()
		t.FillChanged()
	})

	reset.OnClick(func(e events.Event) {
		cur := changedList.Slice.(*[]*Change)
		*cur = make([]*Change, 0)
		changedList.Update()
	})
}

func (t *Task) FillChanged() {
	t.Changed = make([]string, 0)
	for _, v := range *changedList.Slice.(*[]*Change) {
		t.Changed = append(t.Changed, v.Bug)
	}
}

func addChangedField() {
	cur := changedList.Slice.(*[]*Change)
	tmp := Change{Bug: ""}
	*cur = append(*cur, &tmp)
	changedList.Update()
}
