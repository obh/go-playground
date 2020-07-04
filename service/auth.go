package service

import (
    "context"
    "net/http"
    "github.com/obh/go-playground/domains"
)

// This is the service layer 
type Auth interface {
    Authorize(context.Context, *domains.AuthorizeRequest, *http.Request) (*domains.AuthorizeResponse, error)

    Verify(context.Context) (*domains.AuthorizeResponse, error)

    AddToken(int64, *domains.TokenDetails) error 
}
