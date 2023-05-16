package middleware

import (
	"github.com/gin-gonic/gin"
)

var providerFields = map[string]MaskerFunc{
	"gross": func(field string, skipSanitization bool) string {
		if skipSanitization {
			return field
		}
		return "[Secured information]"
	},
}

type ProviderSanitizer struct {
}

func (s *ProviderSanitizer) Sanitize(c *gin.Context, data any) error {
	for field, maskerFunc := range providerFields {
		SanitizeResponseBody(data, field, maskerFunc, false)
	}
	return nil
}
