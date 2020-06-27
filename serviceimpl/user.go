package serviceimpl

import (
    "log"
    "net/http"
    "context"
    "github.com/obh/go-playground/domains"
    "github.com/obh/go-playground/repo"
)

type User struct {
    UserRepo    repo.User
}

func (u* User) GetUserByEmail(ctx context.Context, email string, httpReq *http.Request) (*domains.CrudResponse, error) {
    log.Println("in serviceimpl of Get User")
    return nil, nil
}
