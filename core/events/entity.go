package events

import (
	"encoding/json"
	"time"
)

func (i *Implementation) ProduceEntityInserted(e EventEntity) error {
	event := Event{
		Type:          "entity",
		Identifier:    e.EntityIdentifier(),
		SubIdentifier: e.PrimaryKeyValue(),
		Action:        "create",
		NewData:       e.String(),
		Timestamp:     time.Now(),
	}

	res, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return i.Produce(res, e.EntityIdentifier(), "insert")
}

func (i *Implementation) ProduceEntityUpdated(new EventEntity, old EventEntity) error {
	event := Event{
		Type:          "entity",
		Identifier:    new.EntityIdentifier(),
		SubIdentifier: new.PrimaryKeyValue(),
		Action:        "update",
		NewData:       new.String(),
		OldData:       old.String(),
		Timestamp:     time.Now(),
	}

	res, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return i.Produce(res, new.EntityIdentifier(), "update")
}
