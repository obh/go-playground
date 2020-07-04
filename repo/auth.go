package repo

import (
    "context"
    "github.com/obh/go-playground/domains"

)

type Auth interface {
    Authorize(context.Context, *domains.AuthorizeRequest) (*domains.AuthorizeIntResponse, error)

    AddToken(string, string, int64, int64) error 

    GetUser(context.Context, string) (*domains.User, error)
}
