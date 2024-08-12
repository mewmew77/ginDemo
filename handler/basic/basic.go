package basic

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewBasicHandler(r gin.IRouter) {
	r.GET("basic", base)
}

func base(c *gin.Context) {
	c.JSON(http.StatusOK, "base success")
}
