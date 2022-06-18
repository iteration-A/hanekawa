package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/iteration-A/hanekawa/constants"
	login "github.com/iteration-A/hanekawa/login"
	"github.com/iteration-A/hanekawa/rooms"
)

type model struct {
	token       string
	loginScreen login.Model
	roomsScreen rooms.Model
}

func initialModel() model {
	return model{
		token:       "",
		loginScreen: login.New(),
		roomsScreen: rooms.New(),
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.loginScreen.Init())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		}
	case constants.TokenMsg:
		m.token = msg.String()
	}
	loginScreenMod, cmd := m.loginScreen.Update(msg)
	m.loginScreen = loginScreenMod.(login.Model)
	return m, tea.Batch(cmd)
}

func (m model) View() string {
	if m.token == "" {
		return m.loginScreen.View()
	} else {
		return m.roomsScreen.View()
	}
}

func main() {
	var Program = tea.NewProgram(initialModel())
	if err := Program.Start(); err != nil {
		fmt.Printf("Ups.\n%v\n", err)
		os.Exit(1)
	}
}
