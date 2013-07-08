package api

import (
	"net/http"
	"github.com/stretchr/goweb"
)

func init() {
	var homeController RestController = new(HomeController)
	http.Handle("/", goweb.DefaultHttpHandler())
	goweb.MapController(homeController, homeController.MatchAccept)
}