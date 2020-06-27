package delivery

import (
    "log"
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/obh/go-playground/service"
    //"github.com/obh/go-playground/domains"
)

func ConfigureUserHandler(e *echo.Echo, userSvc service.User) {
    log.Println("userHandler: configuring user handler")
    userHandler := &userHandler{userSvc: userSvc}
    addUserHandler(e, userHandler)
}

func addUserHandler(e *echo.Echo, handler *userHandler){
    e.POST("/user/findByEmail", handler.findByEmail)
}

type userHandler struct {
    userSvc service.User
}

func (h *userHandler) findByEmail(c echo.Context) error {
    var ar string
    if err := c.Bind(ar); err != nil {
        return c.String(http.StatusBadRequest, "Bad Request")
    }
    log.Println(ar)
    log.Println("userHandler: in FindByEmail")

    resp, err := h.userSvc.GetUserByEmail(c.Request().Context(), ar, c.Request())
    if err != nil {
        log.Println("userHandler: Got error in searching for user by email", err)
    }
    return c.JSON(http.StatusOK, resp)
}
