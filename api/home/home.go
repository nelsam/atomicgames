package home

import (
    "fmt"
    "api/base"
    //"appengine"
    "net/http"
)

type HomeHandler struct {
    base.BaseHandler
    Message string
}

func (self HomeHandler) Get(writer http.ResponseWriter,
                            request *http.Request) {
    fmt.Fprint(writer, self.Message)
}

func MakeHomeHandler(message string) *base.UrlHandler {
    homeHandler := HomeHandler{base.BaseHandler{}, "Herro"}
    return &base.UrlHandler{&homeHandler}
}

