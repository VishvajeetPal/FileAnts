package api

import (
	"FileAnts/api/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	router.Use(cors.Default()) //allows all origin
	v1 := router.Group("/api/v1/")
	{

		users_api := v1.Group("/user/")
		{
			users_api.POST("/upload", handlers.UploadFile)
			users_api.GET("/download/:id", handlers.DownloadFile)

		}

	}

	return router
}
