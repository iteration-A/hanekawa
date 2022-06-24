package websockets

import "fmt"

func subscribe(room string) string {
	return fmt.Sprintf(`
	{
		"command": "subscribe",
		"identifier": "{\"channel\": \"ChatRoomChannel\", \"topic\": \"%s\"}"
	}
	`, room)
}

func sendMessage(room, message string) string {
	return fmt.Sprintf(`
	{
		"command": "message",
		"identifier": "{\"channel\": \"ChatRoomChannel\", \"topic\": \"%s\"}",
		"data": "{\"action\": \"message\", \"content\": \"%s\"}"
	}
	`, room, message)
}
