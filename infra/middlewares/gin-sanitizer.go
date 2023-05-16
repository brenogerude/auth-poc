package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Sanitizer interface {
	Sanitize(*gin.Context, any) error
}

type responseWriter struct {
	gin.ResponseWriter
	body any
}

func (w *responseWriter) Write(b []byte) (int, error) {
	err := json.Unmarshal(b, &w.body)
	if err != nil {
		return 0, err
	}
	return len(b), nil
}

func (w *responseWriter) Override() (int, error) {
	bytes, err := json.Marshal(w.body)
	if err != nil {
		return 0, err
	}
	return w.ResponseWriter.Write(bytes)
}

func Sanitization(sanitizers map[string]Sanitizer) gin.HandlerFunc {
	return func(c *gin.Context) {
		writer := &responseWriter{c.Writer, nil}
		c.Writer = writer
		c.Next()

		// imagine we'll dynamically get the source from somewhere
		source := "osv"
		sanitizer, found := sanitizers[source]
		if !found {
			c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Sanitizer %s not found!", source))
			return
		}
		err := sanitizer.Sanitize(c, writer.body)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		writer.Override()
	}
}
