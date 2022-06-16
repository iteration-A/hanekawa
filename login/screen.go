package login

import (
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var (
	termWidth, termHeight, _ = term.GetSize(int(os.Stderr.Fd()))
)

type model struct {
	selectedInput int
	inputs        []textinput.Model
	pink          bool
}

func InitialModel() model {
	username := textinput.New()
	username.Placeholder = "username"
	username.Width = 30
	username.CharLimit = 28
	username.Prompt = " "
	username.Focus()

	password := textinput.New()
	password.Placeholder = "password"
	password.Width = 30
	password.CharLimit = 28
	password.Prompt = " "
	password.EchoMode = textinput.EchoPassword

	return model{
		selectedInput: 0,
		inputs:        []textinput.Model{username, password},
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		{
			switch msg.String() {
			case "ctrl+c", "esc":
				return m, tea.Quit
			case "tab", "shift+tab", "enter", "up", "down":
				m.updateSelectedField(msg.String())
			default:
				m.pink = !m.pink
			}
		}
	}

	cmds := make([]tea.Cmd, len(m.inputs))
	for index, input := range m.inputs {
		m.inputs[index], cmds[index] = input.Update(msg)
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	doc := strings.Builder{}

	inputs := make([]string, len(m.inputs))
	for index, i := range m.inputs {
		if i.Focused() {
			inputs[index] = selectedInput.Render(i.View())
		} else {
			inputs[index] = input.Render(i.View())
		}
	}

	content := lipgloss.JoinVertical(lipgloss.Center, inputs...)
	var backgroundColor lipgloss.Color
	if m.pink {
		backgroundColor = lipgloss.Color(pink)
	} else {
		backgroundColor = lipgloss.Color(black)
	}

	ui := lipgloss.Place(termWidth-2, termHeight-2, lipgloss.Center, lipgloss.Center, content, lipgloss.WithWhitespaceChars("こんにちは"), lipgloss.WithWhitespaceForeground(backgroundColor))
	doc.WriteString(ui)

	return doc.String()
}
