package repoimpl

import(
  //  "bytes"
    "net/http"
  //  "time"
  //  "encoding/json"
    "context"
    "fmt"
)

type Client struct {
    *http.Client
}

func Init() (*Client) {
    return &Client{
        &http.Client{Timeout: 100},
    }
}

func (c *Client) Get(ctx *context.Context, uri string) (*http.Response, error) {
   fmt.Printf("Calling GET ") 
   r := &http.Response{Status: "200 OK", StatusCode: 200, }
   return r, nil
}
