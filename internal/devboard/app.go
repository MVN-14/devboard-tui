package devboard

import (
	"strings"

	"github.com/charmbracelet/bubbles/list"
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
	projectList  list.Model
	projectTitle textinput.Model
}

func NewApp() app {
	return app{
		screenState: stateList,
		projects:    []Project{},
		projectList: list.New([]list.Item{}, list.NewDefaultDelegate(), 80, 20),
	}
}

func (app app) Init() tea.Cmd {
	return loadProjects
}

func (app app) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case errMsg:
		app.error = msg.Error()
	case projectsMsg:
		var items []list.Item
		for _, v := range msg {
			items = append(items, v)
		}
		cmds = append(cmds, app.projectList.SetItems(items))
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return app, tea.Quit
		case "a":
			app.screenState = stateView
		}
	}

	var cmd tea.Cmd
	app.projectList, cmd = app.projectList.Update(msg)
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
