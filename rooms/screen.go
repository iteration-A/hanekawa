package rooms

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/iteration-A/hanekawa/constants"
	"github.com/iteration-A/hanekawa/headings"
	"github.com/iteration-A/hanekawa/statusbar"
)

// TODO: REMOVE
const content = `
[hanekawa]: hi
[nadeko]: bye
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
[hanekawa]: :(
`

type Model struct {
	content     string
	ready       bool
	viewport    viewport.Model
	input       textinput.Model
	typing      bool
	firstLetter bool
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
			m.viewport.SetContent(content)
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - height
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "i":
			if !m.typing {
				m.typing = true
				m.input.Focus()
				m.firstLetter = true
			}

		case "esc":
			m.typing = false
			m.input.Blur()

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
	return headings.Title("Chat rooms")
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
