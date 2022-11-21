<h1 align="center">
  <img src="./pixiv-logo.svg" alt="pixiv_logo" width ="400">
  <br><a href="https://www.pixiv.net/">pixiv</a> api for golang<br>  
</h1> 

# Introduction

- Partial code reference [everpcpc/pixiv](https://github.com/everpcpc/pixiv) implementation.
- Login part
  reference [DiheChen/pixiv-api](https://github.com/DiheChen/pixiv-api/blob/3e3c5a8690a29eec494e050a7f8a006c3353137b/auth/auth.go)
  implementation.

## install

```bash
go get github.com/VeronicaAlexia/pixiv-api-go
```

## import module

```go 
"github.com/VeronicaAlexia/pixiv-api-go/pixiv" 
```

## example

``` 
package main

import "github.com/VeronicaAlexia/pixiv-api-go/pixiv"  

func main() {
	if login, err := request.ChromeDriverLogin(); err != nil {
		panic(err)
	} else {
		init_pixiv := pixiv.InitPixivAppApi(login.AccessToken, login.RefreshToken)
		if Detail, err := init_pixiv.IllustDetail("87454525"); err == nil {
			t.Log(Detail)
		} else {
			t.Error(err)
		}
	}
}
```
