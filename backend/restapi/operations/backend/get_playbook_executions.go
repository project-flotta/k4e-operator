// Code generated by go-swagger; DO NOT EDIT.

package backend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetPlaybookExecutionsHandlerFunc turns a function with the right signature into a get playbook executions handler
type GetPlaybookExecutionsHandlerFunc func(GetPlaybookExecutionsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPlaybookExecutionsHandlerFunc) Handle(params GetPlaybookExecutionsParams) middleware.Responder {
	return fn(params)
}

// GetPlaybookExecutionsHandler interface for that can handle valid get playbook executions params
type GetPlaybookExecutionsHandler interface {
	Handle(GetPlaybookExecutionsParams) middleware.Responder
}

// NewGetPlaybookExecutions creates a new http.Handler for the get playbook executions operation
func NewGetPlaybookExecutions(ctx *middleware.Context, handler GetPlaybookExecutionsHandler) *GetPlaybookExecutions {
	return &GetPlaybookExecutions{Context: ctx, Handler: handler}
}

/*
	GetPlaybookExecutions swagger:route GET /namespaces/{namespace}/devices/{device-id}/playbookexecutions backend getPlaybookExecutions

Returns the playbook executions.
*/
type GetPlaybookExecutions struct {
	Context *middleware.Context
	Handler GetPlaybookExecutionsHandler
}

func (o *GetPlaybookExecutions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetPlaybookExecutionsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
