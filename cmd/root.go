package cmd

import (
	"log"
	"os"

	"github.com/kolobok-kelbek/commandor/config"
	"github.com/kolobok-kelbek/commandor/execute"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

const env = "COMMANDOR_CONFIG"

var rootCmd = &cobra.Command{
	Use:   "commandor",
	Short: "Command Management System",
	Long:  `Commandor is command management system. This system help for management commands in your project.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MaximumNArgs(1)(cmd, args); err != nil {
			return err
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		configPath, _ := cmd.Flags().GetString("config")
		cmds, err := config.Load(configPath)
		if err != nil {
			log.Fatal(err)
		}

		command := ""
		if len(args) > 0 {
			command = args[0]
		}

		if command != "" {
			for name, data := range cmds {
				if name == command || data.ShortCmd == command {
					execute.Execute(data.Command)
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

			func(name string, data config.Command) {
				list.AddItem(commandName, data.Description, []rune(data.Shortcut)[0], func() {
					app.Stop()
					execute.Execute(data.Command)
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
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("config", "c", os.Getenv(env), "Path for commands file")
}
