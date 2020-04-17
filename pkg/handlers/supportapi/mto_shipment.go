package supportapi

import (
	"fmt"

	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"

	"github.com/transcom/mymove/pkg/gen/supportmessages"
	"github.com/transcom/mymove/pkg/models"
	mtoshipment "github.com/transcom/mymove/pkg/services/mto_shipment"

	"github.com/transcom/mymove/pkg/handlers/supportapi/internal/payloads"

	"github.com/go-openapi/runtime/middleware"

	mtoshipmentops "github.com/transcom/mymove/pkg/gen/supportapi/supportoperations/mto_shipment"
	"github.com/transcom/mymove/pkg/handlers"
	"github.com/transcom/mymove/pkg/services"
)

// Copy of the validators from GHC API with supportmessages instead of ghcmessages
func payloadForClientError(title string, detail string, instance uuid.UUID) *supportmessages.ClientError {
	return &supportmessages.ClientError{
		Title:    handlers.FmtString(title),
		Detail:   handlers.FmtString(detail),
		Instance: handlers.FmtUUID(instance),
	}
}

func payloadForValidationError(title string, detail string, instance uuid.UUID, validationErrors *validate.Errors) *supportmessages.ValidationError {
	return &supportmessages.ValidationError{
		InvalidFields: handlers.NewValidationErrorsResponse(validationErrors).Errors,
		ClientError:   *payloadForClientError(title, detail, instance),
	}
}

// PatchMTOShipmentStatusHandlerFunc updates the status of a MTO Shipment
type PatchMTOShipmentStatusHandlerFunc struct {
	handlers.HandlerContext
	services.Fetcher
	services.MTOShipmentStatusUpdater
}

// Handle updates the status of a MTO Shipment
func (h PatchMTOShipmentStatusHandlerFunc) Handle(params mtoshipmentops.PatchMTOShipmentStatusParams) middleware.Responder {
	_, logger := h.SessionAndLoggerFromRequest(params.HTTPRequest)

	shipmentID := uuid.FromStringOrNil(params.ShipmentID.String())
	status := models.MTOShipmentStatus(params.Body.Status)
	rejectionReason := params.Body.RejectionReason
	eTag := params.IfMatch

	shipment, err := h.UpdateMTOShipmentStatus(shipmentID, status, rejectionReason, eTag)

	if err != nil {
		logger.Error("UpdateMTOShipmentStatus error: ", zap.Error(err))

		switch e := err.(type) {
		case services.NotFoundError:
			notFoundPayload := &supportmessages.Error{Message: handlers.FmtString(fmt.Sprintf("id: '%s' not found", shipmentID.String()))}
			return mtoshipmentops.NewPatchMTOShipmentStatusNotFound().WithPayload(notFoundPayload)
		case services.InvalidInputError:
			payload := payloadForValidationError("Validation errors", "UpdateShipmentMTOStatus", h.GetTraceID(), e.ValidationErrors)
			return mtoshipmentops.NewPatchMTOShipmentStatusUnprocessableEntity().WithPayload(payload)
		case services.PreconditionFailedError:
			return mtoshipmentops.NewPatchMTOShipmentStatusPreconditionFailed().WithPayload(&supportmessages.Error{Message: handlers.FmtString(err.Error())})
		case mtoshipment.ConflictStatusError:
			return mtoshipmentops.NewPatchMTOShipmentStatusConflict().WithPayload(&supportmessages.Error{Message: handlers.FmtString(err.Error())})
		default:
			return mtoshipmentops.NewPatchMTOShipmentStatusInternalServerError()
		}
	}

	payload := payloads.MTOShipment(shipment)
	return mtoshipmentops.NewPatchMTOShipmentStatusOK().WithPayload(payload)
}
