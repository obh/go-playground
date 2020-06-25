package repoimpl

import (
	"database/sql"
	"fmt"
    "log"
    "net/url"
	_ "github.com/go-sql-driver/mysql"

    "github.com/obh/go-playground/config"
)

type MySqlClient struct {
    Conn *sql.DB
}

const (
    getUserByEmail          string = `select * from Users where email = ?`
)

func InitDb(cfg config.DbConfig) (*MySqlClient, error) {
    fmt.Println(" Connecting to mysql", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
    connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
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
        log.Println("Error in mysql connection", err)
		return nil, err
	}
	dbConn.SetMaxOpenConns(cfg.MaxConnections)
    return &MySqlClient{dbConn}, nil
	//return dbConn, nil
}

/*
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
*/
