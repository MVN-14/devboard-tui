package main

import (
	"fmt"
	"os"

	"github.com/MVN-14/devboard-tui/app"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if os.Getenv("DEBUG") != "" {
		tea.LogToFile("debug.log", "debug")
	}
	p := tea.NewProgram(app.New())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
