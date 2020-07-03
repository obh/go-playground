package service

import(
    "net/http"
    "context"
    "github.com/obh/go-playground/domains"
)

type User interface {
    CreateUser(context.Context, *domains.CreateUserRequest, *http.Request) (*domains.CrudResponse, error)

    GetUserByEmail(context.Context, string, *http.Request) (*domains.CrudResponse, error)

    ValidateUserLogin(context.Context, *domains.LoginRequest, *http.Request) (*domains.CrudResponse, error)

//    AuthUser(context.Context, *domains.UserAuthRequest, *http.Request) (*domains.ServiceResponse, error)
    

}
