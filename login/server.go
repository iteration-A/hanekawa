package login

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type token string

func getToken(username, password string) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(time.Second * 5)
		return token("as$!34%")
	}
}
