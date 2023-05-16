package infra

import (
	"fmt"
	"net/http"
	middleware "oauth-poc/infra/middlewares"

	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	*gin.Engine
}

func NewGinRouter() *GinRouter {
	engine := gin.Default()
	engine.Use(
		middleware.Sanitization(getSanitizers()),
		middleware.Authentication(),
		middleware.Authorization(),
	)
	return &GinRouter{engine}

}

func (r *GinRouter) AddRoute(method string, path string, handler gin.HandlerFunc) {
	r.Engine.Handle(method, path, handler)
}

func (r *GinRouter) Use(handlerFunc http.HandlerFunc) {
	r.Engine.Use(gin.WrapH(handlerFunc))
}

func (r *GinRouter) Run(addr string) error {
	fmt.Println("I'm using gin router")
	return r.Engine.Run(addr)
}

func getSanitizers() map[string]middleware.Sanitizer {
	return map[string]middleware.Sanitizer{
		"osv": &middleware.OSVSanitizer{},
	}
}
