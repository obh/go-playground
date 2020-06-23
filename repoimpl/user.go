package repoimpl

import (
    
)

type User struct {
    Client *Client
    UserSvcBase string
}

func (a* User) UserCreate(ctx context.Context, req *domains.UserCreateIntRequest) (*domains.UserCreateIntResponse, error) {
    fmt.Printf("in repoimpl UserCreate")
}
