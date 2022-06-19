package login

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/iteration-A/hanekawa/constants"
)

type tokenMsg string
type badCredentialsMsg struct{}
type serverErrorMsg struct{}

type tokenResp struct {
	Token string `json:"token"`
	Username string `json:"username"`
}

func getToken(username, password string) tea.Cmd {
	return func() tea.Msg {
		body, err := json.Marshal(map[string]string{
			"username": username,
			"password": password,
		})

		if err != nil {
			log.Fatalf("An error ocurred\n%v", err)
		}

		resp, err := http.Post("http://localhost:3000/login", "application/json", bytes.NewBuffer(body))
		if err != nil {
			return serverErrorMsg{}
		}
		defer resp.Body.Close()

		switch resp.Status {
		case "200 OK":
			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalf("An error ocurred\n%v", err)
			}

			var tResp tokenResp
			json.Unmarshal(body, &tResp)

			saveToken([]byte(tResp.Token))

			return constants.TokenMsg{
				Token: tResp.Token,
				Username: tResp.Username,
			}
		case "401 Unauthorized":
			return badCredentialsMsg{}
		default:
			return serverErrorMsg{}
		}
	}
}
