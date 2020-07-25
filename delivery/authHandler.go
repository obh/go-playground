package delivery

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"

	"github.com/obh/go-playground/domains"
	"github.com/obh/go-playground/service"
)

//var validate *validator.Validate

func ConfigureAuthHandler(e *echo.Echo, svc service.Auth) {
	fmt.Printf("in ConigureAuthHandler")
	authHandler := &authHandler{authSvc: svc}
	addAuthHandler(e, authHandler)
}

func addAuthHandler(e *echo.Echo, handler *authHandler) {
	e.POST("/auth", handler.authorize)
	e.POST("/validate", handler.validate)
	e.POST("/logout", handler.logout)
}

type authHandler struct {
	authSvc service.Auth
}

func (h *authHandler) validate(c echo.Context) error {
	tr := new(domains.TokenRequest)
	if err := c.Bind(tr); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	validate = validator.New()
	err := validate.Struct(tr)
	if err != nil {
		log.Println("deliver::authHandler.go:: Found error in request ", err)
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	r, err := h.authSvc.Verify(c.Request().Context(), tr)
	return c.JSON(http.StatusOK, r)
}

func (h *authHandler) logout(c echo.Context) error {
	ar := new(domains.TokenRequest)
	if err := c.Bind(ar); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	validate = validator.New()
	err := validate.Struct(ar)
	if err != nil {
		log.Println("delivery::authHandler.go:: Found error in request ", err)
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	return c.JSON(http.StatusOK, "ok")
}

func (h *authHandler) authorize(c echo.Context) error {
	ar := new(domains.AuthorizeRequest)
	if err := c.Bind(ar); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	validate = validator.New()
	err := validate.Struct(ar)
	if err != nil {
		log.Println("delivery::userHandler.go:: Found error in request ", err)
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	resp, err := h.authSvc.Authorize(c.Request().Context(), ar, c.Request())
	if err != nil {
		fmt.Printf("response from authorize service\n")
	}
	return c.JSON(http.StatusOK, resp)
}
