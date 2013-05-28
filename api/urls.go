package api

import (
    "api/home"
    "base"
)

var Handlers = []base.PathHandler{
    base.PathHandler{"/home", home.MakeHomeHandler("Herro.")},
}