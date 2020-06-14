package main

import (
    "fmt"
    //"log"
    "net/http"
    "strings"

    "obh-crud/config"
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
    http.HandleFunc("/", rootHandler)
    //http.ListenAndServe(port, nil)
    //log.Fatal(http.ListenAndServe(":8081", nil))

    config := config.LoadConfig() 
    print(config)
    print("hello world")
}
