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
    "github.com/obh/go-playground/utils"
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
   hashedPwd, err := utils.HashPassword(req.Password) 
   if err != nil {
        log.Println("Cannot Hash password")
        return &domains.CrudResponse{Status: "OK", Code: 500, Message: "Internal servier error"}, nil
   }
   var now = time.Now().Format("2006-01-02 15:04:05")
   createIntReq := &domains.CreateUserIntRequest{Email: req.Email, Phone: req.Phone, Password: hashedPwd, AddedOn: now};
   user, err := u.UserRepo.CreateNewUser(ctx, createIntReq)
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

func (u *User) ValidateUserLogin(ctx context.Context, loginReq *domains.LoginRequest, httpReq *http.Request) (*domains.CrudResponse, error) {
    // is this even required, Clean out response and other stuff
    if loginReq.Email == "" || loginReq.Password == "" {
        return &domains.CrudResponse{Status: "OK", Code: BAD_REQUEST_CODE, Message: BAD_REQUEST_EMAIL}, nil 
    }
    hashedPwd, err := utils.HashPassword(loginReq.Password) 
    user, err := GetUserByEmail(ctx, loginReq.Email)
    if err != nil {
        log.Println("serviceimpl:user.go:: User not found with email ")
        return &domains.CrudResponse{Status: "OK", Code: NOT_FOUND_CODE, Message: NOT_FOUND_MSG}, nil
    }
    if user.Password != hashedPwd {
        log.Println("serviceimpl:user.go:: User password does not match")
        return &domains.CrudResponse{Status: "OK", Code: NOT_FOUND_CODE, Message: NOT_FOUND_MSG}, nil
    }
    // ok looks user is verified
    // now we will create the token
    token, err := utils.CreateToken(user.Email)
    if err != nil {
        // Failed when creating token
    }
    // save token
    tokens := map[string]string {
        "access_token" : token.AccessToken,
        "refresh_token" : token.RefreshToken
    }
    return c.JSON(http.StatusOK, tokens)
}
