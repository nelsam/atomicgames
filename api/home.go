package api

import (
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"github.com/stretchr/goweb/handlers"
)

type HomeController struct {
	BaseRestController
}

func (controller *HomeController) Path() string {
	return "/home"
}

func (controller *HomeController) Versions() []string {
	return []string{"1.0"}
}

func (controller *HomeController) MatchAccept(ctx context.Context) (handlers.MatcherFuncDecision, error) {
	decision, err := controller.BaseRestController.MatchAccept(ctx)
	if err == nil && decision != handlers.NoMatch {
		if !controller.MatchVersions(controller.Versions()) {
			decision = handlers.NoMatch
		}
	}
	return decision, err
}

func (controller *HomeController) ReadMany(ctx context.Context) error {
	manyHomes := []string{"Test", "test2"}
	return goweb.API.RespondWithData(ctx, manyHomes)
}
