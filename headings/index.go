package headings

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/iteration-A/hanekawa/constants"
)

var title = lipgloss.NewStyle().
	Background(lipgloss.Color(constants.Primary)).
	Foreground(lipgloss.Color(constants.Secondary))

func Title(message string) string {
	gap := strings.Repeat(" ", (constants.TermWidth-lipgloss.Width(message)) / 2)
	msg := gap + message + gap

	return lipgloss.Place(constants.TermWidth,
		lipgloss.Height(msg),
		lipgloss.Center, lipgloss.Center,
		title.Render(msg))
}
