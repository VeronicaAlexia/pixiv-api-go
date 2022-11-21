package request

import (
	"io"
	"net/http"
	"net/url"
)

var PixivKey = struct {
	Token        string
	RefreshToken string
}{}

type Request struct {
	Path     string            // API Path
	Mode     string            // GET, POST, PUT
	Header   map[string]string // Request Header
	Query    map[string]string // Query Params
	Params   url.Values        // init in url.Values
	requests *http.Request     // init in http.NewRequest
}

type Response struct {
	Response *http.Response // Response from http.DefaultClient.Do
	Request  *Request       // Request from type Request
	Body     io.ReadCloser  // Body from Response
	content  []byte         // Body -> Content []byte
	text     string         // Content -> string
	retry    int            // Retry request if failed
}
