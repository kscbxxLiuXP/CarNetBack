package controller

import (
	"github.com/gin-gonic/gin"
)

func NoRoute(c *gin.Context) {
	NOTFOUND(c)
}
