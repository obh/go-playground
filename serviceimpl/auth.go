package serviceimpl

import(
    "context"
    "net/http"
    "fmt"
    "log"
    "github.com/obh/go-playground/domains"
    "github.com/obh/go-playground/repo"
    "github.com/obh/go-playground/utils"
    "github.com/obh/go-playground/config"
    "golang.org/x/crypto/bcrypt"
)

type Auth struct {
    AuthRepo repo.Auth
    Secrets  config.AuthConfig
}

func (a *Auth) Authorize(ctx context.Context, ar *domains.AuthorizeRequest, httpReq *http.Request) (*domains.AuthorizeResponse, error) {
    // verify email regex as well. Any sanitization required?
    if ar.Email == "" || ar.Password == "" {
        return &domains.AuthorizeResponse{Status: 400, Message: "Request not incorrect", }, nil
    }
    user, err := a.AuthRepo.GetUser(ctx, ar.Email)
    if err != nil {
        log.Println("serviceimpl:user.go:: User not found with email ")
        return &domains.AuthorizeResponse{Status: 400,  Message: "Incorrect login credentials", }, nil
    }
    ePwd := [] byte(user.Password)
    attempt := [] byte(ar.Password)
    err = bcrypt.CompareHashAndPassword(ePwd, attempt)
    if err != nil{
        log.Println("serviceimpl:user.go:: User password does not match")
        return &domains.AuthorizeResponse{Status: 400, Message: "Incorrect login credentials",}, nil
    }
    // Creating and saving token in Cache
    token := new(domains.TokenDetails)
    err = utils.CreateToken(a.Secrets.AccessSecret, a.Secrets.RefreshSecret, user.Email, token)
    err = a.AuthRepo.AddToken(token, user.Email)

    resp := &domains.AuthorizeResponse{Status: 100, Message: "OK", AccessToken: token.AccessToken, RefreshToken: token.RefreshToken }
    return resp, nil
}


func (a *Auth) Verify(c context.Context) (*domains.AuthorizeResponse, error) {
    fmt.Printf("calling verify service implementation")
    ar := &domains.AuthorizeResponse{Status: 100, Message: "OK",}
    return ar, nil
}


