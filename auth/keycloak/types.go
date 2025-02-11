package keycloak

import (
	"github.com/Nerzal/gocloak/v13"
	"go.uber.org/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
)

type Interface interface {
	HandleHTTP(w http.ResponseWriter, r *http.Request) error
}

type Implementation struct {
	logger *zap.Logger
	config Config
	client *gocloak.GoCloak
}

type Params struct {
	fx.In
	Logger *zap.Logger
	Config config.Provider
}

type Config struct {
	Hostname     string `yaml:"hostname"`
	Realm        string `yaml:"realm"`
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}
