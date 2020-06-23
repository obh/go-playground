package service

import(

)

type User interface {
    CreateUser(context.Context, *domains.UserCreateRequest, *http.Request) (*domains.ServiceResponse, error)
//    AuthUser(context.Context, *domains.UserAuthRequest, *http.Request) (*domains.ServiceResponse, error)
    
}
