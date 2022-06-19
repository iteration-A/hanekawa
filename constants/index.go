package constants

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
)

var _, base, _, ok = runtime.Caller(0)
var Basepath = filepath.Join(filepath.Dir(base), "../")

func RetrieveTokenWithoutCheck() string {
	bytes, err := ioutil.ReadFile(fmt.Sprintf("%s/.token", Basepath))

	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}
