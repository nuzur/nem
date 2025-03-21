package keycloak

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

func (i *Implementation) HandleToken(ctx context.Context, token string) error {
	rptResult, err := i.client.RetrospectToken(ctx, token, i.config.ClientID, i.config.ClientSecret, i.config.Realm)
	if err != nil {
		return err
	}

	if !*rptResult.Active {
		return errors.New("inactive token")
	}

	//permissions := rptResult.Permissions
	return nil
}

func (i *Implementation) HandleHTTP(w http.ResponseWriter, r *http.Request) error {
	ctx := context.Background()
	headerToken, err := i.TokenFromHeader(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return err
	}

	err = i.HandleToken(ctx, headerToken)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return err
	}
	return nil
}

func (i *Implementation) TokenFromHeader(r *http.Request) (string, error) {
	reqToken := r.Header.Get("Authorization")
	if reqToken == "" {
		reqToken = r.Header.Get("authorization")
	}

	if reqToken == "" {
		return "", errors.New("invalid token len")
	}

	splitToken := strings.Split(reqToken, "bearer ")
	if len(splitToken) < 2 {
		return "", errors.New("invalid token len")
	}

	return splitToken[1], nil
}
