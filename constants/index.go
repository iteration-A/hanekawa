package constants

import (
	"os"

	"golang.org/x/term"
)

const (
	Pink  = "#d2738a"
	Black = "#000000"
)

var (
	TermWidth, TermHeight, _ = term.GetSize(int(os.Stdout.Fd()))
)

type TokenMsg string

func (t TokenMsg) String() string {
	return string(t)
}
