package api

import (
	"github.com/nelsam/atomicgames/settings"
	"github.com/stretchr/goweb/context"
	"github.com/stretchr/goweb/handlers"
	"strings"
)

// The base path for all API calls.  Every API call should
// be in a sub-path of this path.
const baseApiPath string = "/api"

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
	requestedVersion string
	allowedVersions  []string
	MatchedAccept    string
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
		if version_end == -1 {
			version_end = len(controller.MatchedAccept)
		}
		acceptToMatch := controller.MatchedAccept[:version_end]

		version := strings.Replace(acceptToMatch, settings.ApiResponseType, "", 1)

		version = strings.TrimPrefix(version, "-")
		version = strings.TrimPrefix(version, "version")
		version = strings.TrimPrefix(version, "v")
		version = strings.TrimSpace(version)

		if version == "" {
			controller.requestedVersion = "stable"
		} else {
			controller.requestedVersion = version
		}
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

// The base case for all RESTful controllers is that they match
// the "stable" version - which can be explicitly requested
// in the Accept header, or will be the default for any requests
// without an explicit version.
func (controller *BaseRestController) Versions() []string {
	if controller.allowedVersions == nil {
		controller.allowedVersions = []string{"stable"}
	}
	return controller.allowedVersions
}

// The base path for API calls is returned here.  API
// controllers should override this method, appending their
// relative path to the base path.
func (baseHandler *BaseRestController) Path() string {
	return baseApiPath
}

// Searches for an entry in the Accept header that matches
// settings.ApiResponseType.  If a matching entry is found, it is stored
// as BaseRestController.MatchedAccept.
func (controller *BaseRestController) MatchAccept(ctx context.Context) (handlers.MatcherFuncDecision, error) {
	accept := ctx.HttpRequest().Header.Get("Accept")
	matchedAcceptIndex := strings.Index(accept, settings.ApiResponseType)
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
