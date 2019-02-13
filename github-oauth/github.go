package main

import "github.com/gin-gonic/gin"

type Github struct{
	appId string
	appSecret string
}

func (gh *Github) HandleOAuth() {

}

func main() {

	router := gin.Default()
	gh := Github{appId: "f906611190c44776b306",appSecret:"eac043d3ca94c4c89568f53bfb89d094106437eb"}
	oAuth := router.Group("/oauth")
	{oAuth.POST("/github", gh.HandleOauth())
		router.Run("localhost:8080")}
	//r := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}

