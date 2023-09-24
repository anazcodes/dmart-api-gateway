package util

import (
	"net/http"
	"strconv"

	"github.com/anazibinurasheed/dmart-api-gateway/pkg/auth-svc/payload"
	"github.com/gin-gonic/gin"
)

// BindRequest is for bind the data from api request. If bind is success it will return true if any error happens it will return false and also it will write appropriate response.(swag:status 400, util.response)
func BindRequest(c *gin.Context, obj any) bool {

	err := c.ShouldBindJSON(obj)
	if HasError(err) {
		response := Response(http.StatusBadRequest, "json binding failed", nil, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, &response)
		return false
	}

	return true
}

// ValidateStruct will validate the struct fields according to validate tags and if ok it will return true else it will return false. And it will also write appropriate response.(swag:status 417, util.response)
func ValidateStruct(c *gin.Context, obj any) bool {

	if err := payload.ValidateStruct(obj); err != nil {
		response := Response(http.StatusExpectationFailed, "failed, not enough credentials", nil, err.Error())
		c.AbortWithStatusJSON(http.StatusExpectationFailed, &response)
		return false
	}
	return true

}

// RpcHasError checks the error and if there is any error it will write the response and it will return true if there is an error else it will return false.(swag:status 502, util.response)
func RpcHasError(c *gin.Context, err error) bool {

	if HasError(err) {
		response := Response(http.StatusBadGateway, "rpc error", nil, err.Error())
		c.AbortWithStatusJSON(http.StatusBadGateway, &response)
		return true
	}

	return false
}

// ErrorInPageInfo will check the error, if error exist it will write appropriate response and it will return true . If no error it will return false. (swag:status 400, util.response)
func ErrorInPageInfo(c *gin.Context, err error) bool {

	if HasError(err) {
		response := Response(http.StatusBadRequest, "Invalid parameters provided for pagination", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return true
	}
	return false
}

// GetPageNCount will fetch the page and count from the request and it will return the converted values of it and if any error happens it will return an error.
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
