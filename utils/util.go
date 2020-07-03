package utils

import (
    "database/sql"
    "errors"
    "fmt"
    "log"
    "reflect"
    "strconv"
    "strings"
    "time"
    "golang.org/x/crypto/bcrypt"
)

func StructScan(rows *sql.Rows, model interface{}) error {
    v := reflect.ValueOf(model)
    if v.Kind() != reflect.Ptr {
        return errors.New("Must pass a pointer, not a value, to StructScan destination")
    }

    v = reflect.Indirect(v)
    t := v.Type()

    cols, _ := rows.Columns()
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
        for i, colName := range cols {
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
                            v.Field(i).SetString(i2s(item))
                        case reflect.Float32, reflect.Float64:
                            v.Field(i).SetFloat(item.(float64))
                        case reflect.Ptr:
                            if reflect.ValueOf(item).Kind() == reflect.Bool {
                                itemBool := item.(bool)
                                v.Field(i).Set(reflect.ValueOf(&itemBool))
                            }
                        case reflect.Int:
                            v.Field(i).SetInt(item.(int64))
                        case reflect.Int64:
                            v.Field(i).SetInt(item.(int64))
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

func i2s(i interface{}) string {
	switch i.(type) {
	case int:
		return strconv.Itoa(i.(int))
	case int64:
		return strconv.Itoa(int(i.(int64)))
	case []uint8:
		ba := []byte{}
		for _, b := range i.([]uint8) {
			ba = append(ba, byte(b))
		}
		return string(ba)
	case time.Time:
		return i.(time.Time).Format("2006-01-02 15:04:05")
	default:
		// TODO define a general log level
		log.Println("############")
		log.Println(reflect.TypeOf(i))
	}
	return "Default"
}


func HashPassword(password string) (string, error) {
    pwdBytes := []byte(password)
    hashedPassword, err := bcrypt.GenerateFromPassword(pwdBytes, bcrypt.DefaultCost)
    if err != nil {
        log.Println("Found error in generating password")
        return "", err
    }
    return string(hashedPassword), err
}


func CreateToken(accessSecret string, refreshSecret string, email string) (string, error) {
    td := &TokenDetails{}
    td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
    td.AccessUuid = uuid.NewV4().String()
    td.RtExpires = time.Now().Add(time.Hour * 24).Unix()
    td.RefreshUuid = uuid.NewV4().String()

    atClaims := jwt.MapClaims{}
    atClaims["authorized"] = true
    atClaims["access_uuid"] = td.AccessUuid
    atClaims["email"] = email
    atClaims["expiry"] = td.AtExpires
    at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
    td.AccessToken, err := at.SignedString([]byte(accessSecret))
    if err != nil {
        return "", err
    }

    rtClaims := jwt.MapClaims{}
    rtClaims["refresh_uuid"] = td.RefresUuid
    rtClaims["email"] = email
    rtClaims["exp"] = td.RtExpires
    rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
    td.AccessToken, err := at.SignedString([]byte(refreshSecret))
    if err != nil {
        return nil, err
    }
    return td, nil
}
