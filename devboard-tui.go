package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type App struct {
	Choices []string
	Cursor int
	Selected map[int]struct{}
}

func (app App) Init() tea.Cmd {
	return nil
}

func (app App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl-c", "q":
			return app, tea.Quit
		case "up", "k":
			if app.Cursor > 0 {
				app.Cursor--
			}
		case "down", "j":
			if app.Cursor < len(app.Choices) - 1 {
				app.Cursor++
			}
		case "enter", " ":
			_, exists := app.Selected[app.Cursor]
			if exists {
				delete(app.Selected, app.Cursor)
			} else {
				app.Selected[app.Cursor] = struct{}{}
			}
		}
	}
	return app, nil
}

func (app App) View() string {
	var strBuilder strings.Builder
	strBuilder.WriteString("Devboard.\n\n")
	
	for i, choice := range app.Choices {
		cursor := " "
		if app.Cursor == i {
			cursor = ">"
		}

		checkbox := " "
		if _, exists := app.Selected[i]; exists {
			checkbox = "x"
		}

		choiceStr := fmt.Sprintf("%s [%s] - %s\n", cursor, checkbox, choice)
		strBuilder.WriteString(choiceStr)
	}
	
	strBuilder.WriteString("\n\nq to quit.")
	return strBuilder.String()
}

func main() {
	program := tea.NewProgram(App{
		Choices: []string{"List your projects", "Add a project", "Update a project", "Remove a project"},
		Selected: make(map[int]struct{}),
	})
	
	_, err := program.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
