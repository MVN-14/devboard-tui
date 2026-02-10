package list

import (
	"github.com/MVN-14/devboard-tui/app/command"
	"github.com/MVN-14/devboard-tui/app/screen"
	"github.com/MVN-14/devboard-tui/app/style"
	"github.com/MVN-14/devboard-tui/devboard"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type keyMap struct {
	Add    key.Binding
	Delete key.Binding
	Down   key.Binding
	Edit   key.Binding
	Help   key.Binding
	Open   key.Binding
	Quit   key.Binding
	Up     key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Open, k.Edit, k.Add, k.Delete, k.Help}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},
		{k.Add, k.Delete},
		{k.Open, k.Edit},
		{k.Quit, k.Help},
	}
}

type Model struct {
	list list.Model
	keys keyMap
	help help.Model
}

func New() Model {
	d := list.NewDefaultDelegate()
	d.Styles.NormalTitle = style.ListItemTitle
	d.Styles.NormalDesc = style.ListItemDesc
	d.Styles.SelectedTitle = style.ListItemTitleSel
	d.Styles.SelectedDesc = style.ListItemDescSel

	l := list.New([]list.Item{}, d, 0, 0)
	l.Title = "Projects"
	l.Styles.Title = style.ListTitleStyle

	l.DisableQuitKeybindings()
	l.SetShowHelp(false)

	keys := keyMap{
		Add:    key.NewBinding(key.WithKeys("n"), key.WithHelp("N", "New")),
		Delete: key.NewBinding(key.WithKeys("d"), key.WithHelp("D", "Delete")),
		Down:   key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("↓/j", "Down")),
		Edit:   key.NewBinding(key.WithKeys("e"), key.WithHelp("E", "Edit")),
		Help:   key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "help")),
		Open:   key.NewBinding(key.WithKeys("o"), key.WithHelp("O", "Open")),
		Quit:   key.NewBinding(key.WithKeys("ctrl+c"), key.WithHelp("Ctrl-C", "quit")),
		Up:     key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("↑/k", "Up")),
	}

	return Model{
		list: l,
		keys: keys,
		help: help.New(),
	}
}

func (m Model) GetSelected() devboard.Project {
	return m.list.SelectedItem().(ProjectListItem).Project
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Edit):
			return m, tea.Batch(
				command.SendEditMsg(m.GetSelected()),
				command.SendScreenMsg(screen.ScreenView))
		case key.Matches(msg, m.keys.Add):
			return m, tea.Batch(
				command.SendAddMsg(),
				command.SendScreenMsg(screen.ScreenView))
		case key.Matches(msg, m.keys.Delete):
			return m, command.DeleteProject(m.GetSelected().Id)	
		case key.Matches(msg, m.keys.Open):
			return m, command.OpenProject(m.GetSelected().Id)
		}
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.list.View() + "\n\n" + m.help.View(m.keys)
}

func (m *Model) SetSize(w, h int) {
	m.list.SetWidth(w)
	m.list.SetHeight(h - 2)
}

func (m *Model) SetItems(projects []devboard.Project) tea.Cmd {
	items := []list.Item{}
	for _, p := range projects {
		items = append(items, ProjectListItem{p})
	}
	return m.list.SetItems(items)
}

type ProjectListItem struct {
	Project devboard.Project
}

func (i ProjectListItem) FilterValue() string { return i.Project.Name }
func (i ProjectListItem) Title() string       { return i.Project.Name }
func (i ProjectListItem) Description() string { return i.Project.Path }
