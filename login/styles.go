package login

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/iteration-A/hanekawa/constants"
)

var (
	input         = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color(constants.Black)).Width(30)
	selectedInput = input.Copy().BorderForeground(lipgloss.Color(constants.Pink))
)
