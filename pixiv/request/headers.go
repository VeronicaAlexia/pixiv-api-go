package request

import "fmt"

func (req *Request) AddHead(key string, value string) {
	req.Header[key] = value
}

func (req *Request) Headers() {
	// Set default headers for request
	req.requests.Header.Set("Authorization", "Bearer "+PixivKey.Token)
	if req.Mode == "POST" {
		req.requests.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req.requests.Header.Set("Content-Type", "application/json")
	}

	for k, v := range PixivKey.Header {
		req.requests.Header.Set(k, v)
	}
	for k, v := range req.Header {
		req.requests.Header.Set(k, v)
	}
	// req.ShowHeaders()// Show headers key and value

}

func (req *Request) ShowHeaders() {
	for k, v := range req.requests.Header {
		fmt.Println("[", k, "]:", v)
	}
}
