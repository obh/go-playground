package repoimpl

import (
    "fmt"
    "context"
    "log"
    "github.com/obh/go-playground/domains"
    "github.com/obh/go-playground/utils"
)

// Repoimpl does the implemenation for an external service/db call

// This is our AuthRepo client  - it requires
// Mysql -> to get users data
// Memcache -> to get tokens
type Auth struct {
    Client *Client
    Conn    MySqlClient
    AuthSvcBase string
}

const (
    getUser         =       "select * from Users where email = ?";
)

func (a* Auth) Authorize(ctx context.Context, p *domains.AuthorizeRequest) (*domains.AuthorizeIntResponse, error) {
    // return the Authroize response from here
    fmt.Printf("Calling internal authorization service")
    authIntResp := &domains.AuthorizeIntResponse{Status: 100, Message: "OK",}
    return authIntResp, nil
}


func (a *Auth) AddToken(accessUuid string, refreshUuid string, atExpires int64, rtExpires int64) error {
    log.Println("repoimpl:auth.go:: Adding Token to memcache")    
    return nil
}

func (a *Auth) GetUser(ctx context.Context, email string) (*domains.User, error) {
    log.Println("Getting user by email: ", email)
    userRows, err := a.Conn.DB.QueryContext(ctx, getUser, email)
    if err != nil {
        log.Println("User Query failed ", err)
        return nil, err
    }

    if userRows != nil {
        defer userRows.Close()
    }

    user := &domains.User{}
    err = utils.StructScan(userRows, user)
    if err != nil {
        log.Println("got error while parsing row in StructScan")
        return nil, err
    }
    log.Println("Found User: ", user)
    return user, nil
}
