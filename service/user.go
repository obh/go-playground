package service

import(
    "net/http"
    "context"
    "github.com/obh/go-playground/domains"
)

type User interface {
    CreateUser(context.Context, *domains.UserCreateRequest, *http.Request) (*domains.CrudResponse, error)
//    AuthUser(context.Context, *domains.UserAuthRequest, *http.Request) (*domains.ServiceResponse, error)
    
}
