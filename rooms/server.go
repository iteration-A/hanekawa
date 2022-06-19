package rooms

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/iteration-A/hanekawa/constants"
)

type roomsResp struct {
	Rooms []item `json:"rooms"`
}

type item struct {
	Title string `json:"topic"`
}

type roomsMsg []item

func getRoomsCmd() tea.Msg {
	items := getRooms()
	return roomsMsg(items)
}

func getRooms() []item {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:3000/chat_rooms", nil)
	req.Header.Add("Authorization", constants.RetrieveTokenWithoutCheck())
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return []item{}
	}

	defer res.Body.Close()

	var rResp roomsResp
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return []item{}
	}

	err = json.Unmarshal(body, &rResp)
	if err != nil {
		log.Fatal(err)
		return []item{}
	}

	return rResp.Rooms
}
