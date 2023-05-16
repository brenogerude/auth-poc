package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerFunc func(*gin.Context)
type Route struct {
	Method      string
	Path        string
	HandlerFunc gin.HandlerFunc
}

type RouterGroup []Route

func (rg RouterGroup) forEach(handler func(route Route)) {
	for _, r := range rg {
		handler(r)
	}
}

type Router interface {
	AddRoute(method string, path string, handlerFunc gin.HandlerFunc)
	//Use(http.HandlerFunc)
	Run(port string) error
}

func Get(path string, handler gin.HandlerFunc) Route {
	return Route{
		Method:      http.MethodGet,
		Path:        path,
		HandlerFunc: handler,
	}
}

func Post(path string, handler gin.HandlerFunc) Route {
	return Route{
		Method:      http.MethodPost,
		Path:        path,
		HandlerFunc: handler,
	}
}
