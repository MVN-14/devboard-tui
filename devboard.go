package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Project struct {
	Id   int
	Name string
	Path string
}

func (p Project) FilterValue() string { return p.Name }
func (p Project) Title() string       { return p.Name }
func (p Project) Description() string { return p.Path }

type loadMsg struct {
	projects []Project
}

type errMsg error

type successMsg string

type Devboard struct {
	adding    bool
	confirming bool
	editting  bool
	errMsg    string
	list      list.Model
	msg       string
	nameInput textinput.Model
	pathInput textinput.Model
	project   Project
	quitting  bool
}

func loadProjects() tea.Msg {
	cmd := exec.Command("devboard", "list")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return errMsg(
			fmt.Errorf("loadProjects() error: %s\nstdout: %s\nstderr: %s", err, stdout.String(), stderr.String()),
		)
	}

	projects := []Project{}
	if err := json.Unmarshal(stdout.Bytes(), &projects); err != nil {
		return errMsg(
			fmt.Errorf("Error parsing project output: %s", err),
		)
	}

	return loadMsg{projects: projects}
}

func openProject() tea.Msg {
	// TODO Implement
	return nil
}

func updateProject(p Project) tea.Cmd {
	b, err := json.Marshal(p)
	if err != nil {
		return func() tea.Msg { return errMsg(fmt.Errorf("updateProjects() error: %s", err)) }
	}

	var stdout, stderr bytes.Buffer
	cmd := exec.Command("devboard", "update", string(b))
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err = cmd.Run(); err != nil {
		return func() tea.Msg {
			return errMsg(fmt.Errorf("updateProject() error: %s\nstdout: %s\nstderr: %s", err, stdout.String(), stderr.String()))
		}
	}

	return func() tea.Msg { return successMsg("Successfully updated project " + p.Name) }
}

func addProject(p Project) tea.Cmd {
	b, err := json.Marshal(p)
	if err != nil {
		return func() tea.Msg { return errMsg(fmt.Errorf("addProject() error: %s", err)) }
	}

	var stdout, stderr bytes.Buffer
	cmd := exec.Command("devboard", "add", string(b))
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err = cmd.Run(); err != nil {
		return func() tea.Msg {
			return errMsg(fmt.Errorf("updateProject() error: %s\nstdout: %s\nstderr: %s", err, stdout.String(), stderr.String()))
		}
	}

	return func() tea.Msg { return successMsg("Successfully added project " + p.Name) }
}

func deleteProject(id int) tea.Cmd {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("devboard", "remove", strconv.Itoa(id))
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return func() tea.Msg {
			return errMsg(fmt.Errorf("deleteProject() error: %s\nstdout: %s\nstder: %s", err, stdout.String(), stderr.String()))
		}
	}

	return func() tea.Msg {
		return successMsg(fmt.Sprintf("Project with id %d successfully deleted.", id))
	}
}

func (d Devboard) Init() tea.Cmd {
	return loadProjects
}

func (d Devboard) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		d.list.SetWidth(msg.Width)
		d.list.SetHeight(msg.Height - 4)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			if d.editting {
				d.editting = false
				d.project = Project{}
				return d, nil
			}
			d.quitting = true
			return d, tea.Quit
		case "enter":
			if d.editting {
				break
			}
			return d, openProject
		case "n":
			if d.editting || d.adding {
				break
			}
			d.adding = true
			d.project = Project{}
			d.nameInput.Focus()
			return d, textinput.Blink
		case "d":
			if d.editting || d.adding {
				break
			}
			p := d.list.SelectedItem().(Project)
			return d, tea.Batch(deleteProject(p.Id), loadProjects)
		case "e":
			if d.editting || d.adding {
				break
			}
			d.editting = true

			d.project = d.list.SelectedItem().(Project)
			d.nameInput.SetValue(d.project.Name)
			d.pathInput.SetValue(d.project.Path)

			d.nameInput.Focus()
			return d, textinput.Blink
		case "tab":
			if d.editting || d.adding {
				if d.nameInput.Focused() {
					d.nameInput.Blur()
					d.pathInput.Focus()
				} else {
					d.pathInput.Blur()
					d.nameInput.Focus()
				}
			}
		case "ctrl+s":
			d.project.Name = d.nameInput.Value()
			d.project.Path = d.pathInput.Value()
			d.nameInput.SetValue("")
			d.pathInput.SetValue("")
			if d.editting {
				d.editting = false
				return d, tea.Batch(updateProject(d.project), loadProjects)
			}
			if d.adding {
				d.adding = false
				return d, tea.Batch(addProject(d.project), loadProjects)
			}
		}
	case errMsg:
		d.errMsg = msg.Error()
		return d, nil
	case successMsg:
		d.msg = string(msg)
	case loadMsg:
		items := []list.Item{}
		for _, p := range msg.projects {
			items = append(items, p)
		}
		d.list.SetItems(items)
	}

	if d.editting || d.adding {
		d.nameInput, cmd = d.nameInput.Update(msg)
		cmds = append(cmds, cmd)
		d.pathInput, cmd = d.pathInput.Update(msg)
		cmds = append(cmds, cmd)
	} else {
		d.list, cmd = d.list.Update(msg)
		cmds = append(cmds, cmd)
	}

	return d, tea.Batch(cmds...)
}

func (d Devboard) View() string {
	if d.quitting {
		return ""
	}
	view := "Devboard\n"
	if d.errMsg != "" {
		view += d.errMsg
		return view
	}

	if d.editting || d.adding {
		view += "Project Name:\n"
		view += d.nameInput.View()
		view += "\nProject Path:\n"
		view += d.pathInput.View()
		view += "\n\nPress Ctrl-S to save."
	} else {
		view += d.list.View()
		view += fmt.Sprintf("\n\n%s\n", d.msg)
	}

	return view
}

func NewDevboard() Devboard {
	name := textinput.New()
	name.Width = 50
	name.Placeholder = "Project Name"

	path := textinput.New()
	path.Width = 50
	path.Placeholder = "Project Path"

	list := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	list.Title = "Projects"

	return Devboard{
		list:      list,
		nameInput: name,
		pathInput: path,
	}
}

func main() {
	tea.LogToFile("debug.log", "debug")
	p := tea.NewProgram(NewDevboard())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
