// Code generated by go-swagger; DO NOT EDIT.

package queues

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// GetMovesQueueURL generates an URL for the get moves queue operation
type GetMovesQueueURL struct {
	Branch                 *string
	DestinationDutyStation *string
	DodID                  *string
	LastName               *string
	MoveID                 *string
	Status                 []string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetMovesQueueURL) WithBasePath(bp string) *GetMovesQueueURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetMovesQueueURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetMovesQueueURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/queues/moves"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/ghc/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var branchQ string
	if o.Branch != nil {
		branchQ = *o.Branch
	}
	if branchQ != "" {
		qs.Set("branch", branchQ)
	}

	var destinationDutyStationQ string
	if o.DestinationDutyStation != nil {
		destinationDutyStationQ = *o.DestinationDutyStation
	}
	if destinationDutyStationQ != "" {
		qs.Set("destinationDutyStation", destinationDutyStationQ)
	}

	var dodIDQ string
	if o.DodID != nil {
		dodIDQ = *o.DodID
	}
	if dodIDQ != "" {
		qs.Set("dodID", dodIDQ)
	}

	var lastNameQ string
	if o.LastName != nil {
		lastNameQ = *o.LastName
	}
	if lastNameQ != "" {
		qs.Set("lastName", lastNameQ)
	}

	var moveIDQ string
	if o.MoveID != nil {
		moveIDQ = *o.MoveID
	}
	if moveIDQ != "" {
		qs.Set("moveID", moveIDQ)
	}

	var statusIR []string
	for _, statusI := range o.Status {
		statusIS := statusI
		if statusIS != "" {
			statusIR = append(statusIR, statusIS)
		}
	}

	status := swag.JoinByFormat(statusIR, "")

	if len(status) > 0 {
		qsv := status[0]
		if qsv != "" {
			qs.Set("status", qsv)
		}
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetMovesQueueURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetMovesQueueURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetMovesQueueURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetMovesQueueURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetMovesQueueURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetMovesQueueURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
