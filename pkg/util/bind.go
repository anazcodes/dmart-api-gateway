package util

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BindRequest(c *gin.Context, obj any) bool {

	err := c.ShouldBindJSON(obj)
	if err != nil {
		response := Response(http.StatusBadRequest, "failed", nil, nil)
		c.JSON(http.StatusBadRequest, &response)
		return false
	}
	return err == nil
}

func GetPageNCount(c *gin.Context) (page int, count int, err error) {
	page, err = strconv.Atoi(c.Query("page"))
	if err != nil {
		return 0, 0, err
	}

	count, err = strconv.Atoi(c.Query("count"))
	if err != nil {
		return 0, 0, err
	}
	return
}
