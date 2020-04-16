package payloads

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/transcom/mymove/pkg/etag"
	"github.com/transcom/mymove/pkg/gen/supportmessages"
	"github.com/transcom/mymove/pkg/handlers"
	"github.com/transcom/mymove/pkg/models"
)

// MoveTaskOrder payload
func MoveTaskOrder(moveTaskOrder *models.MoveTaskOrder) *supportmessages.MoveTaskOrder {
	if moveTaskOrder == nil {
		return nil
	}

	payload := &supportmessages.MoveTaskOrder{
		ID:                 strfmt.UUID(moveTaskOrder.ID.String()),
		CreatedAt:          strfmt.Date(moveTaskOrder.CreatedAt),
		IsAvailableToPrime: &moveTaskOrder.IsAvailableToPrime,
		IsCanceled:         &moveTaskOrder.IsCanceled,
		MoveOrderID:        strfmt.UUID(moveTaskOrder.MoveOrderID.String()),
		ReferenceID:        moveTaskOrder.ReferenceID,
		UpdatedAt:          strfmt.Date(moveTaskOrder.UpdatedAt),
		ETag:               etag.GenerateEtag(moveTaskOrder.UpdatedAt),
	}
	return payload
}

// Customer payload
func Customer(customer *models.Customer) *supportmessages.Customer {
	if customer == nil {
		return nil
	}
	payload := supportmessages.Customer{
		Agency:             swag.StringValue(customer.Agency),
		CurrentAddress:     Address(&customer.CurrentAddress),
		DestinationAddress: Address(&customer.DestinationAddress),
		DodID:              swag.StringValue(customer.DODID),
		Email:              customer.Email,
		FirstName:          swag.StringValue(customer.FirstName),
		ID:                 strfmt.UUID(customer.ID.String()),
		LastName:           swag.StringValue(customer.LastName),
		Phone:              customer.PhoneNumber,
		UserID:             strfmt.UUID(customer.UserID.String()),
		ETag:               etag.GenerateEtag(customer.UpdatedAt),
	}
	return &payload
}

// MoveOrder payload
func MoveOrder(moveOrder *models.MoveOrder) *supportmessages.MoveOrder {
	if moveOrder == nil {
		return nil
	}
	destinationDutyStation := DutyStation(moveOrder.DestinationDutyStation)
	originDutyStation := DutyStation(moveOrder.OriginDutyStation)
	if moveOrder.Grade != nil {
		moveOrder.Entitlement.SetWeightAllotment(*moveOrder.Grade)
	}
	entitlements := Entitlement(moveOrder.Entitlement)
	payload := supportmessages.MoveOrder{
		DestinationDutyStation: destinationDutyStation,
		Entitlement:            entitlements,
		OrderNumber:            moveOrder.OrderNumber,
		OrderTypeDetail:        moveOrder.OrderTypeDetail,
		ID:                     strfmt.UUID(moveOrder.ID.String()),
		OriginDutyStation:      originDutyStation,
		ETag:                   etag.GenerateEtag(moveOrder.UpdatedAt),
	}

	if moveOrder.Customer != nil {
		payload.Agency = swag.StringValue(moveOrder.Customer.Agency)
		payload.CustomerID = strfmt.UUID(moveOrder.CustomerID.String())
		payload.FirstName = swag.StringValue(moveOrder.Customer.FirstName)
		payload.LastName = swag.StringValue(moveOrder.Customer.LastName)
	}
	if moveOrder.ReportByDate != nil {
		payload.ReportByDate = strfmt.Date(*moveOrder.ReportByDate)
	}
	if moveOrder.DateIssued != nil {
		payload.DateIssued = strfmt.Date(*moveOrder.DateIssued)
	}
	if moveOrder.Grade != nil {
		payload.Grade = *moveOrder.Grade
	}
	if moveOrder.ConfirmationNumber != nil {
		payload.ConfirmationNumber = *moveOrder.ConfirmationNumber
	}
	if moveOrder.OrderType != nil {
		payload.OrderType = *moveOrder.OrderType
	}

	return &payload
}

// Entitlement payload
func Entitlement(entitlement *models.Entitlement) *supportmessages.Entitlements {
	if entitlement == nil {
		return nil
	}
	var proGearWeight, proGearWeightSpouse, totalWeight int64
	if entitlement.WeightAllotment() != nil {
		proGearWeight = int64(entitlement.WeightAllotment().ProGearWeight)
		proGearWeightSpouse = int64(entitlement.WeightAllotment().ProGearWeightSpouse)
		totalWeight = int64(entitlement.WeightAllotment().TotalWeightSelf)
	}
	var authorizedWeight *int64
	if entitlement.AuthorizedWeight() != nil {
		aw := int64(*entitlement.AuthorizedWeight())
		authorizedWeight = &aw
	}
	var sit int64
	if entitlement.StorageInTransit != nil {
		sit = int64(*entitlement.StorageInTransit)
	}
	var totalDependents int64
	if entitlement.TotalDependents != nil {
		totalDependents = int64(*entitlement.TotalDependents)
	}
	return &supportmessages.Entitlements{
		ID:                    strfmt.UUID(entitlement.ID.String()),
		AuthorizedWeight:      authorizedWeight,
		DependentsAuthorized:  entitlement.DependentsAuthorized,
		NonTemporaryStorage:   entitlement.NonTemporaryStorage,
		PrivatelyOwnedVehicle: entitlement.PrivatelyOwnedVehicle,
		ProGearWeight:         proGearWeight,
		ProGearWeightSpouse:   proGearWeightSpouse,
		StorageInTransit:      sit,
		TotalDependents:       totalDependents,
		TotalWeight:           totalWeight,
		ETag:                  etag.GenerateEtag(entitlement.UpdatedAt),
	}
}

// DutyStation payload
func DutyStation(dutyStation *models.DutyStation) *supportmessages.DutyStation {
	if dutyStation == nil {
		return nil
	}
	address := Address(&dutyStation.Address)
	payload := supportmessages.DutyStation{
		Address:   address,
		AddressID: address.ID,
		ID:        strfmt.UUID(dutyStation.ID.String()),
		Name:      dutyStation.Name,
		ETag:      etag.GenerateEtag(dutyStation.UpdatedAt),
	}
	return &payload
}

// Address payload
func Address(address *models.Address) *supportmessages.Address {
	if address == nil {
		return nil
	}
	return &supportmessages.Address{
		ID:             strfmt.UUID(address.ID.String()),
		StreetAddress1: &address.StreetAddress1,
		StreetAddress2: address.StreetAddress2,
		StreetAddress3: address.StreetAddress3,
		City:           &address.City,
		State:          &address.State,
		PostalCode:     &address.PostalCode,
		Country:        address.Country,
		ETag:           etag.GenerateEtag(address.UpdatedAt),
	}
}

// MTOShipment payload
func MTOShipment(mtoShipment *models.MTOShipment) *supportmessages.MTOShipment {
	strfmt.MarshalFormat = strfmt.RFC3339Micro

	payload := &supportmessages.MTOShipment{
		ID:                       strfmt.UUID(mtoShipment.ID.String()),
		MoveTaskOrderID:          strfmt.UUID(mtoShipment.MoveTaskOrderID.String()),
		ShipmentType:             mtoShipment.ShipmentType,
		Status:                   string(mtoShipment.Status),
		CustomerRemarks:          mtoShipment.CustomerRemarks,
		RejectionReason:          mtoShipment.RejectionReason,
		PickupAddress:            Address(mtoShipment.PickupAddress),
		SecondaryDeliveryAddress: Address(mtoShipment.SecondaryDeliveryAddress),
		SecondaryPickupAddress:   Address(mtoShipment.SecondaryPickupAddress),
		DestinationAddress:       Address(mtoShipment.DestinationAddress),
		CreatedAt:                strfmt.DateTime(mtoShipment.CreatedAt),
		UpdatedAt:                strfmt.DateTime(mtoShipment.UpdatedAt),
		ETag:                     etag.GenerateEtag(mtoShipment.UpdatedAt),
	}

	if mtoShipment.RequestedPickupDate != nil {
		payload.RequestedPickupDate = *handlers.FmtDatePtr(mtoShipment.RequestedPickupDate)
	}

	if mtoShipment.ApprovedDate != nil {
		payload.ApprovedDate = strfmt.Date(*mtoShipment.ApprovedDate)
	}

	return payload
}

// MTOShipments payload
func MTOShipments(mtoShipments *models.MTOShipments) *supportmessages.MTOShipments {
	payload := make(supportmessages.MTOShipments, len(*mtoShipments))

	for i, m := range *mtoShipments {
		payload[i] = MTOShipment(&m)
	}
	return &payload
}

// MTOAgent payload
func MTOAgent(mtoAgent *models.MTOAgent) *supportmessages.MTOAgent {
	payload := &supportmessages.MTOAgent{
		ID:            strfmt.UUID(mtoAgent.ID.String()),
		MtoShipmentID: strfmt.UUID(mtoAgent.MTOShipmentID.String()),
		CreatedAt:     strfmt.Date(mtoAgent.CreatedAt),
		UpdatedAt:     strfmt.Date(mtoAgent.UpdatedAt),
		FirstName:     mtoAgent.FirstName,
		LastName:      mtoAgent.LastName,
		AgentType:     string(mtoAgent.MTOAgentType),
		Email:         mtoAgent.Email,
		Phone:         mtoAgent.Phone,
		ETag:          etag.GenerateEtag(mtoAgent.UpdatedAt),
	}
	return payload
}

// MTOAgents payload
func MTOAgents(mtoAgents *models.MTOAgents) *supportmessages.MTOAgents {
	payload := make(supportmessages.MTOAgents, len(*mtoAgents))
	for i, m := range *mtoAgents {
		payload[i] = MTOAgent(&m)
	}
	return &payload
}

// PaymentRequest payload
func PaymentRequest(pr *models.PaymentRequest) *supportmessages.PaymentRequest {
	return &supportmessages.PaymentRequest{
		ID:                   *handlers.FmtUUID(pr.ID),
		IsFinal:              &pr.IsFinal,
		MoveTaskOrderID:      *handlers.FmtUUID(pr.MoveTaskOrderID),
		PaymentRequestNumber: pr.PaymentRequestNumber,
		RejectionReason:      pr.RejectionReason,
		Status:               supportmessages.PaymentRequestStatus(pr.Status),
		ETag:                 etag.GenerateEtag(pr.UpdatedAt),
	}
}