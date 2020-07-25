package serviceimpl

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/obh/go-playground/config"
	"github.com/obh/go-playground/domains"
	"github.com/obh/go-playground/repo"
	"github.com/obh/go-playground/utils"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	AuthRepo repo.Auth
	Secrets  config.AuthConfig
}

func (a *Auth) Logout(ctx context.Context, lr *domains.TokenRequest) (*domains.LogoutResponse, error) {
	// requires verify token
	return &domains.LogoutResponse{Status: 200, Message: "Logout successful"}, nil
}

func (a *Auth) Authorize(ctx context.Context, ar *domains.AuthorizeRequest, httpReq *http.Request) (*domains.AuthorizeResponse, error) {
	// verify email regex as well. Any sanitization required?
	if ar.Email == "" || ar.Password == "" {
		return &domains.AuthorizeResponse{Status: 400, Message: "Request not incorrect"}, nil
	}
	user, err := a.AuthRepo.GetUser(ctx, ar.Email)
	if err != nil {
		log.Println("serviceimpl:user.go:: User not found with email ")
		return &domains.AuthorizeResponse{Status: 400, Message: "Incorrect login credentials"}, nil
	}
	ePwd := []byte(user.Password)
	attempt := []byte(ar.Password)
	err = bcrypt.CompareHashAndPassword(ePwd, attempt)
	if err != nil {
		log.Println("serviceimpl:user.go:: User password does not match")
		return &domains.AuthorizeResponse{Status: 400, Message: "Incorrect login credentials"}, nil
	}
	// Creating and saving token in Cache
	token := new(domains.TokenDetails)
	err = utils.CreateToken(a.Secrets.AccessSecret, a.Secrets.RefreshSecret, user.Email, token)
	err = a.AuthRepo.AddToken(token, user.Email)

	resp := &domains.AuthorizeResponse{Status: 100, Message: "OK", AccessToken: token.AccessToken, RefreshToken: token.RefreshToken}
	return resp, nil
}

// We assume that the token is in the header of the HTTP request
func (a *Auth) ValidateToken(ctx context.Context, req *domains.TokenRequest) *domains.ValidateResponse {
	auth := req.AccessToken
	if auth == "" {
		log.Println("serviceimpl:auth.go:: Missing Authorization header")
		return &domains.ValidateResponse{Status: 400, Message: "Failed to verify token"}
	}
	t := auth
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
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
	return &domains.ValidateResponse{Status: 400, Message: "Failed to verify token"}
}

func (a *Auth) Verify(c context.Context, req *http.Request) (*domains.AuthorizeResponse, error) {
	fmt.Printf("calling verify service implementation")
	ar := &domains.AuthorizeResponse{Status: 100, Message: "OK"}
	return ar, nil
}

//func (a *auth) VerifyAndExtractToken(req *http.Request) (*domains.TokenDetails, error) {
//
//}

func (a *Auth) VerifyToken(req *http.Request) (*domains.ValidateResponse, error) {
	bearerToken := req.Header.Get("Authorization")
	if bearerToken == "" {
		log.Println("serviceimpl:auth.go:: Missing Authorization header")
		return nil, nil
	}
	strArr := strings.Split(bearerToken, " ")
	if len(strArr) != 2 {
		log.Println("serviceimpl:auth.go:: Incorrect Authorization header")
	}
	t := strArr[1]
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(a.Secrets.AccessSecret), nil
	})
	if err != nil {
		log.Println("serviceimpl:auth.go:: Failed to verify token")
		return &domains.ValidateResponse{Status: 400, Message: "Failed to verify token"}, nil
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		log.Println("serviceimpl:auth.go:: Failed to verify token")
		return &domains.ValidateResponse{Status: 400, Message: "Failed to verify token"}, nil
	}
	return &domains.ValidateResponse{Status: 400, Message: "Failed to verify token"}, nil
}

func (a *Auth) ExtractToken(token jwt.Token) {
	claims := token.Claims.(jwt.MapClaims)
	accessUuid, ok := claims["access_uuid"].(string)
	if !ok {
		return
	}
	userEmail, ok := claims["email"].(string)
	if !ok {
		log.Println("Email not found in token")
	}
	ar := &domains.AccessDetails{AccessUuid: accessUuid, Email: userEmail}
	// Get access_uuid from cache
	it, e := a.AuthRepo.GetToken(accessUuid)
	if e != nil {
		log.Println("serviceimpl:auth.go:: Token not in cache", e)
	}
	log.Println("Got access details from token: ", ar)
	log.Println("Got access details from cache: ", it)
}
