package statusbar

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/iteration-A/hanekawa/constants"
)

var (
	left = lipgloss.NewStyle().
		Background(lipgloss.Color(constants.Pink)).
		Foreground(lipgloss.Color(constants.Black)).
		Padding(0, 1)

	center = lipgloss.NewStyle().
		Background(lipgloss.Color(constants.Black)).
		Foreground(lipgloss.Color(constants.Pink))

	right = left.Copy()
)

func StatusLine(l, c, r string) string {
	w := lipgloss.Width

	l = left.Render(l)
	r = right.Render(r)

	gapWidth := constants.TermWidth - w(l) - w(r)
	gap := center.Width(gapWidth).Align(lipgloss.Center).Render(c)

	content := lipgloss.JoinHorizontal(lipgloss.Center, l, gap, r)
	return lipgloss.Place(constants.TermWidth, 1, lipgloss.Center, lipgloss.Center, content)
}
