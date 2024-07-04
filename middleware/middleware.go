package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fields := logrus.Fields{}
		startTime := time.Now()
		traceID := c.GetHeader("X-Request-ID")
		if traceID == "" {
			traceID = uuid.New().String()
		}
		c.Set("traceID", traceID)

		fields["traceID"] = traceID
		fields["method"] = c.Request.Method
		fields["path"] = c.Request.URL.Path
		if len(c.Request.URL.RawQuery) > 0 {
			fields["raw_query"] = c.Request.URL.RawQuery
		}

		if c.ContentType() == gin.MIMEJSON {
			buf, _ := c.GetRawData()
			c.Request.Body = io.NopCloser(bytes.NewBuffer(buf))

			var body interface{}
			_ = json.Unmarshal(buf, &body)
			bodyBytes, _ := json.Marshal(body)
			fields["request_body"] = string(bodyBytes)
		}

		defer func() {
			if r := recover(); r != nil {
				logrus.WithFields(logrus.Fields{
					"traceID": traceID,
					"error":   r,
				}).Error("[HTTP Server] Recovered from a panic")
			}
		}()

		c.Next()

		fields["latency"] = time.Since(startTime).Milliseconds()
		status := c.Writer.Status()
		fields["status_code"] = status

		if status >= 500 {
			fields["error"] = c.Errors.String()
			logrus.WithFields(fields).Error("[HTTP Server] Request completed")
			return
		} else if status >= 400 {
			fields["error"] = c.Errors.String()
			logrus.WithFields(fields).Warn("[HTTP Server] Request completed")
			return
		} else {
			logrus.WithFields(fields).Info("[HTTP Server] Request completed")
			return
		}
	}
}
