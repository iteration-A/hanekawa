package chat

import tea "github.com/charmbracelet/bubbletea"

type GoToRooms struct{}

func goToRoomsCmd() tea.Msg {
	return GoToRooms{}
}
