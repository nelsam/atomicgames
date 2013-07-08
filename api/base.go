package api

import (
	"strings"
	"appengine"
	"github.com/nelsam/atomicgames"
	"github.com/stretchr/goweb/context"
	"github.com/stretchr/goweb/handlers"
)

// Our RESTful controllers have a MatchAccept method, matching the
// definition of a MatcherFunc from the goweb package, which reads
// the Accept header and makes sure that it matches a given pattern
// before serving a response.  They also need a list of supported
// API versions, which is required by the Versions() method.
// We recommend that the Versions() method is used in the
// MatchAccept method for versioning purposes.
type RestController interface {
	MatchAccept(context.Context) (handlers.MatcherFuncDecision, error)
	Versions() []string
}

// Implements some basic parts of a RestController,
// but does NOT implement RestController.  You must implement your own
// Version method, and you should use your Version method to extend
// BaseRestController.MatchAccept, matching your controller's version
// in the Accept header.
type BaseRestController struct {
	appengineContext appengine.Context
	requestedVersion string
	MatchedAccept string
}

// Uses a Context argument (from the goweb.context package) to create
// the appengine context, which can be used for logging and various
// other things related to the appengine server.
func (controller *BaseRestController) AppengineContext(ctx context.Context) appengine.Context {
	if controller.appengineContext == nil {
		controller.appengineContext = appengine.NewContext(ctx.HttpRequest())
	}
	return controller.appengineContext
}

// Makes sure that the Vary header is set to Accept
func (controller *BaseRestController) After(ctx context.Context) error {
	ctx.HttpResponseWriter().Header().Set("Vary", "Accept")
	return nil
}

// Uses BaseRestController.MatchedAccept to find the requested API version.
func (controller *BaseRestController) RequestedVersion() string {
	if controller.requestedVersion == "" {
		version_end := strings.LastIndex(controller.MatchedAccept, "+")
		acceptToMatch := controller.MatchedAccept[:version_end]
		
		version_start := strings.LastIndex(acceptToMatch, "-") + 1
		version := acceptToMatch[version_start:]
		
		version = strings.TrimPrefix(version, "version")
		version = strings.TrimPrefix(version, "v")
		version = strings.TrimSpace(version)
		
		controller.requestedVersion = version
	}
	return controller.requestedVersion
}

// Sugar to check a list of versions (such as what would
// be returned by RestController.Versions()) against the requested
// API version.
func (controller *BaseRestController) MatchVersions(versions []string) bool {
	found := false
	version := controller.RequestedVersion()
	for _, supportedVersion := range versions {
		if version == supportedVersion {
			found = true
			break
		}
	}
	return found
}

// Searches for an entry in the Accept header that matches
// atomicgames.ApiResponseType.  If a matching entry is found, it is stored
// as BaseRestController.MatchedAccept.
func (controller *BaseRestController) MatchAccept(ctx context.Context) (handlers.MatcherFuncDecision, error) {
	accept := ctx.HttpRequest().Header.Get("Accept")
	matchedAcceptIndex := strings.Index(accept, atomicgames.ApiResponseType)
	switch matchedAcceptIndex {
	case -1:
		return handlers.NoMatch, nil
	default:
		acceptFromMatch := accept[matchedAcceptIndex:]
		var matchedAcceptEnd int
		if matchedAcceptEnd = strings.Index(acceptFromMatch, ","); matchedAcceptEnd == -1 {
			matchedAcceptEnd = len(acceptFromMatch)
		}
		matchedAccept := acceptFromMatch[:matchedAcceptEnd]
		if controller.MatchedAccept != matchedAccept {
			controller.MatchedAccept = matchedAccept
			controller.requestedVersion = ""
		}
		return handlers.DontCare, nil
	}
}
