package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	login "github.com/iteration-A/emma/login"
)

func main() {
	p := tea.NewProgram(login.InitialModel())

	if err := p.Start(); err != nil {
		fmt.Printf("Ups.\n%v\n", err)
		os.Exit(1)
	}
}
