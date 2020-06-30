package serviceimpl

import (
    "log"
    "time"
    "net/http"
    "context"
    "regexp"
    "strings"
    "github.com/obh/go-playground/domains"
    "github.com/obh/go-playground/repo"
)

type User struct {
    UserRepo    repo.User
}

const (
    REQUEST_SUCCESS      =       100
    BAD_REQUEST_CODE     =       201
    NOT_FOUND_CODE       =       202
    BAD_REQUEST_EMAIL    =       "Invalid Email provided"
    NOT_FOUND_MSG        =       "Email not found"
)

func (u *User) CreateUser(ctx context.Context, req *domains.CreateUserRequest, httpReq *http.Request) (*domains.CrudResponse, error) {
   // validations should already be done 

   // find if we have any duplicates
   existingUser, err := u.UserRepo.GetUserByEmail(ctx, strings.TrimSpace(req.Email) )
   if existingUser != nil && existingUser.Id > 0 {
        log.Println("Duplicate email found")
        return &domains.CrudResponse{Status: "OK", Code: 401, Message: "Email already exists"}, nil
   }
   var hashedPwd = utils.HashPassword(req.Password) 
   var now = time.Now().Format("yyyy-MM-dd HH:mm:ss")
   createIntReq = &domains.CreateUserIntRequest{Email: req.Email, Phone: req.Phone, Password: hashedPwd, addedOn: now};
   user, err := u.UserRepo.CreateNewUser(ctx, req)
   if err != nil {
        log.Println("serviceimpl:user.go:: Failed when creating new user", err)
        return nil, err
   }
   log.Println("serviceimpl:user.go:: Got user ", user)
   return &domains.CrudResponse{Status: "OK", Code: 200, Message: "user created successfully"}, nil
}

func (u* User) GetUserByEmail(ctx context.Context, email string, httpReq *http.Request) (*domains.CrudResponse, error) {
    if email == "" {
        return &domains.CrudResponse{Status: "OK", Code: BAD_REQUEST_CODE, Message: BAD_REQUEST_EMAIL}, nil
    }
    var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$");
    if len(email) > 254 || !rxEmail.MatchString(email) {
        return &domains.CrudResponse{Status: "OK", Code: BAD_REQUEST_CODE, Message: BAD_REQUEST_EMAIL}, nil
    }
    
    user, err := u.UserRepo.GetUserByEmail(ctx, email)
    if err != nil {
        log.Println("serviceImpl:user.go:: Found error in finding user")
        return &domains.CrudResponse{Status: "OK", Code: NOT_FOUND_CODE, Message: NOT_FOUND_MSG}, nil
    }
    log.Println("serviceImpl:user.go:: Found user ", user)
    return &domains.CrudResponse{Status: "OK", Code: REQUEST_SUCCESS, Message: "Email found"}, nil
    //log.Println("in serviceimpl of Get User")
}
