package service

import (
	"context"
	"github.com/obh/go-playground/domains"
	"net/http"
)

// This is the service layer
type Auth interface {
	Authorize(context.Context, *domains.AuthorizeRequest, *http.Request) (*domains.AuthorizeResponse, error)

	Verify(context.Context, *domains.TokenRequest) (*domains.AuthorizeResponse, error)

	Logout(context.Context, *domains.TokenRequest) (*domains.LogoutResponse, error)
}
