package httpserver

import (
	"net/http"
	"time"

	"fossa/pkg/logging"

	"github.com/gin-gonic/gin"
)

const (
	XRequestIDHeader = "X-Request-Id"
)

func LoggerMiddleware(logger *logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		cntx := RequestContext{
			ID:        c.GetHeader(XRequestIDHeader),
			Timestamp: time.Now(),
			Method:    c.Request.Method,
			URL:       c.Request.URL.String(),
			Status:    c.Writer.Status(),
		}

		logger := logger.WithContext("request", cntx)

		if c.Writer.Status() == http.StatusInternalServerError {
			// All errors wrapped into single error with stacktrace
			logger.Error("request processing error", "error", c.Errors.Last().Err)
		} else {
			logger.Info("request processed")
		}
	}
}

type RequestContext struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"ts"`
	Method    string    `json:"method"`
	URL       string    `json:"url"`
	Status    int       `json:"status"`
}
