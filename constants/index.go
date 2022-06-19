package constants

import (
	"os"

	"golang.org/x/term"
)

const (
	Primary     = "#e9f542"
	PrimaryDark = "#757a2d"
	Secondary   = "#000000"
	Dark        = "#4a4848"
)

var (
	TermWidth, TermHeight, _ = term.GetSize(int(os.Stdout.Fd()))
)

type TokenMsg struct {
	Token    string
	Username string
}

type RoomSelectedMsg string

func (r RoomSelectedMsg) String() string { return string(r) }
