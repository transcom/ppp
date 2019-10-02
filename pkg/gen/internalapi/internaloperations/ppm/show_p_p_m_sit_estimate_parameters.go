// Code generated by go-swagger; DO NOT EDIT.

package ppm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewShowPPMSitEstimateParams creates a new ShowPPMSitEstimateParams object
// no default values defined in spec.
func NewShowPPMSitEstimateParams() ShowPPMSitEstimateParams {

	return ShowPPMSitEstimateParams{}
}

// ShowPPMSitEstimateParams contains all the bound params for the show p p m sit estimate operation
// typically these are obtained from a http.Request
//
// swagger:parameters showPPMSitEstimate
type ShowPPMSitEstimateParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: query
	*/
	DaysInStorage int64
	/*
	  Required: true
	  Pattern: ^(\d{5}([\-]\d{4})?)$
	  In: query
	*/
	DestinationZip string
	/*
	  Required: true
	  Pattern: ^(\d{5}([\-]\d{4})?)$
	  In: query
	*/
	OriginZip string
	/*
	  Required: true
	  In: query
	*/
	OriginalMoveDate strfmt.Date
	/*
	  Required: true
	  In: query
	*/
	WeightEstimate int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewShowPPMSitEstimateParams() beforehand.
func (o *ShowPPMSitEstimateParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qDaysInStorage, qhkDaysInStorage, _ := qs.GetOK("days_in_storage")
	if err := o.bindDaysInStorage(qDaysInStorage, qhkDaysInStorage, route.Formats); err != nil {
		res = append(res, err)
	}

	qDestinationZip, qhkDestinationZip, _ := qs.GetOK("destination_zip")
	if err := o.bindDestinationZip(qDestinationZip, qhkDestinationZip, route.Formats); err != nil {
		res = append(res, err)
	}

	qOriginZip, qhkOriginZip, _ := qs.GetOK("origin_zip")
	if err := o.bindOriginZip(qOriginZip, qhkOriginZip, route.Formats); err != nil {
		res = append(res, err)
	}

	qOriginalMoveDate, qhkOriginalMoveDate, _ := qs.GetOK("original_move_date")
	if err := o.bindOriginalMoveDate(qOriginalMoveDate, qhkOriginalMoveDate, route.Formats); err != nil {
		res = append(res, err)
	}

	qWeightEstimate, qhkWeightEstimate, _ := qs.GetOK("weight_estimate")
	if err := o.bindWeightEstimate(qWeightEstimate, qhkWeightEstimate, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDaysInStorage binds and validates parameter DaysInStorage from query.
func (o *ShowPPMSitEstimateParams) bindDaysInStorage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("days_in_storage", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("days_in_storage", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("days_in_storage", "query", "int64", raw)
	}
	o.DaysInStorage = value

	return nil
}

// bindDestinationZip binds and validates parameter DestinationZip from query.
func (o *ShowPPMSitEstimateParams) bindDestinationZip(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("destination_zip", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("destination_zip", "query", raw); err != nil {
		return err
	}

	o.DestinationZip = raw

	if err := o.validateDestinationZip(formats); err != nil {
		return err
	}

	return nil
}

// validateDestinationZip carries on validations for parameter DestinationZip
func (o *ShowPPMSitEstimateParams) validateDestinationZip(formats strfmt.Registry) error {

	if err := validate.Pattern("destination_zip", "query", o.DestinationZip, `^(\d{5}([\-]\d{4})?)$`); err != nil {
		return err
	}

	return nil
}

// bindOriginZip binds and validates parameter OriginZip from query.
func (o *ShowPPMSitEstimateParams) bindOriginZip(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("origin_zip", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("origin_zip", "query", raw); err != nil {
		return err
	}

	o.OriginZip = raw

	if err := o.validateOriginZip(formats); err != nil {
		return err
	}

	return nil
}

// validateOriginZip carries on validations for parameter OriginZip
func (o *ShowPPMSitEstimateParams) validateOriginZip(formats strfmt.Registry) error {

	if err := validate.Pattern("origin_zip", "query", o.OriginZip, `^(\d{5}([\-]\d{4})?)$`); err != nil {
		return err
	}

	return nil
}

// bindOriginalMoveDate binds and validates parameter OriginalMoveDate from query.
func (o *ShowPPMSitEstimateParams) bindOriginalMoveDate(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("original_move_date", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("original_move_date", "query", raw); err != nil {
		return err
	}

	// Format: date
	value, err := formats.Parse("date", raw)
	if err != nil {
		return errors.InvalidType("original_move_date", "query", "strfmt.Date", raw)
	}
	o.OriginalMoveDate = *(value.(*strfmt.Date))

	if err := o.validateOriginalMoveDate(formats); err != nil {
		return err
	}

	return nil
}

// validateOriginalMoveDate carries on validations for parameter OriginalMoveDate
func (o *ShowPPMSitEstimateParams) validateOriginalMoveDate(formats strfmt.Registry) error {

	if err := validate.FormatOf("original_move_date", "query", "date", o.OriginalMoveDate.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindWeightEstimate binds and validates parameter WeightEstimate from query.
func (o *ShowPPMSitEstimateParams) bindWeightEstimate(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("weight_estimate", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("weight_estimate", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("weight_estimate", "query", "int64", raw)
	}
	o.WeightEstimate = value

	return nil
}
