// +build !appengine

// This is mainly to take care of dependency handling when using the
// "go get" command.  However, if we ever stop using GAE, it will
// be useful for generating a binary to run as a web server, too.
package atomicgames

import (
	"github.com/nelsam/atomicgames/api"
	"github.com/stretchr/goweb"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	api.MapRoutes()

	address := ":" + os.Getenv("PORT")
	server := &http.Server{
		Addr:           address,
		Handler:        goweb.DefaultHttpHandler(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	listener, listenErr := net.Listen("tcp", address)
	if listenErr != nil {
		panic(listenErr)
	}

	server.Serve(listener)
}
