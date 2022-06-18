package statusbar

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/iteration-A/hanekawa/constants"
)

var (
	left = lipgloss.NewStyle().
		Background(lipgloss.Color(constants.Primary)).
		Foreground(lipgloss.Color(constants.Secondary)).
		Padding(0, 1)

	center = lipgloss.NewStyle().
		Background(lipgloss.Color(constants.Secondary)).
		Foreground(lipgloss.Color(constants.Primary))

	right = left.Copy()
)

func StatusLine(l, c, r string) string {
	w := lipgloss.Width

	l = left.Render(l)
	r = right.Render(r)

	gapWidth := constants.TermWidth - w(l) - w(r)
	gap := center.Width(gapWidth).Align(lipgloss.Left).PaddingLeft(1).Render(c)

	content := lipgloss.JoinHorizontal(lipgloss.Center, l, gap, r)
	return lipgloss.Place(constants.TermWidth, 1, lipgloss.Center, lipgloss.Center, content)
}
