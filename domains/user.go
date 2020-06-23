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


