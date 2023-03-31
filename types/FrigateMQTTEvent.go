package types

type FrigateMQTTEvent struct {
	Type   string       `json:"type"`
	Before FrigateEvent `json:"before"`
	After  FrigateEvent `json:"after"`
}
