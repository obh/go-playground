package domains


// This is request for authorize internal service
type AuthorizeRequest struct {
    Username string     `json:"username"`
    Password string     `json:"password"`

}

// This is response from internal service
type AuthorizeIntResponse struct {
    Status int
    Message string
}

// This is service request
type AuthorizeHttpRequest struct {
    Username string
    Password string
}

type CrudResponse struct {
    Status  string
    Code    int
    Message string
    // Should add data as well
}

type AuthorizeResponse struct {
    Status  string
    Code    int
    Message string
    // Should add data as well
}

type TokenDetails struct {
    AccessToken     string
    RefreshToken    string
    AccessUuid      string
    RefreshUuid     string
    AtExpires       int64
    RtExpires       int64
}
