package domain

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

type Conference struct {
	Id               uuid.UUID
	UserId           uint64
	Name             string
	EventTime        *EventTime
	ParticipantCount int32
	SpeakerCount     int32
}

func (conference Conference) String() string {
	return fmt.Sprintf("Conference id:%v name:%v Date:%s", conference.Id, conference.Name, conference.EventTime.String())
}

func NewConference(userId uint64, name string, eventTime *EventTime) *Conference {
	result := Conference{
		EventTime: eventTime,
		Name:      name,
		UserId:    userId,
	}
	result.Id = uuid.New()

	return &result
}

func MakeConference(userId uint64, name string, eventTime *EventTime) Conference {
	result := Conference{
		EventTime: eventTime,
		Name:      name,
		UserId:    userId,
	}
	result.Id = uuid.New()

	return result
}

func NewConferenceFromByteJson(conferenceJson []byte) (*Conference, error) {
	var conference Conference
	err := json.Unmarshal(conferenceJson, &conference)
	return &conference, err
}
