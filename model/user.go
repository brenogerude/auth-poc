package model

import (
	"fmt"
	"strings"
)

type UserRole string

const (
	AdminRole   UserRole = "admin"
	SupportRole UserRole = "support"
	ClientRole  UserRole = "client"
)

func GetUserRoleFromString(role string) (UserRole, error) {
	switch strings.ToLower(role) {
	case "admin":
		return AdminRole, nil
	case "support":
		return SupportRole, nil
	case "client":
		return ClientRole, nil
	default:
		return "", fmt.Errorf("invalid user role %s", role)
	}
}
