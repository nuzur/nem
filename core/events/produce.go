package events

import (
	saramafx "github.com/mklfarha/sarama-fx"

	"fmt"
	"github.com/nuzur/nem/monitoring"
	"golang.org/x/sync/errgroup"
)

func (i *Implementation) Produce(message []byte, entityIdentifier string, action string) error {
	if i.provider.Get("events.enabled").String() == "false" {
		return nil
	}
	eg := errgroup.Group{}

	eg.Go(func() error {
		return i.produceToKafkaTopic("test", message, entityIdentifier, action)
	})

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}

func (i *Implementation) produceToKafkaTopic(topic string, message []byte, entityIdentifier string, action string) error {
	if i.client == nil {
		return nil
	}
	err := i.client.SendMessage(saramafx.SendMessageRequest{
		Topic:   topic,
		Message: message,
	})

	if err != nil {
		i.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: fmt.Sprintf("event_produce_%s", action),
			EntityIdentifier: entityIdentifier,
			Message:          "error producing message",
			Layer:            monitoring.EventServiceLayer,
			LayerSubtype:     "kafka",
			Type:             monitoring.EmitTypeError,
			ExtraData:        map[string]string{"message": string(message), "topic": topic},
			Error:            err,
		})
	}

	i.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: fmt.Sprintf("event_produce_%s", action),
		EntityIdentifier: entityIdentifier,
		Message:          "produced message successfully",
		Layer:            monitoring.EventServiceLayer,
		LayerSubtype:     "kafka",
		Type:             monitoring.EmitTypeSuccess,
		ExtraData:        map[string]string{"message": string(message), "topic": topic},
	})
	return nil
}
