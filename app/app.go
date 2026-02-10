package app

import (
	"fmt"

	"github.com/MVN-14/devboard-tui/app/command"
	"github.com/MVN-14/devboard-tui/app/list"
	"github.com/MVN-14/devboard-tui/app/screen"
	"github.com/MVN-14/devboard-tui/app/style"
	"github.com/MVN-14/devboard-tui/app/toast"
	"github.com/MVN-14/devboard-tui/app/view"
	tea "github.com/charmbracelet/bubbletea"
)

const widthOffset = 2
const heightOffset = 10

type Model struct {
	errMsg      string
	projectList list.Model
	msg         string
	projectView view.Model
	quitting    bool
	screen      screen.Screen
	toast		toast.Model
	width       int
	height      int
}

func (m Model) Init() tea.Cmd {
	return command.LoadProjects
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.projectList.SetSize(
			m.width-widthOffset,
			m.height-heightOffset,
		)
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			m.quitting = true
			return m, tea.Quit
		}
	case command.LoadMsg:
		cmd = m.projectList.SetItems(msg.Projects)
		cmds = append(cmds, cmd)
	case command.ScreenMsg:
		m.screen = msg.Screen
	case command.ErrMsg:
		m.toast.SetToast(msg.Err.Error(), toast.Error)
	case command.SuccessMsg:
		m.toast.SetToast(msg.Str, toast.Success)
	}
	
	m.toast.Update()

	switch m.screen {
	case screen.ScreenList:
		m.projectList, cmd = m.projectList.Update(msg)
		cmds = append(cmds, cmd)
	case screen.ScreenView:
		m.projectView, cmd = m.projectView.Update(msg)
	}
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if m.quitting {
		return ""
	}

	view := style.RenderTitle(m.width-widthOffset, "D E V B O A R D")
	view += "\n"

	switch {
	case m.errMsg != "":
		view += fmt.Sprintf("%s\n", m.errMsg)
	case m.msg != "":
		view += fmt.Sprintf("\n%s\n", m.msg)
	}

	switch m.screen {
	case screen.ScreenList:
		view += m.projectList.View()
	case screen.ScreenView:
		view += m.projectView.View()
	}
	
	return m.toast.Render(style.RenderView(view), m.width - widthOffset)
}

func New() Model {
	return Model{
		projectList: list.New(),
		projectView: view.New(),
		screen:      screen.ScreenList,
	}
}
