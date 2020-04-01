// Code generated by go-swagger; DO NOT EDIT.

package supportoperations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/transcom/mymove/pkg/gen/supportapi/supportoperations/move_task_order"
	"github.com/transcom/mymove/pkg/gen/supportapi/supportoperations/mto_service_item"
	"github.com/transcom/mymove/pkg/gen/supportapi/supportoperations/mto_shipment"
	"github.com/transcom/mymove/pkg/gen/supportapi/supportoperations/payment_requests"
	"github.com/transcom/mymove/pkg/gen/supportapi/supportoperations/uploads"
)

// NewMymoveAPI creates a new Mymove instance
func NewMymoveAPI(spec *loads.Document) *MymoveAPI {
	return &MymoveAPI{
		handlers:              make(map[string]map[string]http.Handler),
		formats:               strfmt.Default,
		defaultConsumes:       "application/json",
		defaultProduces:       "application/json",
		customConsumers:       make(map[string]runtime.Consumer),
		customProducers:       make(map[string]runtime.Producer),
		ServerShutdown:        func() {},
		spec:                  spec,
		ServeError:            errors.ServeError,
		BasicAuthenticator:    security.BasicAuth,
		APIKeyAuthenticator:   security.APIKeyAuth,
		BearerAuthenticator:   security.BearerAuth,
		JSONConsumer:          runtime.JSONConsumer(),
		MultipartformConsumer: runtime.DiscardConsumer,
		JSONProducer:          runtime.JSONProducer(),
		MtoServiceItemCreateMTOServiceItemHandler: mto_service_item.CreateMTOServiceItemHandlerFunc(func(params mto_service_item.CreateMTOServiceItemParams) middleware.Responder {
			return middleware.NotImplemented("operation MtoServiceItemCreateMTOServiceItem has not yet been implemented")
		}),
		PaymentRequestsCreatePaymentRequestHandler: payment_requests.CreatePaymentRequestHandlerFunc(func(params payment_requests.CreatePaymentRequestParams) middleware.Responder {
			return middleware.NotImplemented("operation PaymentRequestsCreatePaymentRequest has not yet been implemented")
		}),
		UploadsCreateUploadHandler: uploads.CreateUploadHandlerFunc(func(params uploads.CreateUploadParams) middleware.Responder {
			return middleware.NotImplemented("operation UploadsCreateUpload has not yet been implemented")
		}),
		MoveTaskOrderFetchMTOUpdatesHandler: move_task_order.FetchMTOUpdatesHandlerFunc(func(params move_task_order.FetchMTOUpdatesParams) middleware.Responder {
			return middleware.NotImplemented("operation MoveTaskOrderFetchMTOUpdates has not yet been implemented")
		}),
		MoveTaskOrderGetMoveTaskOrderCustomerHandler: move_task_order.GetMoveTaskOrderCustomerHandlerFunc(func(params move_task_order.GetMoveTaskOrderCustomerParams) middleware.Responder {
			return middleware.NotImplemented("operation MoveTaskOrderGetMoveTaskOrderCustomer has not yet been implemented")
		}),
		MoveTaskOrderUpdateMTOPostCounselingInformationHandler: move_task_order.UpdateMTOPostCounselingInformationHandlerFunc(func(params move_task_order.UpdateMTOPostCounselingInformationParams) middleware.Responder {
			return middleware.NotImplemented("operation MoveTaskOrderUpdateMTOPostCounselingInformation has not yet been implemented")
		}),
		MtoShipmentUpdateMTOShipmentHandler: mto_shipment.UpdateMTOShipmentHandlerFunc(func(params mto_shipment.UpdateMTOShipmentParams) middleware.Responder {
			return middleware.NotImplemented("operation MtoShipmentUpdateMTOShipment has not yet been implemented")
		}),
	}
}

/*MymoveAPI The Support API for move.mil */
type MymoveAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for a "multipart/form-data" mime type
	MultipartformConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// MtoServiceItemCreateMTOServiceItemHandler sets the operation handler for the create m t o service item operation
	MtoServiceItemCreateMTOServiceItemHandler mto_service_item.CreateMTOServiceItemHandler
	// PaymentRequestsCreatePaymentRequestHandler sets the operation handler for the create payment request operation
	PaymentRequestsCreatePaymentRequestHandler payment_requests.CreatePaymentRequestHandler
	// UploadsCreateUploadHandler sets the operation handler for the create upload operation
	UploadsCreateUploadHandler uploads.CreateUploadHandler
	// MoveTaskOrderFetchMTOUpdatesHandler sets the operation handler for the fetch m t o updates operation
	MoveTaskOrderFetchMTOUpdatesHandler move_task_order.FetchMTOUpdatesHandler
	// MoveTaskOrderGetMoveTaskOrderCustomerHandler sets the operation handler for the get move task order customer operation
	MoveTaskOrderGetMoveTaskOrderCustomerHandler move_task_order.GetMoveTaskOrderCustomerHandler
	// MoveTaskOrderUpdateMTOPostCounselingInformationHandler sets the operation handler for the update m t o post counseling information operation
	MoveTaskOrderUpdateMTOPostCounselingInformationHandler move_task_order.UpdateMTOPostCounselingInformationHandler
	// MtoShipmentUpdateMTOShipmentHandler sets the operation handler for the update m t o shipment operation
	MtoShipmentUpdateMTOShipmentHandler mto_shipment.UpdateMTOShipmentHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *MymoveAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *MymoveAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *MymoveAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *MymoveAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *MymoveAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *MymoveAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *MymoveAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the MymoveAPI
func (o *MymoveAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.MtoServiceItemCreateMTOServiceItemHandler == nil {
		unregistered = append(unregistered, "mto_service_item.CreateMTOServiceItemHandler")
	}

	if o.PaymentRequestsCreatePaymentRequestHandler == nil {
		unregistered = append(unregistered, "payment_requests.CreatePaymentRequestHandler")
	}

	if o.UploadsCreateUploadHandler == nil {
		unregistered = append(unregistered, "uploads.CreateUploadHandler")
	}

	if o.MoveTaskOrderFetchMTOUpdatesHandler == nil {
		unregistered = append(unregistered, "move_task_order.FetchMTOUpdatesHandler")
	}

	if o.MoveTaskOrderGetMoveTaskOrderCustomerHandler == nil {
		unregistered = append(unregistered, "move_task_order.GetMoveTaskOrderCustomerHandler")
	}

	if o.MoveTaskOrderUpdateMTOPostCounselingInformationHandler == nil {
		unregistered = append(unregistered, "move_task_order.UpdateMTOPostCounselingInformationHandler")
	}

	if o.MtoShipmentUpdateMTOShipmentHandler == nil {
		unregistered = append(unregistered, "mto_shipment.UpdateMTOShipmentHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *MymoveAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *MymoveAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *MymoveAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *MymoveAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *MymoveAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *MymoveAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the mymove API
func (o *MymoveAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *MymoveAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/move-task-orders/{moveTaskOrderID}/mto-shipments/{mtoShipmentID}/mto-service-items"] = mto_service_item.NewCreateMTOServiceItem(o.context, o.MtoServiceItemCreateMTOServiceItemHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/payment-requests"] = payment_requests.NewCreatePaymentRequest(o.context, o.PaymentRequestsCreatePaymentRequestHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/payment-requests/{paymentRequestID}/uploads"] = uploads.NewCreateUpload(o.context, o.UploadsCreateUploadHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/move-task-orders"] = move_task_order.NewFetchMTOUpdates(o.context, o.MoveTaskOrderFetchMTOUpdatesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/move-task-orders/{moveTaskOrderID}/customer"] = move_task_order.NewGetMoveTaskOrderCustomer(o.context, o.MoveTaskOrderGetMoveTaskOrderCustomerHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/move-task-orders/{moveTaskOrderID}/post-counseling-info"] = move_task_order.NewUpdateMTOPostCounselingInformation(o.context, o.MoveTaskOrderUpdateMTOPostCounselingInformationHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/move-task-orders/{moveTaskOrderID}/mto-shipments/{mtoShipmentID}"] = mto_shipment.NewUpdateMTOShipment(o.context, o.MtoShipmentUpdateMTOShipmentHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *MymoveAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *MymoveAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *MymoveAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *MymoveAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
