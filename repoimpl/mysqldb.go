package repoimpl

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MySqlClient struct {
    *sql.DB
}

func Init() (*MySqlClient) {
    db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
        // do proper error handling
		panic(err.Error())  
	}
	defer db.Close()
}

func Get(query string, 
