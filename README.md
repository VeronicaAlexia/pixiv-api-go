# This project is a pixiv android api for go.
 
# Introduction 
- Partial code reference [everpcpc/pixiv](https://github.com/everpcpc/pixiv) implementation.
- Login part reference [DiheChen/pixiv-api](https://github.com/DiheChen/pixiv-api/blob/3e3c5a8690a29eec494e050a7f8a006c3353137b/auth/auth.go) implementation.

## install
```bash
go get github.com/VeronicaAlexia/pixiv-api-go
```

## example
``` 
package main

import (
	"fmt"
	"github.com/VeronicaAlexia/pixiv-api-go/pixiv"
	"github.com/VeronicaAlexia/pixiv-api-go/pkg/request"
)

func main() {
	if login, err := request.ChromeDriverLogin(); err != nil {
		panic(err)
	} else {
		request.PixivRefreshToken = login.RefreshToken
		request.PixivToken = login.AccessToken
		println("PixivRefreshToken: ", login.RefreshToken)
		println("PixivToken: ", login.AccessToken)
		init_pixiv := pixiv.InitPixivAppApi()
		if Detail, err := init_pixiv.IllustDetail("87454525"); err == nil {
			fmt.Println(Detail)
		} else {
			panic(err)
		}
	}

}
```
