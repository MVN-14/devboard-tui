package command

import (
	"github.com/MVN-14/devboard-tui/app/screen"
	"github.com/MVN-14/devboard-tui/devboard"
	tea "github.com/charmbracelet/bubbletea"
)

func LoadProjects() tea.Msg {
	projects, err := devboard.LoadProjects()
	if err != nil {
		return ErrMsg{err}
	}
	return LoadMsg{projects}
}

func OpenProject(id int) tea.Cmd {
	err := devboard.OpenProject(id)
	if err != nil {
		return SendErrMsg(err)
	}
	return SendSuccessMsg("Project opened successfully")
}

func DeleteProject(id int) tea.Cmd {
	err := devboard.DeleteProject(id)
	if err != nil {
		return SendErrMsg(err)
	}
	return tea.Batch(
		SendSuccessMsg("Project deleted successfully"),
		LoadProjects,
	)

}

func UpdateProject(p devboard.Project) tea.Cmd {
	err := devboard.UpdateProject(p)
	if err != nil {
		return SendErrMsg(err)
	}

	return tea.Batch(
		SendSuccessMsg(p.Name+" updated successfully"),
		SendScreenMsg(screen.ScreenList),
		LoadProjects)
}

func AddProject(p devboard.Project) tea.Cmd {
	err := devboard.AddProject(p)
	if err != nil {
		return SendErrMsg(err)
	}

	return tea.Batch(
		SendSuccessMsg(p.Name+" added successfully"),
		SendScreenMsg(screen.ScreenList),
		LoadProjects)
}

