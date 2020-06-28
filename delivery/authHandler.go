package delivery

import (
    "fmt"
    "net/http"
    "github.com/labstack/echo/v4"
    
    "github.com/obh/go-playground/service"
    "github.com/obh/go-playground/domains"
)

func ConfigureAuthHandler(e *echo.Echo, svc service.Auth) {
    fmt.Printf("in ConigureAuthHandler")
    authHandler := &authHandler{authSvc: svc}
    addAuthHandler(e, authHandler)
}

func addAuthHandler(e *echo.Echo, handler *authHandler){
    e.POST("/auth", handler.authorize)
}

type authHandler struct {
    authSvc service.Auth
}

func (h *authHandler) authorize(c echo.Context) error {
   ar := new(domains.AuthorizeRequest)
   if err := c.Bind(ar); err != nil {
        return c.String(http.StatusBadRequest, "Bad Request")
   }
   fmt.Printf(ar.Username)
   fmt.Printf("in authorize handler ", ar) 

   resp, err := h.authSvc.Authorize(c.Request().Context(), ar, c.Request())
   if err != nil {
        fmt.Printf("response from authorize service\n")
   }
   return c.JSON(http.StatusOK, resp)
}

