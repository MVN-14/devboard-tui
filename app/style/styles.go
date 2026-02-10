package style

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	fgColor         = lipgloss.Color("#D7D0C8")
	borderColor     = lipgloss.Color("#402A2C")
	accentPrimary   = lipgloss.Color("#D8315B")
	accentSecondary = lipgloss.Color("#006992")
	success         = lipgloss.Color("#00B16A")
	err             = lipgloss.Color("#DC0E0E")
)

var (
	HelpDescStyle = lipgloss.NewStyle().
			Foreground(accentSecondary)
	HelpKeyStyle = lipgloss.NewStyle().
			Foreground(accentPrimary)

	ListTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(accentPrimary)

	ListItemTitleSel = lipgloss.NewStyle().
				Foreground(accentPrimary).
				BorderForeground(accentPrimary).
				Border(lipgloss.ThickBorder(), false, false, false, true).
				Bold(true)
	ListItemDescSel = lipgloss.NewStyle().
			BorderForeground(accentPrimary).
			Border(lipgloss.ThickBorder(), false, false, false, true).
			Foreground(accentPrimary)
	ListItemTitle = lipgloss.NewStyle().
			Foreground(accentSecondary)
	ListItemDesc = lipgloss.NewStyle().
			Foreground(accentSecondary)

	ViewStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(borderColor)

	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(accentPrimary)

	toastStyle = lipgloss.NewStyle().
			Margin(1).
			Padding(1, 0).
			Bold(true).
			AlignHorizontal(lipgloss.Center).
			Foreground(lipgloss.Color("#1f1f1f"))
	ToastSuccess = toastStyle.
			Background(success)
	ToastError = toastStyle.
			Background(err)

	InputFocusedText = lipgloss.NewStyle().
				Bold(true)
	InputDefaultText = lipgloss.NewStyle().
				Bold(false)
	InputFocusedPrompt = lipgloss.NewStyle().
				Foreground(accentPrimary).
				Bold(true)
	InputDefaultPrompt = lipgloss.NewStyle().
				Foreground(fgColor).
				Bold(false)
)

func RenderView(v string) string {
	return ViewStyle.Render(v)
}

func RenderTitle(w int, t string) string {
	return lipgloss.PlaceHorizontal(w, lipgloss.Center, TitleStyle.Render(t))
}
