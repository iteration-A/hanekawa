package websockets

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
