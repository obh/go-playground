package domains

type User struct {
    Id          int64       `json:"id"`
    Email       string      `json:"email"`
    Phone       string      `json:"phone"`
    Password    string      `json:"password"` // represents the hashedpassword
}

type CreateUserRequest struct {
    Email   string      `json:"email" validate:"required,email,max=60"`
    Phone   string      `json:"phone" validate:"required,numeric,min=8,max=14"`
    Password string     `json:"password" validate:"required,min=6,max=40"`
}

type CreateUserIntRequest struct {
    Email       string
    Phone       string
    Password    string
    AddedOn     string
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

