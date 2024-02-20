package api

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

var gitlab_config *oauth2.Config

var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://gitlab.agileserve.org.cn:8001/oauth/authorize",
	TokenURL: "https://gitlab.agileserve.org.cn:8001/oauth/token",
}

func getGitLabOAuthConfig() string {
	gitlab_config = &oauth2.Config{
		ClientID:     "fde4604874a9d50ba861871b344095b3c8c55af513d745d371b0f2107d45f3ef",             // "GithubID"
		ClientSecret: "gloas-gloas-40765d36cb320010098669d88cb29845ea819e51eb38935099b98c062d47b436", // "GithubSecret"
		Scopes:       []string{"read_user"},
		Endpoint:     Endpoint,
		RedirectURL:  "http://127.0.0.1:8081/callback/gitlab",
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
	log.Println(code)

	_, err := gitlab_config.Exchange(r, code)
	if err != nil {
		log.Println("Error during exchange: ", err)
		r.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	//username, done := getInfo(r, token, "user", err)
	//if done {
	//	return
	//}
	//
	//email, done := getInfo(r, token, "email", err)
	//if done {
	//	return
	//}

	r.JSON(http.StatusOK, gin.H{
		"username": "username",
		"email":    "email",
	})
}

//func getInfo(r *gin.Context, token *oauth2.Token, info string, err error) (string, bool) {
//	client := gitlab_config.Client(r, token)
//	var resp *http.Response
//	resp, err = client.Get("https://gitlab.agileserve.org.cn:8001/api/v4/user")
//	if err != nil {
//		log.Println("Error during get: ", err)
//		r.JSON(http.StatusUnauthorized, gin.H{
//			"error": err.Error(),
//		})
//		return "", true
//	}
//	print(resp.Body)
//	return "", false
//}
