package devboard

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ProjectList struct {
	list list.Model
}
func NewProjectList(width, height int) ProjectList {
	return ProjectList{
		list: list.New([]list.Item{}, list.NewDefaultDelegate(), width, height),
	}
}

func (pl *ProjectList) SetProjects(projects []*Project) tea.Cmd {
	items := []list.Item{}
	for _, v := range projects {
		items = append(items, v)
	}
	cmd := pl.list.SetItems(items)
	return cmd
}

func (pl *ProjectList) Update(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	pl.list, cmd = pl.list.Update(msg)
	return cmd
}

func (pl ProjectList) View() string {
	return pl.list.View()
}

func (pl ProjectList) Items() []list.Item {
	return pl.list.Items()
}
