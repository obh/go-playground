package domains

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

