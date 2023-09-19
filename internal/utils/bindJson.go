package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindRequest(c *gin.Context, obj any) {

	if err := c.ShouldBindJSON(obj); err != nil {
		response := Response(http.StatusBadRequest, "failed", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
}
