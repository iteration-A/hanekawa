package chat

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/iteration-A/hanekawa/constants"
	"github.com/iteration-A/hanekawa/headings"
	"github.com/iteration-A/hanekawa/statusbar"
)

// TODO: REMOVE
var tempMessages = []string{
	"[hanekawa]: hi",
	"[nadeko]: hi",
	"[nadeko]: bye",
	"[hanekawa]: :( 1",
	"[hanekawa]: :( 2",
	"[hanekawa]: :( 3",
	"[hanekawa]: :( 4",
	"[hanekawa]: :( 5",
	"[hanekawa]: :( 6",
	"[hanekawa]: :( 7",
	"[hanekawa]: :( 8",
	"[hanekawa]: :( 9",
	"[hanekawa]: :( 10",
	"[hanekawa]: :( 11",
	"[hanekawa]: :( 12",
	"[hanekawa]: :( 13",
	"[hanekawa]: :( 14",
	"[hanekawa]: :( 15",
	"[hanekawa]: :( 16",
	"[hanekawa]: :( 17",
	"[hanekawa]: :( 18",
	"[hanekawa]: :( 19",
	"[hanekawa]: :( 20",
	"[hanekawa]: :( 21",
	"[hanekawa]: :( 22",
	"[hanekawa]: :( 23",
	"[hanekawa]: :( 24",
	"[hanekawa]: :( 25",
	"[hanekawa]: :( 26",
	"[hanekawa]: :( 27",
	"[hanekawa]: :( 28",
	"[hanekawa]: :( 29",
	"[hanekawa]: :( 30",
	"[hanekawa]: :( 31",
	"[hanekawa]: :( 32",
	"[hanekawa]: :( 33",
	"[hanekawa]: :( 34",
	"[hanekawa]: :( 35",
	"[hanekawa]: :( 36",
	"[hanekawa]: :( 37",
	"[hanekawa]: :( 38",
	"[hanekawa]: :( 39",
	"[hanekawa]: :( 40",
	"[hanekawa]: :( 41",
}

type Model struct {
	content     string
	ready       bool
	viewport    viewport.Model
	input       textinput.Model
	typing      bool
	firstLetter bool
	chatName    string
	username    string
}

func initialModel() Model {
	i := textinput.New()
	i.CharLimit = 80
	i.Width = constants.TermWidth / 2
	i.Prompt = ""
	i.Placeholder = "Type something..."
	i.PlaceholderStyle = placeholder
	i.SetCursorMode(textinput.CursorStatic)

	return Model{
		input: i,
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
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		height := headerHeight + footerHeight

		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-height)
			m.viewport.YPosition = headerHeight
			m.content = joinMessages(tempMessages)
			m.viewport.SetContent(m.content)
			m.viewport.YOffset = m.calcExcess()
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - height
		}

	case constants.RoomSelectedMsg:
		m.chatName = string(msg)

	case constants.TokenMsg:
		m.username = string(msg.Username)

	case tea.KeyMsg:
		switch msg.String() {
		case "i":
			if !m.typing {
				m.typing = true
				m.input.Focus()
				m.firstLetter = true
			} else {
				m.firstLetter = false
			}

		case "g":
			if m.typing {
				break
			}
			m.viewport.YOffset = 0

		case "G":
			if m.typing {
				break
			}
			m.viewport.YOffset = 0
			m.viewport.YOffset = m.calcExcess()

		case "esc":
			m.typing = false
			m.input.Blur()

		case "enter":
			h := lipgloss.Height
			msgHeight := h(m.input.Value())
			m.addMessage(m.input.Value())
			m.input.SetValue("")
			m.viewport.SetContent(m.content)
			excess := m.calcExcess()
			m.viewport.YOffset = excess + msgHeight

		default:
			m.firstLetter = false
		}
	}

	var cmds []tea.Cmd
	var cmd tea.Cmd
	if m.typing {
		if !m.firstLetter {
			m.input, cmd = m.input.Update(msg)
			cmds = append(cmds, cmd)
		}
	} else {
		m.viewport, cmd = m.viewport.Update(msg)
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}

	return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView())
}

func (m Model) headerView() string {
	return headings.Title(m.chatName)
}
func (m Model) footerView() string {
	var msg string
	if m.typing {
		msg = "INSERT (esc)"
	} else {
		msg = "j↓ k↑ i(type)"
	}

	return statusbar.StatusLine(msg, m.input.View(), "Hanekawa🍙")
}

func (m Model) calcExcess() int {
	h := lipgloss.Height
	totalHeight := constants.TermHeight - h(m.headerView()) - h(m.footerView())
	excess := h(m.content) - totalHeight

	return excess
}

func (m *Model) addMessage(msg string) {
	formattedMsg := fmt.Sprintf("[%s] %s", m.username, msg)

	tempMessages = append(tempMessages, formattedMsg)
	m.content = joinMessages(tempMessages)
}

func joinMessages(messages []string) string {
	return strings.Join(messages, "\n")
}
