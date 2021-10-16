package services

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/deer-woman-dezigns/deer-woman-dezigns/code/config"
	cv "github.com/nirasan/go-oauth-pkce-code-verifier"
)

var CodeVerifier, _ = cv.CreateCodeVerifier()

type EtsyService struct{}

func NewEtsyService() *EtsyService {
	return &EtsyService{}
}

func (s *EtsyService) Login(csrfToken string) {
	if req, err := http.NewRequest("GET", config.Config.EtsyRequestUrl, nil); err != nil {
		fmt.Println("Error forming Etsy login request", err)
	} else {
		q := req.URL.Query()
		q.Add("response_type", "code")
		q.Add("client_id", config.Config.EtsyClientId)
		q.Add("redirect_uri", "https://deerwoman-dezigns/api/v1/etsy/callback")
		q.Add("scope", "shops_r")
		q.Add("state", "")
		q.Add("code_challenge", CodeVerifier.CodeChallengeS256())
		q.Add("code_challenge_method", "S256")
		req.URL.RawQuery = q.Encode()
		fmt.Println(req.URL.String())

		if resp, err := http.Get(req.URL.String()); err != nil {
			fmt.Println("Error getting Etsy login", err)
		} else {
			fmt.Println(resp)
		}
	}

	return
}

func (s *EtsyService) HandleCallback(code string) {

	return
}

func (s *EtsyService) RandState(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		letterIdxBits = 6
		letterIdxMask = 1<<letterIdxBits - 1
		letterIdxMax  = 63 / letterIdxBits
	)
	sb := strings.Builder{}
	sb.Grow(n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}
