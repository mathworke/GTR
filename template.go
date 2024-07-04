package main

import (
	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/icons"
	"fmt"
	"os"
	"strconv"
)

func Main() {
	parentWindow := core.NewBody("TITLE")
	core.NewButton(parentWindow).SetText("BUTTON").SetIcon(icons.Add).OnClick(btnClick)

	on := true
	core.Bind(&on, core.NewSwitch(parentWindow)).OnChange(func(e events.Event) {
		core.MessageSnackbar(parentWindow, "Value >> "+strconv.FormatBool(on))
	})

	go func() {
		for {
			if !on {
				fmt.Fprint(os.Stdout, "break\n")
				break
			}

			fmt.Fprint(os.Stdout, "asdasd\n")
		}
	}()

	parentWindow.RunMainWindow()
}

func btnClick(event events.Event) {
	fmt.Fprint(os.Stdout, "btn click")
}
