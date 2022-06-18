package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/iteration-A/hanekawa/chat"
	"github.com/iteration-A/hanekawa/constants"
	login "github.com/iteration-A/hanekawa/login"
	"github.com/iteration-A/hanekawa/rooms"
)

type model struct {
	token       string
	screens     []tea.Model
	screenIndex int
	firstRender bool
}

const (
	loginScreen = iota
	roomsScreen
	chatScreen
)

func initialModel() model {
	return model{
		token:       "",
		screens:     []tea.Model{login.New(), rooms.New(), chat.New()},
		screenIndex: loginScreen,
		firstRender: true,
	}
}

func (m model) Init() tea.Cmd {
	cmds := make([]tea.Cmd, len(m.screens))

	for index, screen := range m.screens {
		cmds[index] = screen.Init()
	}

	return tea.Batch(cmds...)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmds []tea.Cmd
	for index := range m.screens {
		if !m.firstRender {
			if index != m.screenIndex {
				continue
			}
		}

		var cmd tea.Cmd
		m.screens[index], cmd = m.screens[index].Update(msg)
		cmds = append(cmds, cmd)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	case constants.TokenMsg:
		m.token = msg.String()
		m.screenIndex = chatScreen
	case tea.WindowSizeMsg:
		m.firstRender = false
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return m.screens[m.screenIndex].View()
}

func main() {
	var Program = tea.NewProgram(initialModel())
	if err := Program.Start(); err != nil {
		fmt.Printf("Ups.\n%v\n", err)
		os.Exit(1)
	}
}
