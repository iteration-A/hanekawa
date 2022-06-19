package constants

import (
	"fmt"
	"io/ioutil"
	"log"
)

func RetrieveTokenWithoutCheck() string {
	bytes, err := ioutil.ReadFile(fmt.Sprintf("%s/.token", Basepath))

	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}
