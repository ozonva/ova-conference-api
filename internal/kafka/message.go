package kafka

type ChangeType int

const (
	CreateConference ChangeType = iota
	MultiCreateConference
	UpdateConference
	DeleteConference
)

type ConferenceChangedMessage struct {
	Type  ChangeType
	Value interface{}
}
