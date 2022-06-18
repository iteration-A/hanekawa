package login

func (m *Model) updateSelectedField(msg string) {
	switch msg {
	case "tab", "down":
		if m.selectedInput < len(m.inputs)-1 {
			m.selectedInput++
		} else {
			m.selectedInput = 0
		}
	case "shift+tab", "up":
		if m.selectedInput == 0 {
			m.selectedInput = len(m.inputs) - 1
		} else {
			m.selectedInput--
		}
	}

	for index := range m.inputs {
		if index == m.selectedInput {
			m.inputs[index].Focus()
		} else {
			m.inputs[index].Blur()
		}
	}
}
