// +build appengine

package api

import (
	"appengine"
	"github.com/stretchr/goweb/context"
)

// A base controller class containing GAE-specific logic.
type BaseAppengineController struct {
	GAEContext appengine.Context
}

// Uses a Context argument (from the goweb.context package) to create
// the appengine context, then returns it.
func (controller *BaseAppengineController) AppengineContext(ctx context.Context) appengine.Context {
	controller.SetGAEContext(ctx)
	return controller.GAEContext
}

// Set the GAEContext member.
func (controller *BaseAppengineController) SetGAEContext(ctx context.Context) {
	if controller.GAEContext == nil {
		controller.GAEContext = appengine.NewContext(ctx.HttpRequest())
	}
}
