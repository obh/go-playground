package repoimpl

import (
    "context"
    "log"
    "github.com/obh/go-playground/domains"   
    "github.com/obh/go-playground/utils"
)

type User struct {
    Conn    MySqlClient
}

const (
    getUserByEmail    =   "select * from Users where email = ?";
)

func (u* User) GetUserByEmail(ctx context.Context, email string) (*domains.User, error) {
    log.Println("Getting user by email: ", email)
    userRows, err := u.Conn.DB.QueryContext(ctx, getUserByEmail, email)
    if err != nil {
        log.Println("User Query failed ", err)
        return nil, err
    }

    if userRows != nil {
        defer userRows.Close()
    }

    user := &domains.User{}

    if userRows.Next() {
        err = utils.StructScan(userRows, user)
        if err != nil {
            log.Println("got error while parsing row in StructScan")
            return nil, err
        }
    }
    return user, nil
}
