package devboard

import (
	"bytes"
	"encoding/json"
	"errors"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

func DevboardAdd() {}

type errMsg error
type projectsMsg []Project

func loadProjects() tea.Msg {
	projects, err := devboardList()
	if err != nil {
		return errMsg(err)
	}
	return projectsMsg(projects)
}

func devboardList() ([]Project, error) {
	projects := []Project{}

	cmd := exec.Command("devboard", "list")

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		return projects, errors.New(stdout.String() + "\n" + err.Error())
	}

	err = json.Unmarshal(stdout.Bytes(), &projects)
	if err != nil {
		return projects, err
	}

	return projects, nil
}
