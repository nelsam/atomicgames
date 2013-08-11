package api

import (
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/controllers"
)

var ControllerMap = map[string]controllers.RestfulController{
	"home": new(HomeController),
}

func MapRoutes() {
	for _, controller := range ControllerMap {
		if restController, ok := controller.(RestController); ok {
			// We only want to map API controllers that implement our
			// requirements for REST controllers.
			goweb.MapController(controller, restController.MatchAccept)
		}
	}
}
