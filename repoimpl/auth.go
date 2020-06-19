package repoimpl

import (
    "context"
    "github.com/obh/go-playground/domains"
)

// Repoimpl does the implemenation for an external service/db call

// This is our AuthRepo client 
type Auth {
    Client *client
    AuthSvcBase string
}

func (a* Auth) Authorize(ctx context.Context, p *domains.AuthorizeRequest) (*domains.AuthorizeResponse, error) {
    // return the Authroize response from here

}
