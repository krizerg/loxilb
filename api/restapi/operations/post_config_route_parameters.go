// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"loxilb/api/models"
)

// NewPostConfigRouteParams creates a new PostConfigRouteParams object
//
// There are no default values defined in the spec.
func NewPostConfigRouteParams() PostConfigRouteParams {

	return PostConfigRouteParams{}
}

// PostConfigRouteParams contains all the bound params for the post config route operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostConfigRoute
type PostConfigRouteParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Attributes for load balance service
	  Required: true
	  In: body
	*/
	Attr *models.RouteEntry
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostConfigRouteParams() beforehand.
func (o *PostConfigRouteParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.RouteEntry
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("attr", "body", ""))
			} else {
				res = append(res, errors.NewParseError("attr", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Attr = &body
			}
		}
	} else {
		res = append(res, errors.Required("attr", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
