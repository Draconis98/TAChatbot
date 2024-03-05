package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"io"
	"log"
	"net/http"
	"os"
)

var github_config *oauth2.Config

func getGitHubOAuthConfig() *oauth2.Config {
	github_config = &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),     // "GithubID"
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"), // "GithubSecret"
		Scopes:       []string{"read:user", "user:email"},
		Endpoint:     github.Endpoint,
		//RedirectURL:  "http://127.0.0.1:9000",
	}

	return github_config
}

func GitHubOAuth(r *gin.Context) {
	config := getGitHubOAuthConfig()
	url := config.AuthCodeURL("state")
	log.Println("URL: ", url)
	r.JSON(http.StatusOK, gin.H{
		"url": url,
	})
}

func GitHubCallback(r *gin.Context) {
	config := getGitHubOAuthConfig()
	code := r.Query("code")
	token, err := config.Exchange(r, code)
	if err != nil {
		log.Println("Error during exchange: ", err)
		r.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	username, done := getInfo(r, token, "user", err)
	if done {
		return
	}

	email, done := getInfo(r, token, "email", err)
	if done {
		return
	}

	r.JSON(http.StatusOK, gin.H{
		"username": username,
		"email":    email,
	})
}

func getInfo(r *gin.Context, token *oauth2.Token, info string, err error) (string, bool) {
	client := github_config.Client(r, token)
	var resp *http.Response
	if info == "user" {
		resp, err = client.Get("https://api.github.com/user")

	} else {
		resp, err = client.Get("https://api.github.com/user/emails")
	}
	if err != nil {
		log.Println("Error during get: ", err)
		r.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return "", true
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

	_info, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error during read: ", err)
		r.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return "", true
	}

	if info == "user" {
		var user map[string]interface{}
		if err = json.Unmarshal(_info, &user); err != nil {
			log.Println("Error during unmarshal: ", err)
			r.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return "", true
		}
		return user["login"].(string), false
	} else {
		var emails []map[string]interface{}
		if err = json.Unmarshal(_info, &emails); err != nil {
			log.Println("Error during unmarshal: ", err)
			r.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return "", true
		}

		for _, e := range emails {
			if e["primary"].(bool) {
				return e["email"].(string), false
			}
		}
	}

	return "", false
}
