package invoice

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/transcom/mymove/pkg/models"

	ediinvoice "github.com/transcom/mymove/pkg/edi/invoice"
	edisegment "github.com/transcom/mymove/pkg/edi/segment"
)

// GHCPaymentRequestInvoiceGenerator is a service object to turn payment requests into 858s
type GHCPaymentRequestInvoiceGenerator struct {
	db *pop.Connection
}

const dateFormat = "060102"
const timeFormat = "1504"

// Generate method takes a payment request and returns an Invoice858C
func (g GHCPaymentRequestInvoiceGenerator) Generate(paymentRequest models.PaymentRequest, sendProductionInvoice bool) (ediinvoice.Invoice858C, error) {
	// TODO: probably need to check if the MTO is loaded on the paymentRequest that is passed in, not sure what is more in line with go standards to error out if it's not there or look it up.
	// TODO: seems ReferenceID is a *string but cannot be saved as nil, do we need to validate it's not nil here

	currentTime := time.Now()

	// TODO: interchangeControlNumber, err := icnSequencer.NextVal()
	// if err != nil {
	// 	return Invoice858C{}, errors.Wrap(err, fmt.Sprintf("Failed to get next Interchange Control Number"))
	// }
	var usageIndicator string
	if sendProductionInvoice {
		usageIndicator = "P"
	} else {
		usageIndicator = "T"
	}

	var edi858 ediinvoice.Invoice858C
	isa := edisegment.ISA{
		AuthorizationInformationQualifier: "00", // No authorization information
		AuthorizationInformation:          "0084182369",
		SecurityInformationQualifier:      "00", // No security information
		SecurityInformation:               "_   _",
		InterchangeSenderIDQualifier:      "ZZ",
		InterchangeSenderID:               "GOVDPIBS",
		InterchangeReceiverIDQualifier:    "12",
		InterchangeReceiverID:             "8004171844",
		InterchangeDate:                   currentTime.Format(dateFormat),
		InterchangeTime:                   currentTime.Format(timeFormat),
		InterchangeControlStandards:       "U",
		InterchangeControlVersionNumber:   "00401",
		InterchangeControlNumber:          100001272,
		AcknowledgementRequested:          0,
		UsageIndicator:                    usageIndicator, // T for test, P for production
		ComponentElementSeparator:         "|",
	}
	edi858.ISA = append(edi858.ISA, &isa)

	bx := edisegment.BX{
		TransactionSetPurposeCode:    "00",
		TransactionMethodTypeCode:    "J",
		ShipmentMethodOfPayment:      "PP",
		ShipmentIdentificationNumber: *paymentRequest.MoveTaskOrder.ReferenceID,
		StandardCarrierAlphaCode:     "TRUS",
		ShipmentQualifier:            "4",
	}
	edi858.Header = append(edi858.Header, &bx)

	paymentRequestNumberSegment := edisegment.N9{
		ReferenceIdentificationQualifier: "CN",
		ReferenceIdentification:          paymentRequest.PaymentRequestNumber,
	}
	edi858.Header = append(edi858.Header, &paymentRequestNumberSegment)

	var paymentServiceItems models.PaymentServiceItems
	error := g.db.Q().
		Eager("MTOServiceItem.ReService").
		Where("payment_request_id", paymentRequest.ID).
		All(&paymentServiceItems)
	if error != nil {
		return edi858, fmt.Errorf("Could not find payment service items: %w", error)
	}

	paymentServiceItemSegments, err := generatePaymentServiceItemSegments(paymentServiceItems)
	if err != nil {
		return edi858, fmt.Errorf("Could not generate payment service item segments: %w", err)
	}
	edi858.ServiceItems = append(edi858.ServiceItems, &paymentServiceItemSegments)
	}

	return edi858, nil
}

func generatePaymentServiceItemSegments(paymentRequest) ([]edisegment.Segment, error) {
	//Initialize empty collection of segments
	var segments []edisegment.Segment

	// Iterate over payment service items
	for idx, serviceItem := range paymentServiceItems {
		// Initialize empty edisegment
		var tariffSegments []edisegment.Segment

		hierarchicalIDNumber := strconv.Itoa(idx + 1)
		// Build and put together the segments
		hlSegment := edisegment.HL{
			HierarchicalIDNumber:  hierarchicalIDNumber, // may need to change if sending multiple payment request in a single edi
			HierarchicalLevelCode: "|",
		}

		n9Segment := edisegment.N9{
			ReferenceIdentificationQualifier: "PO",
			// pending creation of shorter identifier for payment service item
			// https://dp3.atlassian.net/browse/MB-3718
			ReferenceIdentification: serviceItem.ID.String(),
		}

		// TODO: add another n9 for SIT
		// TODO: add a L5 segment/definition

		var weight models.PaymentServiceItemParam
		// TODO: update to have a case statement as different service items may or may not have weight
		// and the distance key can differ (zip3 v zip5, and distances for SIT)
		error := g.db.Q().
			Join("service_item_param_key sk", "payment_service_item_params.service_item_param_key_id = sk.id").
			Where("payment_service_item_id", serviceItem.ID).
			Where("sk.key = ?", ServiceItemParamNameWeightBilledActual)
		First(&weight)
		var distance models.PaymentServiceItemParam
		error := g.db.Q().
			Join("service_item_param_key sk", "payment_service_item_params.service_item_param_key_id = sk.id").
			Where("payment_service_item_id", serviceItem.ID).
			Where("sk.key = ?", ServiceItemParamNameDistanceZip3)
		First(&distance)
		if error != nil {
			return edi858, fmt.Errorf("Could not lookup PaymentServiceItemParam: %w", error)
		}
		l0Segment := edisegment.L0{
			LadingLineItemNumber:   hierarchicalIDNumber,
			BilledRatedAsQuantity:  distance.Value,
			BilledRatedAsQualifier: "DM",
			Weight:                 weight.Value,
			WeightQualifier:        "B",
			WeightUnitCode:         "L",
		}

		tariffSegments = append(tariffSegments, hlSegment, n9Segment, l0Segment)

		segments = append(segments, tariffSegments...)

		return segments, nil;
}
