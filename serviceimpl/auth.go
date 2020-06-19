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

func (a *Auth) Authorize(c context.Context, req *domains.AuthRequest, httpReq *http.Request) (resp *domains.Response, error) {
    fmt.Printf("Calling authorize service implementation ")
}

func (a *Auth) Verify(c context.Context) (resp *domains.Response, error) {
    fmt.Printf("calling verify service implementation")
}
