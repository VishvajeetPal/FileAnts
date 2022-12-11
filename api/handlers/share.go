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
		c.AbortWithStatusJSON(404, "Something went wrong")
	}
	c.FileAttachment(file, name)
	c.JSON(200, "ok")
}
