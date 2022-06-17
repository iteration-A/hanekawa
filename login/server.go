package login

import (
	tea "github.com/charmbracelet/bubbletea"
)

type token string

func getToken(username, password string) tea.Cmd {
	return func() tea.Msg {
		return token("as$!34%")
	}
}
