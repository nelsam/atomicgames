package api

import (
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"net/http"
	"path"
)

type HomeController struct {
	BaseRestController
	BaseAppengineController
}

func (controller *HomeController) Path() string {
	return path.Join(controller.BaseRestController.Path(), "home")
}

func (controller *HomeController) ReadMany(ctx context.Context) error {
	manyHomes := []string{"Test", "test2"}
	return goweb.API.WriteResponseObject(ctx, http.StatusOK, manyHomes)
}
