package services

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"

	"github.com/deer-woman-dezigns/deer-woman-dezigns/code/models"

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
			RedirectURL: fmt.Sprintf("%s://%s/api/v1/etsy/callback", config.Config.BackendProtocol, config.Config.BackendDomain),
			ClientID:    config.Config.EtsyClientId,
			Scopes:      []string{"shops_r"},
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://www.etsy.com/oauth/connect",
				TokenURL: "https://api.etsy.com/v3/public/oauth/token",
			},
		},
	}
}

func (s *EtsyService) Login(c *gin.Context) string {
	stateCookie := s.GenerateStateCookie(c)
	codeChallenge := CodeVerifier.CodeChallengeS256()
	challengeOpt := oauth2.SetAuthURLParam("code_challenge", codeChallenge)
	challengeTypeOpt := oauth2.SetAuthURLParam("code_challenge_method", "S256")
	c.SetCookie("codeVer", CodeVerifier.Value, 60*60*12, "/", config.Config.BackendDomain, false, true)
	redirectUrl := s.EtsyOauthConfig.AuthCodeURL(stateCookie, challengeOpt, challengeTypeOpt)
	return redirectUrl
}

func (s *EtsyService) HandleCallback(c *gin.Context) {
	tokenState, _ := c.Cookie("oauthstate")

	if reqState := c.Query("state"); reqState == "" || reqState != tokenState {
		log.Println("invalid or missing state token")
	}
	code := c.Query("code")
	s.GetAuthToken(c, code)
}

func (s *EtsyService) GetAuthToken(c *gin.Context, code string) {
	codeVer, _ := c.Cookie("codeVer")
	if resp, err := http.PostForm(s.EtsyOauthConfig.Endpoint.TokenURL, url.Values{
		"grant_type":    {"authorization_code"},
		"client_id":     {config.Config.EtsyClientId},
		"redirect_uri":  {s.EtsyOauthConfig.RedirectURL},
		"code":          {code},
		"code_verifier": {codeVer},
	}); err != nil {
		log.Println("error retrieving oath token", err)
	} else {
		tokens := models.Tokens{}
		json.NewDecoder(resp.Body).Decode(&tokens)
		c.SetCookie("access_token", tokens.AccessToken, tokens.ExpiresIn, "/", config.Config.BackendDomain, false, true)
		c.SetCookie("refresh_token", tokens.RefreshToken, 60*60*24*7, "/", config.Config.BackendDomain, false, true)
	}
}

func (s *EtsyService) GenerateStateCookie(c *gin.Context) string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	c.SetCookie("oauthstate", state, 60*60*12, "/", config.Config.BackendDomain, false, true)
	return state
}
