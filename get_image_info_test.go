package main

import (
	"github.com/VeronicaAlexia/pixiv-api-go/pixiv"
	"testing"
)

// Path: get_image_info_test.go

func TestGetImageInfo(t *testing.T) {
	// You need to login first to get the token and refresh token
	PixivRefreshToke, PixivToken := "", ""
	init_pixiv := pixiv.InitPixivAppApi(PixivToken, PixivRefreshToke)
	if Detail, err := init_pixiv.IllustDetail("87454525"); err == nil {
		t.Log(Detail)
	} else {
		t.Error(err)
	}
}
