// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewFetchMTOUpdatesParams creates a new FetchMTOUpdatesParams object
// no default values defined in spec.
func NewFetchMTOUpdatesParams() FetchMTOUpdatesParams {

	return FetchMTOUpdatesParams{}
}

// FetchMTOUpdatesParams contains all the bound params for the fetch m t o updates operation
// typically these are obtained from a http.Request
//
// swagger:parameters fetchMTOUpdates
type FetchMTOUpdatesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Only return move task orders updated since this time.
	  In: query
	*/
	Since *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewFetchMTOUpdatesParams() beforehand.
func (o *FetchMTOUpdatesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qSince, qhkSince, _ := qs.GetOK("since")
	if err := o.bindSince(qSince, qhkSince, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindSince binds and validates parameter Since from query.
func (o *FetchMTOUpdatesParams) bindSince(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("since", "query", "int64", raw)
	}
	o.Since = &value

	return nil
}
