package repo

import (
    "context"
    "github.com/obh/go-playground/domains"

)

type Auth interface {

    AddToken(*domains.TokenDetails, string) error 

    GetUser(context.Context, string) (*domains.User, error)
}
