package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"io"
	"log"
	"net/http"
)

var gitlab_config *oauth2.Config

//var Endpoint =

func getGitLabOAuthConfig() string {
	gitlab_config = &oauth2.Config{
		ClientID:     "fde4604874a9d50ba861871b344095b3c8c55af513d745d371b0f2107d45f3ef",       // "GithubID"
		ClientSecret: "gloas-7c56cc2423a1db1644c94b0000b2afd995dbc5d9ad2e3bfeba089bef4987e520", // "GithubSecret"
		Scopes:       []string{"read_user"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://gitlab.agileserve.org.cn:8001/oauth/authorize",
			TokenURL: "https://gitlab.agileserve.org.cn:8001/oauth/token",
		},
		RedirectURL: "http://127.0.0.1:8081/callback/gitlab",
	}

	return gitlab_config.AuthCodeURL("callme")
}

func GitLabOAuth(r *gin.Context) {
	url := getGitLabOAuthConfig()
	log.Println("URL: ", url)
	r.JSON(http.StatusOK, gin.H{
		"url": url,
	})
}

func GitLabCallback(r *gin.Context) {
	code := r.Query("code")

	token, err := gitlab_config.Exchange(r, code)
	if err != nil {
		log.Println("Error during exchange: ", err)
		r.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	client := gitlab_config.Client(r, token)
	resp, err := client.Get("https://gitlab.agileserve.org.cn:8001/api/v4/user")
	if err != nil {
		log.Println("Error during get: ", err)
		r.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
	}

	info, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error during read: ", err)
		r.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
	}

	var user map[string]interface{}
	if err = json.Unmarshal(info, &user); err != nil {
		log.Println("Error during unmarshal: ", err)
		r.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
	}

	r.JSON(http.StatusOK, gin.H{
		"username": user["username"],
		"email":    user["email"],
	})
}
