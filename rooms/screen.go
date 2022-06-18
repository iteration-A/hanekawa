package rooms

import (
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/iteration-A/hanekawa/constants"
)

func roomSelectedCmd(room item) tea.Cmd {
	return func() tea.Msg {
		return constants.RoomSelectedMsg(room.title)
	}
}

type Model struct {
	list   list.Model
	items  []list.Item
	choice item
}

func initialModel() Model {
	items := []list.Item{
		item{title: "Wired sounds"},
		item{title: "Wired people"},
		item{title: "Sek"},
		item{title: "General"},
	}

	width := constants.TermWidth - 4
	height := constants.TermHeight - 4

	l := list.New(items, item{}, width, height)
	l.Title = "Which room do you want to join?"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return Model{
		list:  l,
		items: items,
	}
}

func New() Model {
	return initialModel()
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch key := msg.String(); key {
		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = i
			}

		case "j", "down":
			if m.list.Index() == len(m.items)-1 {
				m.list.Select(-1)
			}

		case "k", "up":
			if m.list.Index() == 0 {
				m.list.Select(len(m.items))
			}
		}
	}

	if m.choice.title != "" {
		return m, roomSelectedCmd(m.choice)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	str := strings.Builder{}

	str.WriteString(m.list.View())

	return str.String()
}
