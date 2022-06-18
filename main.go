package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/iteration-A/emma/constants"
	login "github.com/iteration-A/emma/login"
)

type model struct {
	token       string
	loginScreen login.Model
	ticks       int
}

func initialModel() model {
	return model{
		token:       "",
		loginScreen: login.New(),
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.loginScreen.Init())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case constants.TokenMsg:
		m.token = msg.String()
	}
	m.ticks++
	loginScreenMod, cmd := m.loginScreen.Update(msg)
	m.loginScreen = loginScreenMod.(login.Model)
	return m, tea.Batch(cmd)
}

func (m model) View() string {
	if m.token == "" {
		return fmt.Sprintf("%v\n%d", m.loginScreen.View(), m.ticks)
	} else {
		return "ok"
	}
}

func main() {
	var Program = tea.NewProgram(initialModel())
	if err := Program.Start(); err != nil {
		fmt.Printf("Ups.\n%v\n", err)
		os.Exit(1)
	}
}
