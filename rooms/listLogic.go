package rooms

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type item struct {
	title string
	desc  string
}

func (i item) FilterValue() string { return i.title }

func (i item) String() string { return i.title }

func (i item) Height() int { return 1 }

func (i item) Spacing() int { return 0 }

func (i item) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }

func (i item) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	item, ok := listItem.(item)

	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, item)

	render := itemStyle.Render
	if index == m.Index() {
		render = selectedItemStyle.Render
	}

	fmt.Fprintf(w, render(str))
}
