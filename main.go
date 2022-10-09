package main

import (
	"github.com/rivo/tview"
	"log"
	"os"
)

func main() {
	cmds, err := configLoad()
	if err != nil {
		log.Fatal(err)
	}

	cmd := ""
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	if cmd != "" {
		for name, data := range cmds {
			if name == cmd || data.ShortCmd == cmd {
				execute(data.Command)
				os.Exit(0)
			}
		}

		log.Fatal("command not found")
	}

	app := tview.NewApplication()
	list := tview.NewList()

	for name, data := range cmds {
		commandName := data.Title
		if commandName == "" {
			commandName = name
		}

		func(name string, data command) {
			list.AddItem(commandName, data.Description, []rune(data.Shortcut)[0], func() {
				app.Stop()
				execute(data.Command)
			})
		}(name, data)
	}

	list.AddItem("Quit", "Press to exit", 'q', func() {
		app.Stop()
	})

	app.SetRoot(list, true).SetFocus(list)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
