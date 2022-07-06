package login

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/iteration-A/hanekawa/constants"
)

var art string

const ASCII = "rice"

func Art() string {
	if art != "" {
		return art
	}

	content, err := ioutil.ReadFile(fmt.Sprintf("%s/ascii/%s", constants.Basepath, ASCII))

	if err != nil {
		log.Fatal("Could not read ascii art", err)
	}

	return string(content[0 : len(content)-1])
}
