package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/obh/go-playground/config"
	"github.com/obh/go-playground/delivery"
	"github.com/obh/go-playground/repoimpl"
	"github.com/obh/go-playground/serviceimpl"

	//echo
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4/middleware"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if strings.TrimSpace(r.URL.Path[1:]) == "" {
		fmt.Fprintf(w, "request routed to rootHandler : enter your name after /")
		return
	}

	//fmt.Fprintf will stream the output to http response stream
	fmt.Fprintf(w, "request routed to rootHandler:\n Hello! %s, welcome to Cashfree", r.URL.Path[1:])
	//fmt.Printf will stream the output to standard output device in our case shell console
	fmt.Printf("request routed to rootHandler:\n Hello! %s, welcome to Cashfree\n", r.URL.Path[1:])
}

func main() {
	port := ":8081"
	fmt.Println("Starting webservice on port {}....", port)
	//http.HandleFunc("/", rootHandler)
	//http.ListenAndServe(port, nil)
	//log.Fatal(http.ListenAndServe(":8081", nil))

	config := config.LoadConfig()
	print(config)
	print("hello world")

	// Echo instance
	e := echo.New()

	// client for Repo
	client := repoimpl.Init()

	// db client
	mysqlClient, err := repoimpl.InitDb(config.DbConfig)
	if err != nil {
		fmt.Println("MySql connection failed ", err)
	}

	// cache client
	cache, err := repoimpl.InitCache(config.CacheConfig)
	if err != nil {
		log.Println("Cache connection failed ", err)
	}
	cache.Client.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})

	// Auth service goes here. Start with repo implementation here
	authRepo := &repoimpl.Auth{Client: client, AuthSvcBase: "localhost", Conn: mysqlClient, Cache: cache}

	// inject the rep to service
	authSvc := &serviceimpl.Auth{AuthRepo: authRepo, Secrets: config.AuthConfig}
	// configure service
	delivery.ConfigureAuthHandler(e, authSvc)

	// User service goes here
	userRepo := &repoimpl.User{Conn: mysqlClient}
	userSvc := &serviceimpl.User{UserRepo: userRepo}
	delivery.ConfigureUserHandler(e, userSvc)

	if userRepo != nil {
		fmt.Println("Got user Repo")
	}

	e.Logger.Fatal(e.Start(":1323"))

}
