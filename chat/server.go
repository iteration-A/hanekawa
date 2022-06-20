package chat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/iteration-A/hanekawa/constants"
)

type MessagesMsg []message

type message struct {
	Content string `json:"content"`
	User    struct {
		Username string
	} `json:"user"`
}
type getLastMessagesResp struct {
	Chat struct {
		Topic    string    `json:"topic"`
		Messages []message `json:"messages"`
	} `json:"chat_room"`
}

func getLastMessagesCmd(topic string) tea.Cmd {
	return func() tea.Msg {
		return getLastMessages(topic)
	}
}

func getLastMessages(topic string) MessagesMsg {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("http://localhost:3000/chat_rooms/%s", topic), nil)
	req.Header.Add("Authorization", constants.RetrieveTokenWithoutCheck())
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var mResp getLastMessagesResp
	json.Unmarshal(body, &mResp)

	return MessagesMsg(mResp.Chat.Messages)
}
