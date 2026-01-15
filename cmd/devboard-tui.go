package main

import (
	"fmt"
	"os"

	"github.com/MVN-14/devboard-tui/internal/devboard"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	program := tea.NewProgram(devboard.NewApp())
	_, err := program.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
