package service

import (
    "context"
    "net/http"
    "github.com/obh/go-playground/domains"
)

// This is the service layer 
type Auth interface {
    Authorize(context.Context, *domains.AuthorizeRequest, *http.Request) (*domains.AuthorizeResponse, error)

    ValidateToken(context.Context, *http.Request) (*domains.ValidateResponse)

    Verify(context.Context, *http.Request) (*domains.AuthorizeResponse, error)

}
