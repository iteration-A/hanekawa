package login

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/iteration-A/hanekawa/constants"
)

type Model struct {
	selectedInput  int
	inputs         []textinput.Model
	pink           bool
	loading        bool
	token          tokenMsg
	loader         progress.Model
	badCredentials bool
}

func New() Model {
	return initialModel()
}

func initialLoaderModel() progress.Model {
	loader := progress.New(progress.WithGradient(pink, black))
	loader.ShowPercentage = false
	loader.SetPercent(0.01)
	return loader
}

func initialModel() Model {
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

	return Model{
		selectedInput: 0,
		inputs:        []textinput.Model{username, password},
		loader:        initialLoaderModel(),
	}
}

type tickMsg time.Time
type clearErrorMsg struct{}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func clearErrorCmd() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(time.Second * 2)
		return clearErrorMsg{}
	}
}

func sleepAndThenPassTokenCmd(token string) tea.Cmd {
	return tea.Cmd(func() tea.Msg {
		time.Sleep(time.Second * 1)
		return constants.TokenMsg(token)
	})
}

func (m Model) Init() tea.Cmd {
	return tea.Batch([]tea.Cmd{textinput.Blink, tickCmd()}...)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		{
			switch msg.String() {
			case "ctrl+c", "esc":
				return m, tea.Quit
			case "tab", "shift+tab", "up", "down":
				m.updateSelectedField(msg.String())
			case "enter":
				if m.selectedInput == len(m.inputs)-1 {
					m.loading = true
					return m, getToken(m.inputs[0].Value(), m.inputs[1].Value())
				}
			default:
				m.pink = !m.pink
			}
		}

	case badCredentialsMsg:
		m.badCredentials = true
		return m, clearErrorCmd()

	case clearErrorMsg:
		m.badCredentials = false
		m.loading = false
		m.loader = initialLoaderModel()
		return m, nil

	case tokenMsg:
		m.token = msg
		cmd := m.loader.SetPercent(1.0)
		tokenCmd := sleepAndThenPassTokenCmd(string(m.token))
		return m, tea.Batch([]tea.Cmd{tickCmd(), cmd, tokenCmd}...)

	case tickMsg:
		var cmd tea.Cmd
		if m.loading {
			cmd = m.loader.IncrPercent(0.15)
		}
		return m, tea.Batch(tickCmd(), cmd)

	case progress.FrameMsg:
		loaderMod, cmd := m.loader.Update(msg)
		m.loader = loaderMod.(progress.Model)
		return m, cmd
	}

	cmds := make([]tea.Cmd, len(m.inputs))
	for index, input := range m.inputs {
		m.inputs[index], cmds[index] = input.Update(msg)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
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
	if m.loading {
		content = m.loader.View()
	}

	if m.badCredentials {
		content = "Bad credentials!"
	}

	var backgroundColor lipgloss.Color
	if m.pink {
		backgroundColor = lipgloss.Color(pink)
	} else {
		backgroundColor = lipgloss.Color(black)
	}

	ui := lipgloss.Place(constants.TermWidth-1, constants.TermHeight-2, lipgloss.Center, lipgloss.Center, content, lipgloss.WithWhitespaceChars("░"), lipgloss.WithWhitespaceForeground(backgroundColor))
	doc.WriteString(ui)

	return doc.String()
}
