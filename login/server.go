package login

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type tokenMsg string
type badCredentialsMsg struct{}

type tokenResp struct {
	Token string `json:"token"`
}

func getToken(username, password string) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(time.Second * 3)
		body, err := json.Marshal(map[string]string{
			"username": username,
			"password": password,
		})

		if err != nil {
			log.Fatalf("An error ocurred\n%v", err)
		}

		resp, err := http.Post("http://localhost:3000/login", "application/json", bytes.NewBuffer(body))
		if resp.Status != "200 OK" {
			return badCredentialsMsg{}
		}

		if err != nil {
			log.Fatalf("An error ocurred\n%v", err)
		}
		defer resp.Body.Close()

		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("An error ocurred\n%v", err)
		}

		var tResp tokenResp
		json.Unmarshal(body, &tResp)
		return tokenMsg(tResp.Token)
	}
}
