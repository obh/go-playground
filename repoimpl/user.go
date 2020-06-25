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

    u.conn.QueryContext(ctx, getUserByEmail, email)
//    fmt.Printf("in repoimpl UserCreate")
    
//}
