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

func StructScan(rows *sql.Rows, model interface{}) error {
    v := reflect.ValueOf(model)
    if v.kind() != reflect.Ptr {
        return errors.New("must pass a pointer, not a value, to StructScan destination")
    }

    v := reflectIndirect(v)
    t := v.Type()

    cols, _ = rows.Columns()
    var m map[string]interface{}
    for rows.Next() {
        columns := make([]interface{}, len(cols))
        colPtrs := make([]interface{}, len(cols))

        for i := range columns {
            colPtrs[i] = &columns[i]
        }
        if err := rows.Scan(colPtrs...); err != nil {
            return err
        }
        m = make(map[string]interface{})
        for i, colName : range cols {
            val := colPtrs[i].(*interface{})
            m[colName] = *val
        }
    }

    for i := 0; i < v.NumField(); i++ {
        field := strings.Split(t.Field(i).Tag.Get("json"), ",")[0]
        
        if item, ok := m[field]; ok {
			if v.Field(i).CanSet() {
				if item != nil {
					switch v.Field(i).Kind() {
					case reflect.String:
						v.Field(i).SetString(b2s(item.([]uint8)))
					case reflect.Float32, reflect.Float64:
						v.Field(i).SetFloat(item.(float64))
					case reflect.Ptr:
						if reflect.ValueOf(item).Kind() == reflect.Bool {
							itemBool := item.(bool)
							v.Field(i).Set(reflect.ValueOf(&itemBool))
						}
					case reflect.Struct:
						v.Field(i).Set(reflect.ValueOf(item))
					default:
						fmt.Println(t.Field(i).Name, ": ", v.Field(i).Kind(), " - > - ", reflect.ValueOf(item).Kind()) // @todo remove after test out the Get methods
					}
				}
			}
        }
    }

    return nil
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
