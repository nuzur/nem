package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/nuzur/nem/auth"
	"github.com/nuzur/nem/monitoring"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthUnaryServerInterceptor(auth auth.Interface, m *monitoring.Implementation) grpc.UnaryServerInterceptor {
	return grpc.UnaryServerInterceptor(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		if info.FullMethod == "/grpc.health.v1.Health/Check" || info.FullMethod == "/grpc.health.v1.Health/Watch" {
			return handler(ctx, req)
		}
		token, err := GetAuthorizationToken(ctx)
		if err != nil {
			m.Emit(monitoring.EmitRequest{
				ActionIdentifier: "invalid_auth_token",
				Message:          "error parsing auth token",
				Layer:            monitoring.ProtocolServiceLayer,
				LayerSubtype:     "protobuf",
				Type:             monitoring.EmitTypeError,
				Error:            err,
			})
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}

		err = auth.HandleToken(ctx, token)
		if err != nil {
			m.Emit(monitoring.EmitRequest{
				ActionIdentifier: "invalid_auth",
				Message:          "error handling auth",
				Layer:            monitoring.ProtocolServiceLayer,
				LayerSubtype:     "protobuf",
				Type:             monitoring.EmitTypeError,
				Error:            err,
			})
			return nil, status.Error(codes.Unauthenticated, fmt.Sprintf("invalid token %v", err))
		}
		return handler(ctx, req)
	})
}

func extractToken(header []string) (string, error) {
	if len(header) != 1 {
		return "", status.Error(codes.Unauthenticated, "invalid token format")
	}
	splitToken := strings.Split(header[0], "bearer ")
	if len(splitToken) < 2 {
		return "", status.Error(codes.Unauthenticated, "invalid token format")
	}
	return splitToken[1], nil
}

func GetAuthorizationToken(ctx context.Context) (string, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("missing metadata")
	}
	if len(meta["authorization"]) != 1 {
		return "", errors.New("authorization metadata not found")
	}

	return extractToken(meta["authorization"])
}
