package login

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
)

var art string

var _, base, _, ok = runtime.Caller(0)
var basepath = filepath.Join(filepath.Dir(base), "../")

func Art() string {
	if !ok {
		log.Fatal("Could not read ascii art")
	}
	if art != "" {
		return art
	}

	content, err := ioutil.ReadFile(fmt.Sprintf("%s/ascii/hanekawa", basepath))

	if err != nil {
		log.Fatal("Could not read ascii art", err)
	}

	return string(content[0:len(content)-1])
}
