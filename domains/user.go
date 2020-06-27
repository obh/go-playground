package domains

type User struct {
    id          int64
    email       string
    phone       string
}

type UserRequest struct {
    Id      int64       `json:"id"`
    Email   string      `json:"email"`
    Phone   string      `json:"phone"`
}

type UserAuthRequest struct {
    username       string
    password       string
}


type UserAuthResponse struct {
    username    string
    status      int
    token       string
}

type UserCreateRequest struct {
    email       string
    phone       string
    password    string
}

type UserCreateIntRequest struct {
    email       string
    phone       string
    password    string
}

type UserCreateIntResponse struct {
    email       string
    phone       string
    password    string
    status      string
    message     string
}

