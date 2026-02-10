package command

import (
	"github.com/MVN-14/devboard-tui/app/screen"
	"github.com/MVN-14/devboard-tui/devboard"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	LoadMsg struct {
		Projects []devboard.Project
	}

	EditMsg struct {
		Project devboard.Project
	}

	AddMsg struct{}

	ScreenMsg struct {
		Screen screen.Screen
	}

	ErrMsg struct {
		Err error
	}

	SuccessMsg struct {
		Str string
	}
)

func SendScreenMsg(s screen.Screen) tea.Cmd {
	return func() tea.Msg { return ScreenMsg{s} }
}

func SendAddMsg() tea.Cmd {
	return func() tea.Msg { return AddMsg{} }
}

func SendEditMsg(p devboard.Project) tea.Cmd {
	return func() tea.Msg { return EditMsg{p} }
}

func SendSuccessMsg(m string) tea.Cmd {
	return func() tea.Msg { return SuccessMsg{m} }
}

func SendErrMsg(err error) tea.Cmd {
	return func() tea.Msg { return ErrMsg{err} }
}
