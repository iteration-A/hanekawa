package login

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/iteration-A/hanekawa/constants"
)

var filename = fmt.Sprintf("%s/.token", constants.Basepath)

func saveToken(token []byte) bool {
	err := ioutil.WriteFile(filename, token, 0644)

	if err != nil {
		log.Fatal(err)
	}

	return true
}

type noToken struct{}
func retrieveToken() tea.Msg {
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		return noToken{}
	}

	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	token := string(bytes)

	if !validateToken(token) {
		return noToken{}
	}

	return tokenMsg(token)
}

func validateToken(token string) bool {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:3000/test", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", token)
	res, err := client.Do(req)
	if err != nil {
		return false
	}

	return res.Status == "200 OK"
}
