package repo

import (
    "context"
    "github.com/obh/go-playground/domains"
)

type User interface {
    GetUserByEmail(context.Context, string) (*domains.User, error)

    CreateNewUser(context.Context, *domains.CreateUserIntRequest) (*domains.User, error)

    // Verifies email and password for a user
    VerifyUser(context.Context, string, string) (*domains.User, error)
}
