package entity

type PushMessage struct {
	RoomID string
	Event  *Event
}
