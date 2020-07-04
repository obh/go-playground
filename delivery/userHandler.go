package delivery

import (
    "log"
    "net/http"
    "gopkg.in/go-playground/validator.v9"
    "github.com/labstack/echo/v4"

    "github.com/obh/go-playground/service"
    "github.com/obh/go-playground/domains"
)

// use a single instance of Validate, it caches struct info
// Todo This should go to a single parent handler
var validate *validator.Validate

func ConfigureUserHandler(e *echo.Echo, userSvc service.User) {
    log.Println("userHandler: configuring user handler")
    userHandler := &userHandler{userSvc: userSvc}
    addUserHandler(e, userHandler)
}

func addUserHandler(e *echo.Echo, handler *userHandler){
    e.POST("/users", handler.createUser)
}

type userHandler struct {
    userSvc service.User
}


func (h *userHandler) createUser(c echo.Context) error {
    
    ar := new(domains.CreateUserRequest)
    if err := c.Bind(ar); err != nil {
        return c.String(http.StatusBadRequest, "Bad Request")
    }
    // Validate the request
    validate = validator.New()
    err := validate.Struct(ar)
    if err != nil {
        log.Println("delivery:userHandler.go:: Found error in Create user requirest ", err)
        return c.JSON(http.StatusInternalServerError, "")
    }

    log.Println("delivery:userHandler.go:: Got request to create user ", ar)

    resp, err := h.userSvc.CreateUser(c.Request().Context(), ar, c.Request())
    if err != nil {
        return c.JSON(http.StatusInternalServerError, "")
        log.Println("userHandler: Got error in searching for user by email", err)
    }
    return c.JSON(http.StatusOK, resp)
}
