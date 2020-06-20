package repo

import (
    "context"
    "github.com/obh/go-playground/domains"

)

type Auth interface {
    Authorize(context.Context, *domains.AuthorizeRequest) (*domains.AuthorizeIntResponse, error)
}
