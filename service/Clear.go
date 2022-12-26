package user

import (
	"github.com/gin-gonic/gin"
	"os"
)

func Clear(c *gin.Context) (interface{}, error) {

	err := os.RemoveAll("./app/recordsTemp/download/")
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll("./app/recordsTemp/download/", 777)
	if err != nil {
		return nil, err
	}

	return "Done", err

}
