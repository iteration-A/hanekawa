package rooms

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/iteration-A/hanekawa/constants"
)

func (i item) FilterValue() string { return i.Title }

func (i item) String() string { return i.Title }

func (i item) Height() int { return 1 }

func (i item) Spacing() int { return 0 }

func (i item) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }

func (i item) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	item, ok := listItem.(item)

	if !ok {
		return
	}

	str := fmt.Sprintf("%s\n", item)

	style := itemStyle
	if index == m.Index() {
		style = selectedItemStyle
	}

	content := style.Width(constants.TermWidth - 10).Render(str)
	fmt.Fprintf(w, content)
}
