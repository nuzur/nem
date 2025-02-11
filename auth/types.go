package auth

import (
	"context"
	"github.com/gofrs/uuid"
	"net/http"
)

type Interface interface {
	HandleHTTP(w http.ResponseWriter, r *http.Request) error
	HandleToken(ctx context.Context, token string) error
	GetUserID(ctx context.Context, token string) (uuid.UUID, error)
}
