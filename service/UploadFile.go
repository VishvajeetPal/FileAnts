package user

import (
	"FileAnts/model"
	repository "FileAnts/repository/db"
	"FileAnts/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
	"path/filepath"
	"time"
)

func UploadFile(c *gin.Context) (interface{}, error) {

	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return nil, err
	}

	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String()

	if err := c.SaveUploadedFile(file, "./recordsTemp/upload/"+newFileName+extension); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return nil, err
	}

	err = utils.Encrypt("./recordsTemp/upload/"+newFileName+extension, newFileName)
	if err != nil {
		return nil, err
	}
	err = utils.AwsFileUpload("./recordsTemp/upload/"+newFileName+".bin", newFileName)
	if err != nil {
		return nil, err
	}

	record := model.Aws_File{
		AwsLink:   newFileName,
		FileId:    uuid.New().String(),
		Extension: extension,
	}

	max := c.Param("min")

	if max == "" {
		return nil, fmt.Errorf("max time param is empty")
	}

	if max == "5" {
		record.MaxTime = time.Now().Add(5 * time.Minute).Unix()
	} else if max == "10" {
		record.MaxTime = time.Now().Add(10 * time.Minute).Unix()
	} else if max == "30" {
		record.MaxTime = time.Now().Add(30 * time.Minute).Unix()
	} else {
		record.MaxTime = time.Now().Add(30 * time.Second).Unix()
	}

	if err := validator.New().Struct(record); err != nil {
		return nil, err
	}
	_, err = repository.CreateOne(&record)
	if err != nil {
		return nil, err
	}

	err = utils.DeleteFiles("./recordsTemp/upload/" + newFileName + extension)
	if err != nil {
		return nil, err
	}
	err = utils.DeleteFiles("./recordsTemp/upload/" + newFileName + ".bin")
	if err != nil {
		return nil, err
	}
	return c.Request.Host + "/api/v1/user/download/" + record.AwsLink, nil
	/*	return "Your file has been successfully uploaded.", nil
	 */
}
