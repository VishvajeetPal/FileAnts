package main

import (
	"FileAnts/api"
	repository "FileAnts/repository/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"

	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		return
	}
	dbUrl := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	DB_URL := "postgres://" + dbName + ":" + dbPassword + "@" + dbUrl

	dialector := postgres.Open(DB_URL)
	db := repository.NewDatabse(dialector)
	db.Connect()

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	routersInit := api.InitRouter()
	gin.SetMode(gin.DebugMode)
	routersInit.Run(port)
}
