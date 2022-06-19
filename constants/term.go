package constants

import (
	"os"

	"golang.org/x/term"
)

var (
	TermWidth, TermHeight, _ = term.GetSize(int(os.Stdout.Fd()))
)
