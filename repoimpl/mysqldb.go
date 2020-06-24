package repoimpl

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MySqlClient struct {
    Conn *sql.DB
}

const (
    getUserByEmail          string = `select * from Users where email = ?`
)

func Init(cfg config.DBConfig) (*MySqlClient, error) {
    connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.Name)
	val := url.Values{}
	val.Add("charset", "utf8")
	val.Add("parseTime", "True")
	val.Add("loc", "Local")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		return nil, err
	}
	err = dbConn.Ping()
	if err != nil {
		return nil, err
	}
	dbConn.SetMaxOpenConns(cfg.MaxConnections)
	return dbConn, nil
}

func (sql *MySqlClient) GetUserByEmail(ctx context.Context, email string) ([]domains.UserDetails, error) {
    userRows, err := sql.Conn.QueryContext(ctx, getUserByEmail, email)
    if err != nil {
        fmt.Printf("Error while fetching user ", err)
        return nil
    }
    if userRows != nil {
        defer userRows.Close()
    }

    var userDetailRows []domains.UserDetails
    for userRows.Next() {
        user := domains.UserDetails{}
        err := StructScan(userRows, &user)

        if err != nil {
            fmt.Printf("cannot understand this struct")
            return nil
        }
        userDetailRows = append(userDetailRows, user) 
    }
    return userDetailRows, nil
}
