package websockets

type Message struct {
	Content string `json:"content"`
}

type GenericResp struct {
	Message struct {
		Type string `json:"type"`
	} `json:"message"`
}

type UserJoinedResp struct {
	Message struct {
		Type string `json:"type"`
		Data    struct {
			User string `json:"user"`
		} `json:"data"`
	} `json:"message"`
}

type UserLeftResp struct {
	Message struct {
		Type string `json:"type"`
		Data    struct {
			User string `json:"user"`
		} `json:"data"`
	} `json:"message"`
}

type NewMessageResp struct {
	Message struct {
		Type string `json:"type"`
		Data    struct {
			Content string `json:"content"`
			Username string `json:"username"`
		} `json:"data"`
	} `json:"message"`
}
