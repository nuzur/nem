package keycloak

import (
	"context"
	"github.com/Nerzal/gocloak/v13"
	"github.com/gofrs/uuid"
	base "nem/auth"
)

func New(params Params) (base.Interface, error) {
	var config Config
	if err := params.Config.Get("auth.keycloak").Populate(&config); err != nil {
		return nil, err
	}

	i := Implementation{
		logger: params.Logger,
		config: config,
		client: gocloak.NewClient(config.Hostname),
	}
	return &i, nil
}

func (i *Implementation) GetUserID(ctx context.Context, token string) (uuid.UUID, error) {
	userInfo, err := i.client.GetUserInfo(ctx, token, i.config.Realm)
	if err != nil {
		return uuid.Nil, err
	}
	if userInfo.Sub != nil {
		return uuid.FromStringOrNil(*userInfo.Sub), nil
	}
	return uuid.Nil, err
}

func (i *Implementation) GetUserByID(ctx context.Context, userID string) (*gocloak.User, error) {
	jwt, err := i.client.LoginClient(ctx, i.config.ClientID, i.config.ClientSecret, i.config.Realm)
	if err != nil {
		return nil, err
	}
	return i.client.GetUserByID(ctx, jwt.AccessToken, i.config.Realm, userID)
}
