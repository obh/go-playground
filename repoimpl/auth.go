package repoimpl

import (
    "fmt"
    "context"
    "github.com/obh/go-playground/domains"
)

// Repoimpl does the implemenation for an external service/db call

// This is our AuthRepo client 
type Auth struct {
    Client *Client
    AuthSvcBase string
}

func (a* Auth) Authorize(ctx context.Context, p *domains.AuthorizeRequest) (*domains.AuthorizeIntResponse, error) {
    // return the Authroize response from here
    fmt.Printf("Calling internal authorization service")
    authIntResp := &domains.AuthorizeIntResponse{Status: 100, Message: "OK",}
    return authIntResp, nil
}
