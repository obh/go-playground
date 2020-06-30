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
    insertUser        =   "insert into Users (email, phone, password, addedOn) values (?, ?, ?, ?)"
)

func (u *User) CreateNewUser(ctx context.Context, req *domains.CreateUserIntRequest) (*domains.User, error) {
    log.Println("repoimpl:user.go:: Creating new user")
    insUser, err := u.Conn.DB.Prepare(insertUser)
    if err != nil {
        log.Println("Failed while preparing insert query", err)
    }
    log.Println("running query ", insertUser, req.Email, req.Phone, req.Password, req.AddedOn)
    res, err := insUser.Exec(req.Email, req.Phone, req.Password, req.AddedOn)

    if err != nil {
        log.Println("Insert failed ", err)
    }
    insertId, _ := res.LastInsertId()
    log.Println("repoimpl:user.go:: Got id: ", insertId)
    return &domains.User{}, nil
    
}

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
    err = utils.StructScan(userRows, user)
    if err != nil {
        log.Println("got error while parsing row in StructScan")
        return nil, err
    }
    log.Println("Found User: ", user)
    return user, nil
}
