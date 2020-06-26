package repoimpl

import (
    "log"
    "github.com/obh/go-playground/domains"   
)

type UserSql struct {
    conn *MySqlClient
    UserSvcBase string
}

const (
    getUserByEmail    =   "select * from Users where email = ?";

func (u* UserSql) GetUserByEmail(ctx context.Context, email string) (*domains.User, error) {

    userRows, err := u.conn.QueryContext(ctx, getUserByEmail, email)
    if err != nil {
        log.Println("User Query failed ", err)
        return nil, err
    }

    if userRows != nil {
        defer userRows.close()
    }

    user := &domains.User{}
//    fmt.Printf("in repoimpl UserCreate")
    
//}
