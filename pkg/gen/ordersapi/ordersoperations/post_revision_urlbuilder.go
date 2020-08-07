// Code generated by go-swagger; DO NOT EDIT.

package ordersoperations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

// PostRevisionURL generates an URL for the post revision operation
type PostRevisionURL struct {
	Issuer    string
	MemberID  string
	OrdersNum string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PostRevisionURL) WithBasePath(bp string) *PostRevisionURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PostRevisionURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *PostRevisionURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/orders"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/orders/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	issuerQ := o.Issuer
	if issuerQ != "" {
		qs.Set("issuer", issuerQ)
	}

	memberIDQ := o.MemberID
	if memberIDQ != "" {
		qs.Set("memberId", memberIDQ)
	}

	ordersNumQ := o.OrdersNum
	if ordersNumQ != "" {
		qs.Set("ordersNum", ordersNumQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *PostRevisionURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *PostRevisionURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *PostRevisionURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on PostRevisionURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on PostRevisionURL")
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
func (o *PostRevisionURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
