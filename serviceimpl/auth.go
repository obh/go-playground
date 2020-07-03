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


func (a *Auth) AddToken(userId int64, td *TokenDetails) error {
    at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
    rt := time.Unix(td.RtExpires, 0)
    now := time.Now()
    
    a.AuthRepo.AddToken(td.AccessUuid, td.RefreshUuid, td.AtExpires, td.RtExpires)
    log.Println("serviceimpl:auth.go:: Token added successfully")
}
