package main

import (
	"github.com/VeronicaAlexia/pixiv-api-go/pixiv"
	"github.com/VeronicaAlexia/pixiv-api-go/pkg/request"
	"testing"
)

// Path: get_image_info_test.go
// Compare this snippet from pkg\request\request.go:
// package request

func TestGetImageInfo(t *testing.T) {
	init_pixiv := pixiv.InitPixivAppApi()
	request.PixivRefreshToken = ""
	request.PixivToken = ""
	if Detail, err := init_pixiv.IllustDetail("87454525"); err == nil {
		t.Log(Detail)
	} else {
		t.Error(err)
	}
}
