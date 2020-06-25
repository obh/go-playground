package repo

import (
    "context"
    "github.com/obh/go-playground/domains"
)

type User interface {
    GetUserByEmail(context.Context, string) (*domains.User, error)
}
