package events

// EndpointType stores the details of the endpoint API name and operationID
type EndpointType struct {
	APIName     string
	OperationID string
}

// EndpointKeyType is used to key into the map of EndpointTypes
type EndpointKeyType string

// -------------------- API NAMES --------------------

// primeAPIName is a const string to use the EndpointTypes
const supportAPIName string = "supportapi"

// -------------------- ENDPOINT KEYS --------------------

// SupportListMTOsEndpointKey is the key for the listMTOs endpoint in support
const SupportListMTOsEndpointKey = "Support.ListMTOs"

// SupportCreateMoveTaskOrderEndpointKey is the key for the createMoveTaskOrder endpoint in support
const SupportCreateMoveTaskOrderEndpointKey = "Support.CreateMoveTaskOrder"

// SupportGetMoveTaskOrderEndpointKey is the key for the getMoveTaskOrder endpoint in support
const SupportGetMoveTaskOrderEndpointKey = "Support.GetMoveTaskOrder"

// SupportMakeMoveTaskOrderAvailableEndpointKey is the key for the makeMoveTaskOrderAvailable endpoint in support
const SupportMakeMoveTaskOrderAvailableEndpointKey = "Support.MakeMoveTaskOrderAvailable"

// SupportListMTOPaymentRequestsEndpointKey is the key for the listMTOPaymentRequests endpoint in support
const SupportListMTOPaymentRequestsEndpointKey = "Support.ListMTOPaymentRequests"

// SupportUpdatePaymentRequestStatusEndpointKey is the key for the updatePaymentRequestStatus endpoint in support
const SupportUpdatePaymentRequestStatusEndpointKey = "Support.UpdatePaymentRequestStatus"

// SupportUpdateMTOServiceItemStatusEndpointKey is the key for the updateMTOServiceItemStatus endpoint in support
const SupportUpdateMTOServiceItemStatusEndpointKey = "Support.UpdateMTOServiceItemStatus"

// SupportUpdateMTOShipmentStatusEndpointKey is the key for the updateMTOShipmentStatus endpoint in support
const SupportUpdateMTOShipmentStatusEndpointKey = "Support.UpdateMTOShipmentStatus"

// -------------------- ENDPOINT MAP --------------------
var endpoints map[EndpointKeyType]EndpointType = map[EndpointKeyType]EndpointType{
	SupportListMTOsEndpointKey: {
		APIName:     supportAPIName,
		OperationID: "listMTOs",
	},
	SupportCreateMoveTaskOrderEndpointKey: {
		APIName:     supportAPIName,
		OperationID: "createMoveTaskOrder",
	},
	SupportGetMoveTaskOrderEndpointKey: {
		APIName:     supportAPIName,
		OperationID: "getMoveTaskOrder",
	},
	SupportMakeMoveTaskOrderAvailableEndpointKey: {
		APIName:     supportAPIName,
		OperationID: "makeMoveTaskOrderAvailable",
	},
	SupportListMTOPaymentRequestsEndpointKey: {
		APIName:     supportAPIName,
		OperationID: "listMTOPaymentRequests",
	},
	SupportUpdatePaymentRequestStatusEndpointKey: {
		APIName:     supportAPIName,
		OperationID: "updatePaymentRequestStatus",
	},
	SupportUpdateMTOServiceItemStatusEndpointKey: {
		APIName:     supportAPIName,
		OperationID: "updateMTOServiceItemStatus",
	},
	SupportUpdateMTOShipmentStatusEndpointKey: {
		APIName:     supportAPIName,
		OperationID: "updateMTOShipmentStatus",
	},
}

// String returns the string representation of the endpoint name
func (e EndpointType) String() string {
	return e.APIName + "." + e.OperationID
}

// GetEndpointAPI returns the api name of the endpoint
func GetEndpointAPI(key EndpointKeyType) string {
	return endpoints[key].APIName
}

// GetEndpointOperationID retuns the operation ID of the endpoint
func GetEndpointOperationID(key EndpointKeyType) string {
	return endpoints[key].OperationID
}
