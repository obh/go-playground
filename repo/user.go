package repo

import (
    "github.com/obh/go-playground/domains"
)

type User interface {
    UserCreate(context.Context, *domains.UserCreateIntRequest) (*domains.UserCreateIntResponse, error)
}
