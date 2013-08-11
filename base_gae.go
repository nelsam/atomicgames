// +build appengine

package atomicgames

import (
	"github.com/nelsam/atomicgames/api"
	"github.com/stretchr/goweb"
	"net/http"
)

func init() {
	api.MapRoutes()
	http.Handle("/", goweb.DefaultHttpHandler())
}
