package main

import (
	"fmt"
	"github.com/VeronicaAlexia/pixiv-api-go/pixiv"
	"testing"
)

// Path: get_image_info_test.go
// Compare this snippet from pkg\request\request.go:
// package request

func TestGetAuthorInfo(t *testing.T) {
	PixivRefreshToken := ""
	PixivToken := ""
	init_pixiv := pixiv.InitPixivAppApi(PixivToken, PixivRefreshToken)
	if illusts, err := init_pixiv.UserIllusts("27691", ""); err == nil {
		fmt.Println("illusts:", illusts)
		fmt.Println("NextURL: ", illusts.NextURL)
	} else {
		t.Error(err)
	}
}
