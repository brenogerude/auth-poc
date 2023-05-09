package api

import (
	"encoding/json"
	"net/http"
	"reflect"
)

type Body struct {
	Content    interface{}
	Reflection reflect.Value
}

func BuildBody(content interface{}) Body {
	b, err := json.MarshalIndent(&content, "", "\t")
	if err != nil {
		panic(err)
	}
	v := reflect.ValueOf(content)
	return Body{
		Content:    b,
		Reflection: v,
	}
}

type Response struct {
	http.ResponseWriter
	Body      *Body
	Succeeded bool
}

func (r *Response) Success(body Body) {
	r.Body = &body
	r.Succeeded = true
	r.WriteHeader(200)
}

func (r *Response) JSON(body interface{}) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		http.Error(r, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	r.Header().Set("Content-Type", "application/json")

	// Write the JSON data to the response writer
	_, err = r.Write(jsonBytes)
	if err != nil {
		http.Error(r, err.Error(), http.StatusInternalServerError)
		return
	}
}

type HandlerFunc func(*http.Request, *Response)

type PreHandler func(HandlerFunc) HandlerFunc

type Route struct {
	Method      string
	Path        string
	HandlerFunc HandlerFunc
	PreHandlers []PreHandler
}

type RouterGroup []Route

type Router interface {
	AddRoute(method string, path string, handlerFunc http.HandlerFunc)
	//Use(http.HandlerFunc)
	Run(port string) error
}

func Get(path string, handler HandlerFunc, middleware ...PreHandler) Route {
	return Route{
		Method:      http.MethodGet,
		Path:        path,
		HandlerFunc: handler,
		PreHandlers: middleware,
	}
}

func Post(path string, handler HandlerFunc, middleware ...PreHandler) Route {
	return Route{
		Method:      http.MethodPost,
		Path:        path,
		HandlerFunc: handler,
		PreHandlers: middleware,
	}
}
