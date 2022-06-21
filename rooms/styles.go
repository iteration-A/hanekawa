package rooms

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	"github.com/iteration-A/hanekawa/constants"
)

var (
	titleStyle = lipgloss.NewStyle().
			Align(lipgloss.Center).
			Width(constants.TermWidth - 10).
			Foreground(lipgloss.Color(constants.Primary))

	itemStyle = lipgloss.NewStyle().
			Background(lipgloss.Color(constants.Primary)).
			Foreground(lipgloss.Color(constants.Secondary)).
			Align(lipgloss.Center)

	selectedItemStyle = itemStyle.Copy().
				Background(lipgloss.Color(constants.Secondary)).
				Foreground(lipgloss.Color(constants.Primary))

	paginationStyle = list.DefaultStyles().
			PaginationStyle.
			PaddingLeft(4)

	helpStyle = list.DefaultStyles().
			HelpStyle.
			PaddingLeft(4).
			PaddingBottom(1)

	quitTextStyle = lipgloss.NewStyle().
			Margin(1, 0, 2, 4)
)
