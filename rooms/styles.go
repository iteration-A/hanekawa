package rooms

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/iteration-A/hanekawa/constants"
)

var (
	placeholder = lipgloss.NewStyle().
		Background(lipgloss.Color(constants.Black)).
		Foreground(lipgloss.Color(constants.Pink))
)
