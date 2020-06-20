package serviceimpl

import(
    "context"
    "net/http"
    "fmt"
    "github.com/obh/go-playground/domains"
    "github.com/obh/go-playground/repo"

)

type Auth struct {
    AuthRepo repo.Auth
}

func (a *Auth) Authorize(c context.Context, req *domains.AuthorizeRequest, httpReq *http.Request) (*domains.AuthorizeResponse, error) {
    fmt.Printf("Calling authorize service implementation\n")
    ar := &domains.AuthorizeResponse{Status: "SUCCESS", Code: 100, Message: "OK", }
    return ar, nil
}

func (a *Auth) Verify(c context.Context) (*domains.AuthorizeResponse, error) {
    fmt.Printf("calling verify service implementation")
    ar := &domains.AuthorizeResponse{Status: "SUCCESS", Code: 100, Message: "OK",}
    return ar, nil
}
