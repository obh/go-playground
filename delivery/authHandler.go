package delivery

import (
    "fmt"
    
    "github.com/obh/go-playground/domains"
)

func ConfigureAuthHandler(e *echo.Echo, svc service.Auth) {
    fmt.Printf("in ConigureAuthHandler")
    authHandler := &authHandler{authSvc: s}
    addAuthHandler(e, authHandler)
}

func addAuthHandler(e *echo.Echo, handler *authHandler){
    e.GET("/auth", handler.authorize)
}

type authHandler struct {
    authSvc service.Auth
}

func (h * authHandler) authorize(c echo.Context) {
   fmt.Printf("in authorize handler") 
}

