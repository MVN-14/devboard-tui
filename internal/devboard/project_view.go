package devboard

import (
	tea "github.com/charmbracelet/bubbletea"
)

type ProjectView struct {
	project Project
	nameInput textinput.Model
	pathInput textinput.Model
}

func (pv ProjectView) View() string {
	return ""
}

func (pv Update) Update(msg tea.Msg) tea.Cmd {

}
