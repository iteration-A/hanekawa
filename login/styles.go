package login

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/iteration-A/hanekawa/constants"
)

var (
	input         = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color(constants.Secondary)).Width(30).Padding(0)
	selectedInput = input.Copy().BorderForeground(lipgloss.Color(constants.Primary))
)
