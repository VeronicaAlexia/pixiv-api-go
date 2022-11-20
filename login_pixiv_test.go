package main

import (
	"fmt"
	"github.com/VeronicaAlexia/pixiv-api-go/pixiv"
	"testing"
)

func TestLoginPixiv(t *testing.T) {
	if login, err := pixiv.ChromeDriverLogin(); err != nil {
		t.Error(err)
	} else {
		t.Log("login:", login)
		fmt.Println("PixivRefreshToken: ", login.RefreshToken)
		fmt.Println("PixivToken: ", login.AccessToken)
	}
}
