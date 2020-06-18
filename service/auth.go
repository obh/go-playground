package service

import (
    "context"
    "net/http"
    "obh-crud/domains"

)

type Auth interface {
    Authorize(context.Context, *domains.AuthorizeHttpRequest, *http.Request) (*domains.CrudResponse, error)
    Verify(context.Context) (*domains.CrudResponse, error)
}
