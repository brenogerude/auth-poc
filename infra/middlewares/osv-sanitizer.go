package middleware

import (
	"fmt"
	"oauth-poc/model"

	"github.com/gin-gonic/gin"
)

type fields []string

func (f fields) contains(s string) bool {
	for _, v := range f {
		if v == s {
			return true
		}
	}
	return false
}

var disallowedFields = map[model.UserRole]fields{
	model.AdminRole: {},
	model.ClientRole: {
		"ssn",
		"companyId",
	},
}

type OSVSanitizer struct {
}

func (s *OSVSanitizer) Sanitize(c *gin.Context, data any) error {
	role, err := getRole(c)
	if err != nil {
		return err
	}
	fields := disallowedFields[*role]
	for _, field := range fields {
		sanitize(data, field)
	}
	return nil
}

func sanitize(data any, field string) {
	if nestedMap, ok := data.(map[string]interface{}); ok {
		sanitizeField(nestedMap, field)
	} else if arrayData, ok := data.([]interface{}); ok {
		for _, value := range arrayData {
			if valueAsMap, ok := value.(map[string]interface{}); ok {
				sanitizeField(valueAsMap, field)
			}
		}
	}
}

func sanitizeField(data map[string]interface{}, field string) {
	for key, value := range data {
		if nestedMap, ok := value.(map[string]interface{}); ok {
			sanitizeField(nestedMap, field)
			if len(nestedMap) == 0 {
				delete(data, key)
			}
		} else if arrayData, ok := value.([]interface{}); ok {
			for _, value := range arrayData {
				if valueAsMap, ok := value.(map[string]interface{}); ok {
					sanitizeField(valueAsMap, field)
				}
			}
		} else {
			delete(data, field)
		}
	}
}

func getRole(c *gin.Context) (*model.UserRole, error) {
	role := c.GetHeader("role")
	if role == "" {
		return nil, fmt.Errorf("Role not found in request")
	}
	userRole, err := model.GetUserRoleFromString(role)
	if err != nil {
		return nil, err
	}
	return &userRole, nil
}
