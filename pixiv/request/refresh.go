package request

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/VeronicaAlexia/pixiv-api-go/pixiv/pixivstruct"
	"io"
	"time"
)

func gen_clientHash(clientTime string) string {
	h := md5.New()
	_, _ = io.WriteString(h, clientTime)
	_, _ = io.WriteString(h, "28c1fdd170a5204386cb1313c7077b34f83e4aaf4aa829ce78c231e05b0bae2c")
	return hex.EncodeToString(h.Sum(nil))
}

func RefreshAuth() bool {
	params := map[string]string{
		"get_secure_url": "1",
		"client_id":      "MOBrBDS8blbauoSck0ZfDbtuzpyT",
		"client_secret":  "lsACyCD94FhDUtGTXi3QzcFE2uU1hqtDaKeqrdwj",
		"grant_type":     "refresh_token",
		"refresh_token":  PixivKey.RefreshToken,
	}
	client_time := time.Now().Format(time.RFC3339)
	response := Post(
		"https://oauth.secure.pixiv.net/auth/token",
		params, map[string]string{
			"X-Client-Time": client_time,
			"X-Client-Hash": gen_clientHash(client_time),
		}).Json(&pixivstruct.AccessToken{}).(*pixivstruct.AccessToken)

	if response.AccessToken == "" {
		fmt.Println("refresh auth error  ", response.AccessToken)
		return false
	} else {
		PixivKey.Token = response.AccessToken // update token in memory
		fmt.Println("refresh auth success,new token: ", response.AccessToken)
	}
	return true

}
