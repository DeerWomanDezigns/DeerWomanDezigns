package services

import (
	"bytes"
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
	body := models.AuthRequest{
		GrantType:    "authorization_code",
		ClientID:     config.Config.EtsyClientId,
		RedirectURI:  redirectUrl,
		Code:         code,
		CodeVerifier: codeVer,
	}
	jsonBody, _ := json.Marshal(body)
	if resp, err := http.Post("https://api.etsy.com/v3/public/oauth/token", "application/json", bytes.NewBuffer(jsonBody)); err != nil {
		log.Println("error retrieving oauth token", err)
		c.AbortWithError(http.StatusBadRequest, err)
	} else {
		tokens := models.Tokens{}
		json.NewDecoder(resp.Body).Decode(&tokens)
		c.SetCookie("access_token", tokens.AccessToken, tokens.ExpiresIn, "/", config.Config.BackendDomain, false, true)
		c.SetCookie("refresh_token", tokens.RefreshToken, 60*60*24*7, "/", config.Config.BackendDomain, false, true)
		c.JSON(http.StatusOK, "Tokens acquired")
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
