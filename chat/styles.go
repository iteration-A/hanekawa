package chat

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/iteration-A/hanekawa/constants"
)

var (
	placeholder = lipgloss.NewStyle().
			Background(lipgloss.Color(constants.Secondary)).
			Foreground(lipgloss.Color(constants.Primary))

	joinedMessage = lipgloss.NewStyle().
			Foreground(lipgloss.Color(constants.Primary))
)
