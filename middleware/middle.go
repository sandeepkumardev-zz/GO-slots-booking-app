package middleware

import (
	"io"
	"net/http"
	"strconv"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

var LogMiddleware = logger.SetLogger(
	logger.WithLogger(func(c *gin.Context, out io.Writer, latency time.Duration) zerolog.Logger {
		return zerolog.New(out).With().
			Str("path", c.Request.URL.Path).
			Str("Method", c.Request.Method).
			Dur("latency", latency).
			Logger()
	}),
)

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

func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.MaxAge = 12 * time.Hour
	return cors.New(config)
}
