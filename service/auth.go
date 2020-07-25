package service

import (
	"context"
	"github.com/obh/go-playground/domains"
	"net/http"
)

// This is the service layer
type Auth interface {
	Authorize(context.Context, *domains.AuthorizeRequest, *http.Request) (*domains.AuthorizeResponse, error)

	ValidateToken(context.Context, *domains.TokenRequest) (*domains.ValidateResponse, error)

	Verify(context.Context, *http.Request) (*domains.AuthorizeResponse, error)

	Logout(context.Context, *domains.LogoutRequest) (*domains.LogoutResponse, error)
}
