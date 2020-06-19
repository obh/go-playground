package repoimpl

import(
    "bytes"
    "net/http"
    "time"
    "encoding/json"
    "context"
    "fmt"
)

type Client struct {
    *http.Client
}

func Init() {
    return &Client{ 
        &http.Client{Timeout: 100}
    }
}

func (c *Client) Get(ctx *context.Context, uri string) (*httpResponse, error) {
   fmt.Printf("Calling GET ") 
}
