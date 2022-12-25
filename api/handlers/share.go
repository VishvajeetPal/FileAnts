package handlers

import (
	user "FileAnts/service"
	"FileAnts/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadFile(c *gin.Context) {
	res, err := user.UploadFile(c)
	fmt.Printf("Err %v\n", err)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func DownloadFile(c *gin.Context) {

	file, name, _, err := user.DownloadFile(c)
	fmt.Printf("\n%v\n", err)
	if err != nil {
		fmt.Printf("\n%v\n", err)
		c.AbortWithStatusJSON(404, "File has expired.")
	}

	fmt.Printf("\nFile = %v\n", file)
	fmt.Printf("Namee = %v", name)
	c.FileAttachment(file+name, name)
	//c.JSON(200, "ok")
}

func Clear(c *gin.Context) {
	res, err := user.Clear(c)
	fmt.Printf("Err %v\n", err)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}
