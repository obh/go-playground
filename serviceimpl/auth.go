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

// We assume that the token is in the header of the HTTP request
func (a *Auth) ValidateToken(ctx context.Context, req *http.Request) *domains.ValidateResponse {
    auth := r.Header.Get("Authorization") 
    if auth == "" {
        log.Println("serviceimpl:auth.go:: Missing Authorization header")
        return http
    }
    strArr := strings.Split(bearerToken, " ")
    if len(strArr) != 2 {
        log.Println("serviceimpl:auth.go:: Incorrect Authorization header")
    }
    t := strAttr[1]
    token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nilfmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(a.Secrets.AccessSecret), nil
    })
    if err != nil {
        log.Println("serviceimpl:auth.go:: Failed to verify token")
        return &domains.ValidateResponse{Status: 400, Message: "Failed to verify token"}
    }
    if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
        log.Println("serviceimpl:auth.go:: Failed to verify token")
        return &domains.ValidateResponse{Status: 400, Message: "Failed to verify token"
    }
}

func (a *Auth) Verify(c context.Context) (*domains.AuthorizeResponse, error) {
    fmt.Printf("calling verify service implementation")
    ar := &domains.AuthorizeResponse{Status: 100, Message: "OK",}
    return ar, nil
}

func (a *auth) VerifyAndExtractToken(req *http.Request) (*domains.TokenDetails, error) {
    
}

func (a *Auth) VerifyToken(req *http.Request) (*domains.AccessDetails, error) {
    auth := r.Header.Get("Authorization") 
    if auth == "" {
        log.Println("serviceimpl:auth.go:: Missing Authorization header")
        return http
    }
    strArr := strings.Split(bearerToken, " ")
    if len(strArr) != 2 {
        log.Println("serviceimpl:auth.go:: Incorrect Authorization header")
    }
    t := strAttr[1]
    token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nilfmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(a.Secrets.AccessSecret), nil
    })
    if err != nil {
        log.Println("serviceimpl:auth.go:: Failed to verify token")
        return &domains.ValidateResponse{Status: 400, Message: "Failed to verify token"}
    }
    if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
        log.Println("serviceimpl:auth.go:: Failed to verify token")
        return &domains.ValidateResponse{Status: 400, Message: "Failed to verify token"}
    }
}

func (a *Auth) ExtractToken(token interface{})  {
    claims := token.Claims.(jwt.MapClaims)
    accessUuid, ok := claims["access_uuid"].(string)
    if !ok {
        return "", err
    }
    userEmail, err := claims["email"].(string)
    &domains.AccessDetails{AccessUuid: accessUuid, Email: userEmail}, nil
   // Get access_uuid from cache
   it, err := a.GetToken(accessUuid)
   if err != nil {
        log.Println("serviceimpl:auth.go:: Token not in cache", err)
   }
} 
