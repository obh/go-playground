package repo

import (
    "context"
    "obh-crud/domains"
)

type Auth interface {
    Authorize(context.Context, *domains.AuthorizeRequest) *(domains.AuthorizeIntResp, error)
}
