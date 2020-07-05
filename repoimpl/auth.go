package repoimpl

import (
    "context"
    "log"
    "github.com/bradfitz/gomemcache/memcache"

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
    Cache   *Cache
}

const (
    getUser         =       "select * from Users where email = ?";
)


func (a *Auth) AddToken(td *domains.TokenDetails, email string) error {
    log.Println("repoimpl:auth.go:: Adding Token to memcache")    
  
    v := []byte(email)
    it1 := &memcache.Item{Key: td.AccessUuid, Value: v, Expiration: int32(td.AtExpires)}
    err := a.ache.Client.Set(it1)
    if err != nil {
        log.Println("repimpl:auth.go:: Error when inserting in cache", err)
        return err
    }

    it2 := &memcache.Item{Key: td.RefreshUuid, Value: v, Expiration: int32(td.RtExpires)}
    err = a.Cache.Client.Set(it2)
    if err != nil {
        log.Println("repimpl:auth.go:: Error when inserting in cache", err)
        return err
    }
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

func (a *Auth) GetToken(t string) (interface{}, error) {
    log.Println("repimpl:auth.go:: Getting token from cache")
    it, err := a.Cache.Client.Get(t)
    return it, err
}
