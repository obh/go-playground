package delivery

import (
    "fmt"
    "net/http"
    "github.com/labstack/echo/v4"
    
    "github.com/obh/go-playground/service"
    //"github.com/obh/go-playground/domains"
)

func ConfigureAuthHandler(e *echo.Echo, svc service.Auth) {
    fmt.Printf("in ConigureAuthHandler")
    authHandler := &authHandler{authSvc: svc}
    addAuthHandler(e, authHandler)
}

func addAuthHandler(e *echo.Echo, handler *authHandler){
    e.GET("/auth", handler.authorize)
}

type authHandler struct {
    authSvc service.Auth
}

func (h * authHandler) authorize(c echo.Context) error {
   fmt.Printf("in authorize handler") 
   return c.String(http.StatusOK, "Hello, World!")
}

