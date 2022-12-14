package pixiv

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/VeronicaAlexia/pixiv-api-go/pixiv/pixivstruct"
	"github.com/VeronicaAlexia/pixiv-api-go/pixiv/request"
	"math/rand"
	"net/url"
	"os/exec"
	"runtime"
	"strings"
)

// Generate a random token
func generateURLSafeToken(length int) string {
	str := "-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"
	sb := strings.Builder{}
	sb.Grow(length)
	for i := 0; i < length; i++ {
		sb.WriteByte(str[rand.Intn(len(str))])
	}
	return sb.String()
}

// S256 transformation method.
func s256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}

// Proof Key for Code Exchange by OAuth Public Clients (RFC7636).
func oauthPkce() (string, string) {
	codeVerifier := generateURLSafeToken(32)
	codeChallenge := s256(codeVerifier)
	return codeVerifier, codeChallenge
}

func get_pixiv_login_url() (string, string) {
	codeVerifier, codeChallenge := oauthPkce()
	urlValues := url.Values{
		"code_challenge":        {codeChallenge},
		"code_challenge_method": {"S256"},
		"client":                {"pixiv-android"},
	}
	return codeVerifier, "https://app-api.pixiv.net/web/v1/login" + "?" + urlValues.Encode()
}

func loginPixiv(Verifier string, code string) (*pixivstruct.AccessToken, error) {
	params := map[string]string{
		"client_id":      "MOBrBDS8blbauoSck0ZfDbtuzpyT",
		"client_secret":  "lsACyCD94FhDUtGTXi3QzcFE2uU1hqtDaKeqrdwj",
		"code":           code,
		"code_verifier":  Verifier,
		"grant_type":     "authorization_code",
		"include_policy": "true",
		"redirect_uri":   "https://app-api.pixiv.net/web/v1/users/auth/pixiv/callback",
	}
	response := request.Post("https://oauth.secure.pixiv.net/auth/token", params).Json(&pixivstruct.AccessToken{}).(*pixivstruct.AccessToken)
	if response.AccessToken == "" {
		return nil, fmt.Errorf("login login pixiv error: %s", response.Error)
	} else {
		return response, nil
	}
}

func openbrowser(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	return err
}
func ChromeDriverLogin() (*pixivstruct.AccessToken, error) {
	codeVerifier, loginURL := get_pixiv_login_url() // Get the login URL and code verifier
	fmt.Println("please open the following link in your browser:", loginURL)
	fmt.Println("please press f12 to open the developer console, and switch to the network tab.")
	fmt.Println("now, please send the value of the code parameter in the request url of the remaining request.")
	fmt.Println("after logging in, please enter the code value:")
	fmt.Println("note that the code has a very short lifetime, please make sure that the previous step is completed quickly.")
	if err := openbrowser(loginURL); err != nil {
		fmt.Println("failed to open browser, please open the following link in your browser:", loginURL)
	} else {
		fmt.Printf("browser opened successfully,please input the code value:")
	}
	// input code
	var code string
	fmt.Printf("code:")
	_, err := fmt.Scanln(&code)
	if err != nil {
		fmt.Println("input error:", err)
	} else {
		return loginPixiv(codeVerifier, code)
	}
	return nil, err
}
