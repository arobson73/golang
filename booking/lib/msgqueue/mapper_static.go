package msgqueue

import (
	"encoding/json"
	"fmt"
	"log"

	"andy/booking/contracts"

	"github.com/mitchellh/mapstructure"
)

type StaticEventMapper struct{}

func (e *StaticEventMapper) MapEvent(eventName string, serialized interface{}) (Event, error) {
	var event Event

	log.Println("StaticEventMapper.MapEvent()")

	switch eventName {
	case "eventCreated":
		event = &contracts.EventCreatedEvent{}
	case "locationCreated":
		event = &contracts.LocationCreatedEvent{}
	case "eventBooked":
		event = &contracts.EventBookedEvent{}
	case "userCreated":
		event = &contracts.UserCreatedEvent{}
	case "adminUserCreated":
		event = &contracts.AdminUserCreatedEvent{}
	default:
		return nil, fmt.Errorf("unknown event type %s", eventName)
	}

	switch s := serialized.(type) {
	case []byte:
		err := json.Unmarshal(s, event)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal event %s: %s", eventName, err)
		}
	default:
		cfg := mapstructure.DecoderConfig{
			Result:  event,
			TagName: "json",
		}
		dec, err := mapstructure.NewDecoder(&cfg)
		if err != nil {
			return nil, fmt.Errorf("could not initialize decoder for event %s: %s", eventName, err)
		}

		err = dec.Decode(s)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal event %s: %s", eventName, err)
		}
	}

	return event, nil
}
