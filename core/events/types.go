package events

import (
	"go.uber.org/fx"

	saramafx "github.com/mklfarha/sarama-fx"

	"go.uber.org/config"
	"nem/monitoring"
	"time"
)

type Implementation struct {
	client *saramafx.Client

	monitoring *monitoring.Implementation
	provider   config.Provider
}

type Params struct {
	fx.In

	Client *saramafx.Client

	Monitoring *monitoring.Implementation
	Provider   config.Provider
}

type EventEntity interface {
	EntityIdentifier() string
	PrimaryKeyValue() string
	String() string
}

type Event struct {
	Type          string    `json:"type"`
	Identifier    string    `json:"identifier"`
	SubIdentifier string    `json:"sub-identifier"`
	Action        string    `json:"action"`
	NewData       string    `json:"new-data"`
	OldData       string    `json:"old-data"`
	Timestamp     time.Time `json:"timestamp"`
}
