package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if method == http.MethodPost || method == http.MethodPatch {
			log.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
			if len(c.Request.URL.Query()) > 0 {
				log.Printf("Query Params: %v", c.Request.URL.Query())
			}
			if strings.HasPrefix(c.ContentType(), "application/json") && c.Request.Body != nil {
				bodyBytes, err := io.ReadAll(c.Request.Body)
				if err != nil {
					log.Printf("Error reading body: %v", err)
				} else {
					var pretty bytes.Buffer
					if err := json.Indent(&pretty, bodyBytes, "", "  "); err == nil {
						log.Printf("JSON Body:\n%s", pretty.String())
					} else {
						log.Printf("Raw Body: %s", string(bodyBytes))
					}
					c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
				}
			}
		}

		c.Next()
	}
}
