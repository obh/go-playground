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

func (a *Auth) Authorize(c context.Context, ar *domains.AuthorizeRequest, httpReq *http.Request) (*domains.AuthorizeResponse, error) {
    if ar.Email == "" || ar.Password == "" {
        return &domains.CrudResponse{Status: "OK", Code: BAD_REQUEST_CODE, Message:BAD_REQUEST_EMAIL}, nil
    }
    hashedPwd, err := utils.HashPassword(loginReq.Password)
    user, err := GetUserByEmail(ctx, loginReq.Email)
    if err != nil {
        log.Println("serviceimpl:user.go:: User not found with email ")
        return &domains.CrudResponse{Status: "OK", Code: NOT_FOUND_CODE, Message: NOT_FOUND_MSG}, nil
    }
    if user.Password != hashedPwd {
        log.Println("serviceimpl:user.go:: User password does not match")
        return &domains.CrudResponse{Status: "OK", Code: NOT_FOUND_CODE, Message: NOT_FOUND_MSG},nil
    }
    token, err := utils.CreateToken(user.Email)
    tokens := map[string]string {
        "access_token" : token.AccessToken,
        "refresh_token" : token.RefreshToken
    }
    log.Println(tokens)
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
