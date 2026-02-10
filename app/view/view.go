package view

import (
	"github.com/MVN-14/devboard-tui/app/command"
	"github.com/MVN-14/devboard-tui/app/screen"
	"github.com/MVN-14/devboard-tui/app/style"
	"github.com/MVN-14/devboard-tui/devboard"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type input uint8

const (
	nameInput input = iota
	pathInput
	cmdInput
)

type keyMap struct {
	Next   key.Binding
	Prev   key.Binding
	Save   key.Binding
	Cancel key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Next, k.Prev, k.Save, k.Cancel}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Next, k.Prev},
		{k.Save, k.Cancel},
	}
}

type Model struct {
	focused input
	inputs  []textinput.Model
	help    help.Model
	keys    keyMap
	project devboard.Project
}

func New() Model {
	inputs := []textinput.Model{
		textinput.New(),
		textinput.New(),
		textinput.New(),
	}
	inputs[nameInput].Prompt = "Name    > "
	inputs[nameInput].Placeholder = "Project Name"

	inputs[pathInput].Prompt = "Path    > "
	inputs[pathInput].Placeholder = "Project Path"

	inputs[cmdInput].Prompt = "Command > "
	inputs[cmdInput].Placeholder = "Open Comand"

	for i := range len(inputs) {
		inputs[i].Width = 50
	}

	help := help.New()
	help.ShowAll = true

	return Model{
		inputs:  inputs,
		focused: nameInput,
		help:    help,
		keys: keyMap{
			Next: key.NewBinding(
				key.WithKeys("tab", "down"),
				key.WithHelp("tab/↓", "Next")),
			Prev: key.NewBinding(
				key.WithKeys("shift+tab", "up"),
				key.WithHelp("S-tab/↑", "Prev")),
			Save: key.NewBinding(
				key.WithKeys("ctrl+s"),
				key.WithHelp("C-s", "Save")),
			Cancel: key.NewBinding(
				key.WithKeys("ctrl+x", "esc"),
				key.WithHelp("C-x/Esc", "Cancel")),
		},
	}
}

func (m *Model) startEdit(p devboard.Project) tea.Cmd {
	m.project = p
	m.inputs[nameInput].SetValue(m.project.Name)
	m.inputs[pathInput].SetValue(m.project.Path)
	m.inputs[cmdInput].SetValue(m.project.Command)
	m.focused = nameInput
	focus(&m.inputs[m.focused])
	return textinput.Blink
}

func (m *Model) startAdd() tea.Cmd {
	m.inputs[nameInput].SetValue("")
	m.inputs[pathInput].SetValue("")
	m.inputs[cmdInput].SetValue("")
	m.focused = nameInput
	focus(&m.inputs[m.focused])
	return textinput.Blink
}

func unfocus(i *textinput.Model) {
	i.Blur()
	i.TextStyle = style.InputDefaultText.Inherit(i.TextStyle)
	i.PromptStyle = style.InputDefaultPrompt
}

func focus(i *textinput.Model) {
	i.Focus()
	i.TextStyle = style.InputFocusedText.Inherit(i.TextStyle)
	i.PromptStyle = style.InputFocusedPrompt
}

func (m *Model) focusNext() {
	unfocus(&m.inputs[m.focused])
	if m.focused == cmdInput {
		m.focused = nameInput
	} else {
		m.focused++
	}

	focus(&m.inputs[m.focused])
}

func (m *Model) focusPrev() {
	unfocus(&m.inputs[m.focused])
	if m.focused == nameInput {
		m.focused = cmdInput
	} else {
		m.focused--
	}
	focus(&m.inputs[m.focused])
}

func (m *Model) bindInputs() {
	m.project.Name = m.inputs[nameInput].Value()
	m.project.Path = m.inputs[pathInput].Value()
	m.project.Command = m.inputs[cmdInput].Value()
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case command.EditMsg:
		cmds = append(cmds, m.startEdit(msg.Project))
	case command.AddMsg:
		cmds = append(cmds, m.startAdd())
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Next):
			m.focusNext()
		case key.Matches(msg, m.keys.Prev):
			m.focusPrev()
		case key.Matches(msg, m.keys.Cancel):
			return m, command.SendScreenMsg(screen.ScreenList)
		case key.Matches(msg, m.keys.Save):
			m.bindInputs()
			if m.project.Id != 0 {
				return m, command.UpdateProject(m.project)
			} else {
				return m, command.AddProject(m.project)
			}
		}
	}

	m.inputs[m.focused], cmd = m.inputs[m.focused].Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	view := style.TitleStyle.Render(m.project.Name+"\n") + "\n"

	view += m.inputs[nameInput].View() + "\n"
	view += m.inputs[pathInput].View() + "\n"
	view += m.inputs[cmdInput].View() + "\n"

	view += "\n" + m.help.View(m.keys)

	return view
}
