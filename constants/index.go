package constants

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"golang.org/x/term"
)

var _, base, _, ok = runtime.Caller(0)
var Basepath = filepath.Join(filepath.Dir(base), "../")

const (
	Primary     = "#e9f542"
	PrimaryDark = "#757a2d"
	Secondary   = "#000000"
	Dark        = "#4a4848"
)

var (
	TermWidth, TermHeight, _ = term.GetSize(int(os.Stdout.Fd()))
)

func RetrieveTokenWithoutCheck() string {
	bytes, err := ioutil.ReadFile(fmt.Sprintf("%s/.token", Basepath))

	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}
