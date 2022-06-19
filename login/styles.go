package login

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/iteration-A/hanekawa/constants"
)

var (
	input = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(constants.Secondary)).
		Width(30).
		Padding(0).
		Margin(0, 1)

	selectedInput = input.Copy().
			BorderForeground(lipgloss.Color(constants.Primary))
)
