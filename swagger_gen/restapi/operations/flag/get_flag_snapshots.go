// Code generated by go-swagger; DO NOT EDIT.

package flag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetFlagSnapshotsHandlerFunc turns a function with the right signature into a get flag snapshots handler
type GetFlagSnapshotsHandlerFunc func(GetFlagSnapshotsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFlagSnapshotsHandlerFunc) Handle(params GetFlagSnapshotsParams) middleware.Responder {
	return fn(params)
}

// GetFlagSnapshotsHandler interface for that can handle valid get flag snapshots params
type GetFlagSnapshotsHandler interface {
	Handle(GetFlagSnapshotsParams) middleware.Responder
}

// NewGetFlagSnapshots creates a new http.Handler for the get flag snapshots operation
func NewGetFlagSnapshots(ctx *middleware.Context, handler GetFlagSnapshotsHandler) *GetFlagSnapshots {
	return &GetFlagSnapshots{Context: ctx, Handler: handler}
}

/* GetFlagSnapshots swagger:route GET /flags/{flagID}/snapshots flag getFlagSnapshots

GetFlagSnapshots get flag snapshots API

*/
type GetFlagSnapshots struct {
	Context *middleware.Context
	Handler GetFlagSnapshotsHandler
}

func (o *GetFlagSnapshots) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetFlagSnapshotsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
