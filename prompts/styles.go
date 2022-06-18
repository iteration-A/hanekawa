package prompts

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/iteration-A/hanekawa/constants"
)

var (
	Warning = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(constants.Primary)).
		Padding(1, 2).
		Foreground(lipgloss.Color(constants.Primary))

	Title = lipgloss.NewStyle().
		Padding(0).
		Margin(1).
		Align(lipgloss.Center).
		Foreground(lipgloss.Color(constants.Primary))
)
