package ws

type message struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}
