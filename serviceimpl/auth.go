package serviceimpl

import(
    "context"
    "net/http"
    "fmt"
    "log"
    "github.com/obh/go-playground/domains"
    "github.com/obh/go-playground/repo"
    "github.com/obh/go-playground/utils"
)

type Auth struct {
    AuthRepo repo.Auth
}

func (a *Auth) Authorize(ctx context.Context, ar *domains.AuthorizeRequest, httpReq *http.Request) (*domains.AuthorizeResponse, error) {
    // verify email regex as well. Any sanitization required?
    if ar.Email == "" || ar.Password == "" {
        return &domains.AuthorizeResponse{Status: "SUCCESS", Code: 400, Message: "OK", }, nil
    }
    hashedPwd, err := utils.HashPassword(ar.Password)
    user, err := a.GetUserByEmail(ctx, ar.Email)
    if err != nil {
        log.Println("serviceimpl:user.go:: User not found with email ")
        return &domains.AuthorizeResponse{Status: "SUCCESS", Code: 400, Message: "OK", }, nil
    }
    if user.Password != hashedPwd {
        log.Println("serviceimpl:user.go:: User password does not match")
        return &domains.AuthorizeResponse{Status: "SUCCESS", Code: 400, Message: "OK",}, nil
    }
    token, err := utils.CreateToken(user.Email)
    tokens := map[string]string {
        "access_token" : token.AccessToken,
        "refresh_token" : token.RefreshToken,
    }
    log.Println(tokens)
    resp := &domains.AuthorizeResponse{Status: "SUCCESS", Code: 100, Message: "OK", }
    return resp, nil
}



func (a *Auth) Verify(c context.Context) (*domains.AuthorizeResponse, error) {
    fmt.Printf("calling verify service implementation")
    ar := &domains.AuthorizeResponse{Status: "SUCCESS", Code: 100, Message: "OK",}
    return ar, nil
}


func (a *Auth) AddToken(userId int64, td *domains.TokenDetails) error {
    at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
    rt := time.Unix(td.RtExpires, 0)
    now := time.Now()
    
    a.AuthRepo.AddToken(td.AccessUuid, td.RefreshUuid, td.AtExpires, td.RtExpires)
    log.Println("serviceimpl:auth.go:: Token added successfully")
}
