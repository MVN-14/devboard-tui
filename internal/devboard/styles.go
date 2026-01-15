package devboard

import "github.com/charmbracelet/lipgloss"


const (
	errColor = lipgloss.Color("#FA5C5C")
	fgColor = lipgloss.Color("#A5158C")
	borderColor = lipgloss.Color("#FF2DF1")
	titleColor = lipgloss.Color("#FF2DF1")
)

var viewStyle = lipgloss.NewStyle().
	Foreground(fgColor).
	PaddingTop(1).
	PaddingBottom(1).
	PaddingLeft(5).
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(borderColor)

var titleStyle = lipgloss.NewStyle().
	Bold(true).
	PaddingBottom(1).
	Foreground(titleColor)

var textStyle = lipgloss.NewStyle()

var errorStyle = lipgloss.NewStyle().
	Foreground(errColor)
