package user

import (
	"FileAnts/model"
	repository "FileAnts/repository/db"
	"FileAnts/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func DownloadFile(c *gin.Context) (string, string, string, error) {

	var file model.Aws_File

	id := c.Param("id")
	file.AwsLink = id

	res, err := repository.GetOne(file)

	if err != nil {
		return "", "", "", err
	}

	err = utils.S3FileDownloader(res.AwsLink)
	if err != nil {
		fmt.Printf("%v", err)
		return "", "", "", err
	}

	err = utils.Decrypt("./recordsTemp/download/"+res.AwsLink, res.AwsLink+"."+res.Extension)
	if err != nil {
		fmt.Printf("%v", err)
		return "", "", "", err
	}

	return "./recordsTemp/download/", res.AwsLink + res.Extension, "./recordsTemp/download/" + res.AwsLink, nil
}
