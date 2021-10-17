package services

import (
	"encoding/base64"
	"log"
	"math/rand"
	"net/http"

	"golang.org/x/oauth2"

	"github.com/deer-woman-dezigns/deer-woman-dezigns/code/config"
	"github.com/gin-gonic/gin"
)

var etsyOauthConfig = &oauth2.Config{
	RedirectURL: "https://deerwoman-dezigns/api/v1/etsy/callback",
	ClientID:    config.Config.EtsyClientId,
	Scopes:      []string{"shops_r"},
	Endpoint: oauth2.Endpoint{
		AuthURL:  config.Config.EtsyRequestUrl,
		TokenURL: config.Config.EtsyAccessTokenUrl,
	},
}

type EtsyService struct{}

func NewEtsyService() *EtsyService {
	return &EtsyService{}
}

func (s *EtsyService) Login(c *gin.Context) {

	stateCookie := s.GenerateStateCookie(c)
	redirectUrl := etsyOauthConfig.AuthCodeURL(stateCookie)
	c.Redirect(http.StatusTemporaryRedirect, redirectUrl)
	return
}

func (s *EtsyService) HandleCallback(c *gin.Context) string {
	tokenState, _ := c.Cookie("oauthstate")

	if reqState := c.Param("token"); reqState != tokenState {
		log.Println("invalid or missing state token")
		return ""
	}
	code := c.Param("code")
	token := s.GetAuthToken(c, code)

	return token
}

func (s *EtsyService) GetAuthToken(c *gin.Context, code string) string {
	if token, err := etsyOauthConfig.Exchange(c, code); err != nil {
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
	c.SetCookie("oauthstate", state, 60*60*12, "/", "etsy.com", true, false)

	return state
}
