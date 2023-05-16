package middleware

import (
	"github.com/gin-gonic/gin"
)

var piiFields = map[string]MaskerFunc{
	"ssn": func(field string, skipSanitization bool) string {
		if skipSanitization {
			return field
		}
		return "***-**-**" + field[len(field)-2:]
	},
	"firstName": func(field string, skipSanitization bool) string {
		if skipSanitization {
			return field
		}
		return "***-**-****"
	},
	"lastName": func(field string, skipSanitization bool) string {
		if skipSanitization {
			return field
		}
		return "***-**-****"
	},
}

type PIISanitizer struct {
}

func (s *PIISanitizer) Sanitize(c *gin.Context, data any) error {
	for field, maskerFunc := range piiFields {
		headerField := c.Query(field)
		SanitizeResponseBody(data, field, maskerFunc, headerField != "")
	}
	return nil
}
