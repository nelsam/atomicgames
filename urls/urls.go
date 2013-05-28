package urls

import (
    "net/http"
    "api"
    "base"
    "ui"
)

func PathPrefixHandle(prefix string, handlers []base.PathHandler) {
    for _, handler := range handlers {
        path := prefix + handler.Path
        http.Handle(path, handler.Handler)
    }
}

func init() {
    PathPrefixHandle("/api", api.Handlers)
    http.Handle("/", ui.Handler)
}