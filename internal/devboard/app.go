package devboard

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type state uint8

const (
	stateList state = iota
	stateView
)

type app struct {
	cursor       uint8
	error        string
	screenState  state
	projects     []Project
	projectList  ProjectList
	projectTitle textinput.Model
}

func NewApp() app {
	return app{
		screenState: stateList,
		projects:    []Project{},
	}
}

func (app app) Init() tea.Cmd {
	return loadProjects
}

func (app app) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		app.projectList = NewProjectList(msg.Width, msg.Height - 10)
	case errMsg:
		app.error = msg.Error()
	case projectsMsg:
		cmds = append(cmds, app.projectList.SetProjects(msg))
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return app, tea.Quit
		case "a":
			app.screenState = stateView
		}
	}

	cmd := app.projectList.Update(msg)
	cmds = append(cmds, cmd)

	return app, tea.Batch(cmds...)
}

func (app app) View() string {
	sb := strings.Builder{}
	sb.WriteString(titleStyle.Render("Devboard TUI") + "\n")

	if app.error != "" {
		sb.WriteString(errorStyle.Render("An error occured:\n"+app.error) + "\n")
	} else {
		switch app.screenState {
		case stateList:
			if len(app.projectList.Items()) == 0 {
				sb.WriteString("\nNo projects are being tracked by devboard.\n")
			} else {
				sb.WriteString(app.projectList.View())
			}
		case stateView:
		}
	}

	return viewStyle.Render(sb.String())
}
