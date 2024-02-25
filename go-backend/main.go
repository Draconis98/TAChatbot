package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-backend/api"
	"go-backend/internal/repository"
	"log"
)

func main() {
	router := gin.Default()

	// create a new repository
	mongoClient, err := repository.Connect2MongoDB()
	if err != nil {
		log.Panic(err)
	}
	defer repository.Disconnect2MongoDB(mongoClient, nil)
	repository.DB = repository.SelectDB(mongoClient, "llm")

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:8081",
			"http://127.0.0.1:8081",
			"https://10.30.19.40:8081",
			"https://callme.agileserve.org.cn:30940",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	//router.Use(LoadTls())

	// register the routes
	// TODO: auth这部分，现在如果用户碰巧手动输入了一个正确的userid的话，会进到别的用户的界面中
	router.GET("/auth", api.GitLabOAuth)

	callbackGroup := router.Group("/callback")
	{
		callbackGroup.GET("/github", api.GitHubCallback)
		callbackGroup.GET("/gitlab", api.GitLabCallback)
	}

	checkGroup := router.Group("/check")
	{
		checkGroup.GET("/user", api.CheckUserExistence)
		//checkGroup.GET("/card", api.CheckCardExistence)
	}

	getGroup := router.Group("/get")
	{
		getGroup.GET("/username", api.GetUsername)
		getGroup.GET("card", api.GetCard)
		getGroup.GET("/question", api.GetQuestion)
		getGroup.GET("isdisplay", api.IsDisplay)
	}

	newGroup := router.Group("/new")
	{
		newGroup.GET("/card", api.NewCard)
		newGroup.GET("/question", api.NewQuestion)
	}

	showGroup := router.Group("/show")
	{
		showGroup.GET("/", api.Show)
		showGroup.GET("/latest", api.ShowLatest)
		showGroup.GET("/hottest", api.ShowHottest)
	}

	//if err := router.RunTLS("10.30.19.40:8080", "/home/jss40/.vite-plugin-mkcert/rootCA.pem", "/home/jss40/.vite-plugin-mkcert/rootCA-key.pem"); err != nil {
	//	log.Panic(err)
	//}

	if err := router.Run("10.30.19.40:8080"); err != nil {
		log.Panic(err)
	}
}

//func LoadTls() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		middleware := secure.New(secure.Options{
//			SSLRedirect: true,
//			SSLHost:     "10.30.19.40:8080",
//		})
//		err := middleware.Process(c.Writer, c.Request)
//		if err != nil {
//			log.Println(err)
//			return
//		}
//	}
//}
