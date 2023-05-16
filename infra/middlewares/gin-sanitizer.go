package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Sanitizer interface {
	Sanitize(context *gin.Context, responseBody any) error
}

type SanitizerResponseWriter struct {
	gin.ResponseWriter
	body any
}

func (w *SanitizerResponseWriter) Write(b []byte) (int, error) {
	err := json.Unmarshal(b, &w.body)
	if err != nil {
		return 0, err
	}
	return len(b), nil
}

func (w *SanitizerResponseWriter) UpdateResponseBody() (int, error) {
	bodyBytes, err := json.Marshal(w.body)
	if err != nil {
		return 0, err
	}
	return w.ResponseWriter.Write(bodyBytes)
}

func Sanitization(sanitizers ...Sanitizer) gin.HandlerFunc {
	return func(context *gin.Context) {
		sanitizerResponseWriter := &SanitizerResponseWriter{context.Writer, nil}
		context.Writer = sanitizerResponseWriter
		context.Next()

		for _, sanitizer := range sanitizers {
			err := sanitizer.Sanitize(context, sanitizerResponseWriter.body)
			if err != nil {
				context.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		}

		sanitizerResponseWriter.UpdateResponseBody()
	}
}
