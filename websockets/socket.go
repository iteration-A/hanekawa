package websockets

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
	"github.com/iteration-A/hanekawa/constants"
)

var ChatroomChanIn = make(chan interface{})
var ChatroomChanOut = make(chan interface{})

type Subscribe struct {
	Token string
	Room  string
}
type Unsubscribe struct{}

type SendMessage struct {
	Room    string
	Content string
}

type NewMessageMsg struct {
	From    string
	Content string
}

type UserJoinedMsg struct {
	Username string
}

type UserLeftMsg struct {
	Username string
}

func SubscribeToChatRoom(room, token string) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	url := url.URL{Scheme: "ws", Host: constants.SocketUrl, Path: "/cable"}

	log.Printf("connecting to %s", url.String())

	headers := http.Header{"Authorization": []string{token}}
	conn, _, err := websocket.DefaultDialer.Dial(url.String(), headers)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	err = conn.WriteMessage(websocket.TextMessage, []byte(subscribe(room)))

	if err != nil {
		log.Println(err)
		return
	}

	done := make(chan struct{})

	go func() {
		defer close(done)

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}

			var response GenericResp
			json.Unmarshal(message, &response)
			switch response.Message.Type {
			case "user_joined":
				var response UserJoinedResp
				json.Unmarshal(message, &response)
				ChatroomChanOut <- UserJoinedMsg{
					Username: response.Message.Data.User,
				}

			case "user_left":
				var response UserLeftResp
				json.Unmarshal(message, &response)
				ChatroomChanOut <- UserLeftMsg{
					Username: response.Message.Data.User,
				}

			case "new_message":
				var response NewMessageResp
				json.Unmarshal(message, &response)
				ChatroomChanOut <- NewMessageMsg{
					From:    response.Message.Data.Username,
					Content: response.Message.Data.Content,
				}
			}
		}
	}()

	for {
		select {
		case msg := <-ChatroomChanIn:
			switch msg.(type) {
			case Unsubscribe:
				return

			case SendMessage:
				messageToSend := sendMessage(msg.(SendMessage).Room, msg.(SendMessage).Content)
				conn.WriteMessage(websocket.TextMessage, []byte(messageToSend))
			}

		case <-done:
			return

		case <-interrupt:
			log.Println("interrupted")

			err := conn.WriteMessage(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))

			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
