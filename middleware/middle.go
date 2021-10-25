package middleware

import (
	"bytes"
	"fmt"
	"io"

	"time"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	limiter "github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
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

func RateLimit() gin.HandlerFunc {
	rate, err := limiter.NewRateFromFormatted("2-H")
	if err != nil {
		panic(err)
	}

	store := memory.NewStore()

	middleware := mgin.NewMiddleware(limiter.New(store, rate))

	return middleware
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LimitExceed(c *gin.Context) {
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw
	c.Next()
	statusCode := c.Writer.Status()
	if statusCode >= 429 {
		fmt.Println("Over")
	}
}
