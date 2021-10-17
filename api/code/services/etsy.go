package services

import (
	"encoding/base64"
	"log"
	"math/rand"
	"net/http"

	cv "github.com/nirasan/go-oauth-pkce-code-verifier"
	"golang.org/x/oauth2"

	"github.com/deer-woman-dezigns/deer-woman-dezigns/code/config"
	"github.com/gin-gonic/gin"
)

var CodeVerifier, _ = cv.CreateCodeVerifier()

type EtsyService struct {
	EtsyOauthConfig oauth2.Config
}

func NewEtsyService() *EtsyService {
	return &EtsyService{
		EtsyOauthConfig: oauth2.Config{
			RedirectURL: "https://backend.deerwoman-dezigns.com/api/v1/etsy/callback/",
			ClientID:    config.Config.EtsyClientId,
			Scopes:      []string{"shops_r"},
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://www.etsy.com/oauth/connect/",
				TokenURL: "https://openapi.etsy.com/v2/oauth/token/",
			},
		},
	}
}

func (s *EtsyService) Login(c *gin.Context) {

	stateCookie := s.GenerateStateCookie(c)
	codeChallenge := CodeVerifier.CodeChallengeS256()
	challengeOpt := oauth2.SetAuthURLParam("code_challenge", codeChallenge)
	challengeTypeOpt := oauth2.SetAuthURLParam("code_challenge_method", "S256")
	redirectUrl := s.EtsyOauthConfig.AuthCodeURL(stateCookie, challengeOpt, challengeTypeOpt)
	c.Redirect(http.StatusTemporaryRedirect, redirectUrl)
	return
}

func (s *EtsyService) HandleCallback(c *gin.Context) string {
	tokenState, _ := c.Cookie("oauthstate")

	if reqState := c.Query("state"); reqState == "" || reqState != tokenState {
		log.Println("invalid or missing state token")
		return "invalid or missing state token"
	}
	code := c.Query("code")
	token := s.GetAuthToken(c, code)

	return token
}

func (s *EtsyService) GetAuthToken(c *gin.Context, code string) string {
	if token, err := s.EtsyOauthConfig.Exchange(c, code); err != nil {
		log.Println("error retrieving oath token", err)
		return ""
	} else {
		return token.AccessToken
	}
}

func (s *EtsyService) GenerateStateCookie(c *gin.Context) string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	c.SetCookie("oauthstate", state, 60*60*12, "/", "backend.deerwoman-dezigns.com", false, true)

	return state
}
