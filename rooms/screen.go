package rooms

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/iteration-A/hanekawa/constants"
)

func roomSelectedCmd(room item) tea.Cmd {
	return func() tea.Msg {
		return constants.RoomSelectedMsg(room.Title)
	}
}

type Model struct {
	list         list.Model
	items        []list.Item
	choice       item
	gettingRooms bool
}

func initialModel() Model {
	width := constants.TermWidth - 4
	height := constants.TermHeight - 4

	l := list.New([]list.Item{}, item{}, width, height)
	l.Title = "Which room do you want to join?"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return Model{
		list:         l,
		gettingRooms: true,
	}
}

func New() Model {
	return initialModel()
}

func (m Model) Init() tea.Cmd {
	return getRoomsCmd
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case roomsMsg:
		m.gettingRooms = false
		items := make([]list.Item, len(msg))
		for index, item := range msg {
			items[index] = item
		}
		m.items = items
		cmd := m.list.SetItems(m.items)
		cmds = append(cmds, cmd)

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

	if m.choice.Title != "" {
		return m, roomSelectedCmd(m.choice)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, tea.Batch(append(cmds, cmd)...)
}

func (m Model) View() string {
	str := strings.Builder{}

	if m.gettingRooms {
		str.WriteString(fmt.Sprintf("%v", m.list.Items()))
	} else {
		str.WriteString(m.list.View())
	}

	return str.String()
}
