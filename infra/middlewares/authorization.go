package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

const BEARER_SCHEMA = "Bearer"

func authorization(header string, req *http.Request) (bool, error) {
	token := getToken(header)
	path := req.URL.Path
	session, err := getSessionInfo(token)
	if err != nil {
		return false, err
	}

	if session.Role.IsAdmin() {
		return true, nil
	}
	//authorizeRole(session)

	sessionOk, err := authorizeSession(session, path)
	if err != nil {
		return false, err
	}
	if !sessionOk {
		return false, nil
	}
	return true, nil
}

func getToken(header string) string {
	tokenString := header[len(BEARER_SCHEMA):]
	return tokenString
}

func authorizeSession(session Session, path string) (bool, error) {
	// /users/45s6df-4fs6d-5f4s/whatever...
	//    ^         ^
	//    |         |
	pices, err := getPathPieces(path)
	if err != nil {
		return false, err
	}

	// users <- first part
	noun := pices[0]

	// 45s6df-4fs6d-5f4s <- second part
	subject, err := getPathSubject(pices)
	if err != nil {
		return false, err
	}

	return authorizationsMethods[noun](session, subject)
}

func getPathPieces(path string) ([]string, error) {
	if len(path) == 0 {
		return nil, errors.New("empty path")
	}
	pieces := strings.Split(path, "/")
	return []string{pieces[1], pieces[2]}, nil
}

func getPathSubject(pieces []string) (string, error) {
	if len(pieces) <= 1 {
		return "", errors.New("no route subject")
	}

	if !IsValidUUID(pieces[1]) {
		return "", errors.New("invalid subject id")
	}

	return pieces[1], nil
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
