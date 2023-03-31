package types

type FrigateClipsEvent struct {
	FrigateMQTTEvent
	ClipUri string `json:"clip_uri"`
}
