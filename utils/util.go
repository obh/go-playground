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

func GenerateRandString() string {
    rand.Seed(time.Now().UnixNano())
    digits := "0123456789"
    specials := "~=+%^*/()[]{}/!@#$?|"
    all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
        "abcdefghijklmnopqrstuvwxyz" +
        digits + specials
    length := 8
    buf := make([]byte, length)
    buf[0] = digits[rand.Intn(len(digits))]
    buf[1] = specials[rand.Intn(len(specials))]
    for i := 2; i < length; i++ {
        buf[i] = all[rand.Intn(len(all))]
    }
    rand.Shuffle(len(buf), func(i, j int) {
        buf[i], buf[j] = buf[j], buf[i]
    })
    str := string(buf)
}

func HashPassword(password string) (int, error) {
    pwdBytes := []byte(password)
    hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    if err != nil {
        log.Println("Found error in generating password")
        return -1, err
    }
    return 0, nil
}
