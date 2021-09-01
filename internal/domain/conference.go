package domain

import (
	"encoding/json"
	"fmt"
)

type Conference struct {
	Id               int64      `db:"id"`
	Name             string     `db:"name"`
	EventTime        *EventTime `db:"event_time"`
	ParticipantCount int32      `db:"participant_count"`
	SpeakerCount     int32      `db:"speaker_count"`
}

func (conference Conference) String() string {
	return fmt.Sprintf("Conference id:%v name:%v Date:%s", conference.Id, conference.Name, conference.EventTime.String())
}

func NewConference(name string, eventTime *EventTime) *Conference {
	result := Conference{
		EventTime: eventTime,
		Name:      name,
	}

	return &result
}

func MakeConference(name string, eventTime *EventTime, participantCount int32, speakerCount int32) Conference {
	result := Conference{
		EventTime:        eventTime,
		Name:             name,
		ParticipantCount: participantCount,
		SpeakerCount:     speakerCount,
	}
	return result
}

func NewConferenceFromByteJson(conferenceJson []byte) (*Conference, error) {
	var conference Conference
	err := json.Unmarshal(conferenceJson, &conference)
	return &conference, err
}
