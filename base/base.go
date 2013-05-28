package base

import (
    "net/http"
)

type PathHandler struct {
    Path string
    Handler http.Handler
}