package login

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/iteration-A/hanekawa/constants"
)

var art string

func Art() string {
	if art != "" {
		return art
	}

	content, err := ioutil.ReadFile(fmt.Sprintf("%s/ascii/hanekawa", constants.Basepath))

	if err != nil {
		log.Fatal("Could not read ascii art", err)
	}

	return string(content[0 : len(content)-1])
}
