package prompts

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/iteration-A/hanekawa/constants"
)

var (
	warning = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(constants.Pink)).
		Padding(1, 2).
		Foreground(lipgloss.Color(constants.Pink))
)
