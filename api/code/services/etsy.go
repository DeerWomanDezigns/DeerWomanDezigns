package services

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/deer-woman-dezigns/deer-woman-dezigns/code/models"

	"golang.org/x/oauth2"

	"github.com/deer-woman-dezigns/deer-woman-dezigns/code/config"
	"github.com/gin-gonic/gin"
)

type EtsyService struct {
	EtsyOauthConfig oauth2.Config
}

func NewEtsyService() *EtsyService {
	return &EtsyService{}
}

func (s *EtsyService) GetAuthToken(c *gin.Context, code string, codeVer string, redirectUrl string) {
	if resp, err := http.PostForm("https://api.etsy.com/v3/public/oauth/token", url.Values{
		"grant_type":    {"authorization_code"},
		"client_id":     {config.Config.EtsyClientId},
		"redirect_uri":  {redirectUrl},
		"code":          {code},
		"code_verifier": {codeVer},
	}); err != nil {
		log.Println("error retrieving oauth token", err)
	} else {
		tokens := models.Tokens{}
		json.NewDecoder(resp.Body).Decode(&tokens)
		c.SetCookie("access_token", tokens.AccessToken, tokens.ExpiresIn, "/", config.Config.BackendDomain, false, true)
		c.SetCookie("refresh_token", tokens.RefreshToken, 60*60*24*7, "/", config.Config.BackendDomain, false, true)
	}
}

func (s *EtsyService) RefreshAuthToken(c *gin.Context, refreshToken string) {
	if resp, err := http.PostForm(s.EtsyOauthConfig.Endpoint.TokenURL, url.Values{
		"grant_type":    {"refresh_token"},
		"client_id":     {config.Config.EtsyClientId},
		"refresh_token": {refreshToken},
	}); err != nil {
		log.Println("error retrieving oath token", err)
	} else {
		tokens := models.Tokens{}
		json.NewDecoder(resp.Body).Decode(&tokens)
		c.SetCookie("access_token", tokens.AccessToken, tokens.ExpiresIn, "/", config.Config.BackendDomain, false, true)
		c.SetCookie("refresh_token", tokens.RefreshToken, 60*60*24*7, "/", config.Config.BackendDomain, false, true)
	}
}
