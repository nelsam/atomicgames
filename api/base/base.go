package base

import (
    "fmt"
    "net/http"
)

type MethodHandler interface {
    Get(http.ResponseWriter, *http.Request)
    Post(http.ResponseWriter, *http.Request)
    Put(http.ResponseWriter, *http.Request)
    Delete(http.ResponseWriter, *http.Request)
}

type UrlHandler struct {
    Methods MethodHandler
}

func (self UrlHandler) ServeHTTP(writer http.ResponseWriter,
                                 request *http.Request) {
    switch request.Method {
    case "GET":
        self.Methods.Get(writer, request)
    case "POST":
        self.Methods.Post(writer, request)
    case "PUT":
        self.Methods.Put(writer, request)
    case "DELETE":
        self.Methods.Delete(writer, request)
    }
}

type BaseHandler struct {
}

func (self BaseHandler) Get(writer http.ResponseWriter,
                            request *http.Request) {
    fmt.Fprint(writer, "Not Implemented")
}

func (self BaseHandler) Post(writer http.ResponseWriter,
                             request *http.Request) {
    fmt.Fprint(writer, "Not Implemented")
}

func (self BaseHandler) Put(writer http.ResponseWriter,
                            request *http.Request) {
    fmt.Fprint(writer, "Not Implemented")
}

func (self BaseHandler) Delete(writer http.ResponseWriter,
                               request *http.Request) {
    fmt.Fprint(writer, "Not Implemented")
}