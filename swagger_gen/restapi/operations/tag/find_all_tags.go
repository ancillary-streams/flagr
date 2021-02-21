// Code generated by go-swagger; DO NOT EDIT.

package tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// FindAllTagsHandlerFunc turns a function with the right signature into a find all tags handler
type FindAllTagsHandlerFunc func(FindAllTagsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FindAllTagsHandlerFunc) Handle(params FindAllTagsParams) middleware.Responder {
	return fn(params)
}

// FindAllTagsHandler interface for that can handle valid find all tags params
type FindAllTagsHandler interface {
	Handle(FindAllTagsParams) middleware.Responder
}

// NewFindAllTags creates a new http.Handler for the find all tags operation
func NewFindAllTags(ctx *middleware.Context, handler FindAllTagsHandler) *FindAllTags {
	return &FindAllTags{Context: ctx, Handler: handler}
}

/* FindAllTags swagger:route GET /tags tag findAllTags

FindAllTags find all tags API

*/
type FindAllTags struct {
	Context *middleware.Context
	Handler FindAllTagsHandler
}

func (o *FindAllTags) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFindAllTagsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
