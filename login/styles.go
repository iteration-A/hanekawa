package login

import "github.com/charmbracelet/lipgloss"

const (
	pink  = "#d2738a"
	black = "#000000"
)

var (
	input         = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color(black)).Width(30)
	selectedInput = input.Copy().BorderForeground(lipgloss.Color(pink))
)
