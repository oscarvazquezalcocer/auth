package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = oauth2.Config{
		ClientID:     "926578288193-101dm849obhuhdu460s1csr7e5h36j6u.apps.googleusercontent.com",
		ClientSecret: "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",
		//RedirectURL:  "http://localhost:8081/callback",
		Scopes:   []string{"openid", "email", "profile"},
		Endpoint: google.Endpoint,
	}
)

func Login(c *gin.Context) {
	googleOauthConfig.RedirectURL = c.DefaultQuery("redirect_uri", "/")
	url := googleOauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusFound, url)
}

func Root(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
