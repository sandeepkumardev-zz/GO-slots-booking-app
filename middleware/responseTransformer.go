package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ResponseTransformer struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func EnsureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		getInfo := c.Request.Header["Authorization"][0]
		loggedIn, _ := strconv.ParseBool(getInfo)
		if !loggedIn {
			c.AbortWithStatusJSON(http.StatusUnauthorized, &ResponseTransformer{Message: "User is not logged in!", Data: nil, Success: false})
		}
		c.Next()
	}
}
