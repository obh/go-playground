package repo

import (
    "context"
    "github.com/obh/go-playground/domains"
)

type User interface {
    GetUserByEmail(context.Context, string) (*domains.User, error)

    CreateNewUser(context.Context, *domains.CreateUserRequest) (*domains.User, error)
}
