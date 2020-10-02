package invoice

import (
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/suite"

	edisegment "github.com/transcom/mymove/pkg/edi/segment"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/testdatagen"
	"github.com/transcom/mymove/pkg/testingsuite"

	"go.uber.org/zap"
)

type GHCInvoiceSuite struct {
	testingsuite.PopTestSuite
	logger Logger
}

func (suite *GHCInvoiceSuite) SetupTest() {
	errTruncateAll := suite.DB().TruncateAll()
	if errTruncateAll != nil {
		log.Panicf("failed to truncate database: %#v", errTruncateAll)
	}
}

func TestGHCInvoiceSuite(t *testing.T) {
	ts := &GHCInvoiceSuite{
		PopTestSuite: testingsuite.NewPopTestSuite(testingsuite.CurrentPackage().Suffix("ghcinvoice")),
		logger:       zap.NewNop(), // Use a no-op logger during testing
	}
	suite.Run(t, ts)
	ts.PopTestSuite.TearDown()
}

const testDateFormat = "20060102"
const testTimeFormat = "1504"

func (suite *GHCInvoiceSuite) TestAllGenerateEdi() {
	currentTime := time.Now()
	generator := NewGHCPaymentRequestInvoiceGenerator(suite.DB())
	basicPaymentServiceItemParams := []testdatagen.CreatePaymentServiceItemParams{
		{
			Key:     models.ServiceItemParamNameContractCode,
			KeyType: models.ServiceItemParamTypeString,
			Value:   testdatagen.DefaultContractCode,
		},
		{
			Key:     models.ServiceItemParamNameRequestedPickupDate,
			KeyType: models.ServiceItemParamTypeDate,
			Value:   currentTime.Format(dateFormat),
		},
		{
			Key:     models.ServiceItemParamNameWeightBilledActual,
			KeyType: models.ServiceItemParamTypeInteger,
			Value:   "4242",
		},
		{
			Key:     models.ServiceItemParamNameDistanceZip3,
			KeyType: models.ServiceItemParamTypeInteger,
			Value:   "2424",
		},
		{
			Key:     models.ServiceItemParamNameDistanceZip5,
			KeyType: models.ServiceItemParamTypeInteger,
			Value:   "24245",
		},
	}

	mto := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{})
	paymentRequest := testdatagen.MakePaymentRequest(suite.DB(), testdatagen.Assertions{
		Move: mto,
		PaymentRequest: models.PaymentRequest{
			IsFinal:         false,
			Status:          models.PaymentRequestStatusPending,
			RejectionReason: nil,
		},
	})

	requestedPickupDate := time.Date(testdatagen.GHCTestYear, time.September, 15, 0, 0, 0, 0, time.UTC)
	scheduledPickupDate := time.Date(testdatagen.GHCTestYear, time.September, 20, 0, 0, 0, 0, time.UTC)
	actualPickupDate := time.Date(testdatagen.GHCTestYear, time.September, 22, 0, 0, 0, 0, time.UTC)

	mtoShipment := testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move: mto,
		MTOShipment: models.MTOShipment{
			RequestedPickupDate: &requestedPickupDate,
			ScheduledPickupDate: &scheduledPickupDate,
			ActualPickupDate:    &actualPickupDate,
		},
	})

	assertions := testdatagen.Assertions{
		Move:           mto,
		MTOShipment:    mtoShipment,
		PaymentRequest: paymentRequest,
	}

	var paymentServiceItems models.PaymentServiceItems
	dlh := testdatagen.MakePaymentServiceItemWithParams(
		suite.DB(),
		models.ReServiceCodeDLH,
		basicPaymentServiceItemParams,
		assertions,
	)
	fsc := testdatagen.MakePaymentServiceItemWithParams(
		suite.DB(),
		models.ReServiceCodeFSC,
		basicPaymentServiceItemParams,
		assertions,
	)
	ms := testdatagen.MakePaymentServiceItemWithParams(
		suite.DB(),
		models.ReServiceCodeMS,
		basicPaymentServiceItemParams,
		assertions,
	)
	cs := testdatagen.MakePaymentServiceItemWithParams(
		suite.DB(),
		models.ReServiceCodeCS,
		basicPaymentServiceItemParams,
		assertions,
	)
	dsh := testdatagen.MakePaymentServiceItemWithParams(
		suite.DB(),
		models.ReServiceCodeDSH,
		basicPaymentServiceItemParams,
		assertions,
	)
	dop := testdatagen.MakePaymentServiceItemWithParams(
		suite.DB(),
		models.ReServiceCodeDOP,
		basicPaymentServiceItemParams,
		assertions,
	)
	ddp := testdatagen.MakePaymentServiceItemWithParams(
		suite.DB(),
		models.ReServiceCodeDDP,
		basicPaymentServiceItemParams,
		assertions,
	)
	dpk := testdatagen.MakePaymentServiceItemWithParams(
		suite.DB(),
		models.ReServiceCodeDPK,
		basicPaymentServiceItemParams,
		assertions,
	)
	dupk := testdatagen.MakePaymentServiceItemWithParams(
		suite.DB(),
		models.ReServiceCodeDUPK,
		basicPaymentServiceItemParams,
		assertions,
	)

	paymentServiceItems = append(paymentServiceItems, dlh, fsc, ms, cs, dsh, dop, ddp, dpk, dupk)

	serviceMember := testdatagen.MakeExtendedServiceMember(suite.DB(), testdatagen.Assertions{
		ServiceMember: models.ServiceMember{
			ID: uuid.FromStringOrNil("d66d2f35-218c-4b85-b9d1-631949b9d984"),
		},
	})

	// Proceed with full EDI Generation tests
	result, err := generator.Generate(paymentRequest, false)
	suite.NoError(err)

	// Test Invoice Start and End Segments
	suite.T().Run("adds isa start segment", func(t *testing.T) {

		suite.Equal("00", result.ISA.AuthorizationInformationQualifier)
		suite.Equal("0084182369", result.ISA.AuthorizationInformation)
		suite.Equal("00", result.ISA.SecurityInformationQualifier)
		suite.Equal("_   _", result.ISA.SecurityInformation)
		suite.Equal("ZZ", result.ISA.InterchangeSenderIDQualifier)
		suite.Equal("GOVDPIBS", result.ISA.InterchangeSenderID)
		suite.Equal("12", result.ISA.InterchangeReceiverIDQualifier)
		suite.Equal("8004171844", result.ISA.InterchangeReceiverID)
		suite.Equal(currentTime.Format(testDateFormat), result.ISA.InterchangeDate)
		suite.Equal(currentTime.Format(testTimeFormat), result.ISA.InterchangeTime)
		suite.Equal("U", result.ISA.InterchangeControlStandards)
		suite.Equal("00401", result.ISA.InterchangeControlVersionNumber)
		suite.Equal(int64(100001272), result.ISA.InterchangeControlNumber)
		suite.Equal(0, result.ISA.AcknowledgementRequested)
		suite.Equal("T", result.ISA.UsageIndicator)
		suite.Equal("|", result.ISA.ComponentElementSeparator)
	})

	suite.T().Run("adds gs start segment", func(t *testing.T) {
		suite.Equal("SI", result.GS.FunctionalIdentifierCode)
		suite.Equal("MYMOVE", result.GS.ApplicationSendersCode)
		suite.Equal("8004171844", result.GS.ApplicationReceiversCode)
		suite.Equal(currentTime.Format(dateFormat), result.GS.Date)
		suite.Equal(currentTime.Format(timeFormat), result.GS.Time)
		suite.Equal(int64(100001251), result.GS.GroupControlNumber)
		suite.Equal("X", result.GS.ResponsibleAgencyCode)
		suite.Equal("004010", result.GS.Version)
	})

	suite.T().Run("adds st start segment", func(t *testing.T) {
		suite.Equal("858", result.ST.TransactionSetIdentifierCode)
		suite.Equal("0001", result.ST.TransactionSetControlNumber)
	})

	suite.T().Run("adds se end segment", func(t *testing.T) {
		// Will need to be updated as more service items are supported
		suite.Equal(64, result.SE.NumberOfIncludedSegments)
		suite.Equal("0001", result.SE.TransactionSetControlNumber)
	})

	suite.T().Run("adds ge end segment", func(t *testing.T) {
		suite.Equal(1, result.GE.NumberOfTransactionSetsIncluded)
		suite.Equal(int64(100001251), result.GE.GroupControlNumber)
	})

	suite.T().Run("adds iea end segment", func(t *testing.T) {
		suite.Equal(1, result.IEA.NumberOfIncludedFunctionalGroups)
		suite.Equal(int64(100001272), result.IEA.InterchangeControlNumber)
	})

	// Test Header Generation
	suite.T().Run("adds bx header segment", func(t *testing.T) {
		suite.IsType(&edisegment.BX{}, result.Header[0])
		bx := result.Header[0].(*edisegment.BX)
		suite.Equal("00", bx.TransactionSetPurposeCode)
		suite.Equal("J", bx.TransactionMethodTypeCode)
		suite.Equal("PP", bx.ShipmentMethodOfPayment)
		suite.Equal(*paymentRequest.MoveTaskOrder.ReferenceID, bx.ShipmentIdentificationNumber)
		suite.Equal("TRUS", bx.StandardCarrierAlphaCode)
		suite.Equal("4", bx.ShipmentQualifier)
	})

	suite.T().Run("does not error out creating EDI from Invoice858", func(t *testing.T) {
		_, err := result.EDIString()
		suite.NoError(err)
	})

	testData := []struct {
		TestName      string
		Qualifier     string
		ExpectedValue string
	}{
		{TestName: "payment request number", Qualifier: "CN", ExpectedValue: paymentRequest.PaymentRequestNumber},
		{TestName: "contract code", Qualifier: "CT", ExpectedValue: "TRUSS_TEST"},
		{TestName: "service member name", Qualifier: "1W", ExpectedValue: serviceMember.ReverseNameLineFormat()},
		{TestName: "service member rank", Qualifier: "ML", ExpectedValue: string(*serviceMember.Rank)},
		{TestName: "service member branch", Qualifier: "3L", ExpectedValue: string(*serviceMember.Affiliation)},
	}

	for idx, data := range testData {
		suite.T().Run(fmt.Sprintf("adds %s to header", data.TestName), func(t *testing.T) {
			suite.IsType(&edisegment.N9{}, result.Header[idx+1])
			n9 := result.Header[idx+1].(*edisegment.N9)
			suite.Equal(data.Qualifier, n9.ReferenceIdentificationQualifier)
			suite.Equal(data.ExpectedValue, n9.ReferenceIdentification)
		})
	}

	suite.T().Run("adds actual pickup date to header", func(t *testing.T) {
		suite.IsType(&edisegment.G62{}, result.Header[6])
		g62Requested := result.Header[6].(*edisegment.G62)
		suite.Equal(10, g62Requested.DateQualifier)
		suite.Equal(requestedPickupDate.Format(testDateFormat), g62Requested.Date)

		suite.IsType(&edisegment.G62{}, result.Header[7])
		g62Scheduled := result.Header[7].(*edisegment.G62)
		suite.Equal(76, g62Scheduled.DateQualifier)
		suite.Equal(scheduledPickupDate.Format(testDateFormat), g62Scheduled.Date)

		suite.IsType(&edisegment.G62{}, result.Header[8])
		g62Actual := result.Header[8].(*edisegment.G62)
		suite.Equal(86, g62Actual.DateQualifier)
		suite.Equal(actualPickupDate.Format(testDateFormat), g62Actual.Date)
	})

	suite.T().Run("adds orders destination address", func(t *testing.T) {
		// name
		expectedDutyStation := paymentRequest.MoveTaskOrder.Orders.NewDutyStation
		transportationOffice, err := models.FetchDutyStationTransportationOffice(suite.DB(), expectedDutyStation.ID)
		suite.FatalNoError(err)
		suite.IsType(&edisegment.N1{}, result.Header[9])
		n1 := result.Header[9].(*edisegment.N1)
		suite.Equal("ST", n1.EntityIdentifierCode)
		suite.Equal(expectedDutyStation.Name, n1.Name)
		suite.Equal("10", n1.IdentificationCodeQualifier)
		suite.Equal(transportationOffice.Gbloc, n1.IdentificationCode)
		// street address
		address := expectedDutyStation.Address
		suite.IsType(&edisegment.N3{}, result.Header[10])
		n3 := result.Header[10].(*edisegment.N3)
		suite.Equal(address.StreetAddress1, n3.AddressInformation1)
		suite.Equal(*address.StreetAddress2, n3.AddressInformation2)
		// city state info
		suite.IsType(&edisegment.N4{}, result.Header[11])
		n4 := result.Header[11].(*edisegment.N4)
		suite.Equal(address.City, n4.CityName)
		suite.Equal(address.State, n4.StateOrProvinceCode)
		suite.Equal(address.PostalCode, n4.PostalCode)
		suite.Equal(*address.Country, n4.CountryCode)
	})

	suite.T().Run("adds orders origin address", func(t *testing.T) {
		// name
		expectedDutyStation := paymentRequest.MoveTaskOrder.Orders.OriginDutyStation
		suite.IsType(&edisegment.N1{}, result.Header[12])
		n1 := result.Header[12].(*edisegment.N1)
		suite.Equal("SF", n1.EntityIdentifierCode)
		suite.Equal(expectedDutyStation.Name, n1.Name)
		suite.Equal("10", n1.IdentificationCodeQualifier)
		suite.Equal(expectedDutyStation.TransportationOffice.Gbloc, n1.IdentificationCode)
		// street address
		address := expectedDutyStation.Address
		suite.IsType(&edisegment.N3{}, result.Header[13])
		n3 := result.Header[13].(*edisegment.N3)
		suite.Equal(address.StreetAddress1, n3.AddressInformation1)
		suite.Equal(*address.StreetAddress2, n3.AddressInformation2)
		// city state info
		suite.IsType(&edisegment.N4{}, result.Header[14])
		n4 := result.Header[14].(*edisegment.N4)
		suite.Equal(address.City, n4.CityName)
		suite.Equal(address.State, n4.StateOrProvinceCode)
		suite.Equal(address.PostalCode, n4.PostalCode)
		suite.Equal(*address.Country, n4.CountryCode)
	})

	suite.T().Run("adds lines of accounting to header", func(t *testing.T) {
		suite.IsType(&edisegment.FA1{}, result.Header[15])
		fa1 := result.Header[15].(*edisegment.FA1)
		suite.Equal("DF", fa1.AgencyQualifierCode)
		suite.IsType(&edisegment.FA2{}, result.Header[16])
		fa2 := result.Header[16].(*edisegment.FA2)
		suite.Equal("TA", fa2.BreakdownStructureDetailCode)
		suite.Equal(*paymentRequest.MoveTaskOrder.Orders.TAC, fa2.FinancialInformationCode)
	})

	var numOfSegments = 5
	for idx, paymentServiceItem := range paymentServiceItems {
		var hierarchicalNumberInt = idx + 1
		var hierarchicalNumber = strconv.Itoa(hierarchicalNumberInt)
		segmentOffset := numOfSegments * idx

		suite.T().Run("adds hl service item segment", func(t *testing.T) {
			suite.IsType(&edisegment.HL{}, result.ServiceItems[segmentOffset])
			hl := result.ServiceItems[segmentOffset].(*edisegment.HL)
			suite.Equal(hierarchicalNumber, hl.HierarchicalIDNumber)
			suite.Equal("|", hl.HierarchicalLevelCode)
		})

		suite.T().Run("adds n9 service item segment", func(t *testing.T) {
			suite.IsType(&edisegment.N9{}, result.ServiceItems[segmentOffset+1])
			n9 := result.ServiceItems[segmentOffset+1].(*edisegment.N9)
			suite.Equal("PO", n9.ReferenceIdentificationQualifier)
			suite.Equal(paymentServiceItem.ID.String(), n9.ReferenceIdentification)
		})
		serviceCode := paymentServiceItem.MTOServiceItem.ReService.Code
		switch serviceCode {
		case models.ReServiceCodeCS, models.ReServiceCodeMS:
			suite.T().Run("adds l5 service item segment", func(t *testing.T) {
				suite.IsType(&edisegment.L5{}, result.ServiceItems[segmentOffset+2])
				l5 := result.ServiceItems[segmentOffset+2].(*edisegment.L5)
				suite.Equal(hierarchicalNumberInt, l5.LadingLineItemNumber)
				suite.Equal(string(serviceCode), l5.LadingDescription)
				suite.Equal("TBD", l5.CommodityCode)
				suite.Equal("D", l5.CommodityCodeQualifier)
			})

			suite.T().Run("adds l0 service item segment", func(t *testing.T) {
				suite.IsType(&edisegment.L0{}, result.ServiceItems[segmentOffset+3])
				l0 := result.ServiceItems[segmentOffset+3].(*edisegment.L0)
				suite.Equal(hierarchicalNumberInt, l0.LadingLineItemNumber)
			})

			suite.T().Run("adds l3 service item segment", func(t *testing.T) {
				suite.IsType(&edisegment.L3{}, result.ServiceItems[segmentOffset+4])
				l3 := result.ServiceItems[segmentOffset+4].(*edisegment.L3)
				suite.Equal(paymentServiceItem.PriceCents.Int64(), l3.PriceCents)
			})
		case models.ReServiceCodeDOP, models.ReServiceCodeDUPK,
			models.ReServiceCodeDPK, models.ReServiceCodeDDP:
			suite.T().Run("adds l5 service item segment", func(t *testing.T) {
				suite.IsType(&edisegment.L5{}, result.ServiceItems[segmentOffset+2])
				l5 := result.ServiceItems[segmentOffset+2].(*edisegment.L5)
				suite.Equal(hierarchicalNumberInt, l5.LadingLineItemNumber)
				suite.Equal(string(serviceCode), l5.LadingDescription)
				suite.Equal("TBD", l5.CommodityCode)
				suite.Equal("D", l5.CommodityCodeQualifier)
			})

			suite.T().Run("adds l0 service item segment", func(t *testing.T) {
				suite.IsType(&edisegment.L0{}, result.ServiceItems[segmentOffset+3])
				l0 := result.ServiceItems[segmentOffset+3].(*edisegment.L0)
				suite.Equal(hierarchicalNumberInt, l0.LadingLineItemNumber)
				suite.Equal(float64(4242), l0.Weight)
				suite.Equal("B", l0.WeightQualifier)
				suite.Equal("L", l0.WeightUnitCode)
			})

			suite.T().Run("adds l3 service item segment", func(t *testing.T) {
				suite.IsType(&edisegment.L3{}, result.ServiceItems[segmentOffset+4])
				l3 := result.ServiceItems[segmentOffset+4].(*edisegment.L3)
				suite.Equal(float64(4242), l3.Weight)
				suite.Equal("B", l3.WeightQualifier)
				suite.Equal(paymentServiceItem.PriceCents.Int64(), l3.PriceCents)
			})
		default:
			suite.T().Run("adds l5 service item segment", func(t *testing.T) {
				suite.IsType(&edisegment.L5{}, result.ServiceItems[segmentOffset+2])
				l5 := result.ServiceItems[segmentOffset+2].(*edisegment.L5)
				suite.Equal(hierarchicalNumberInt, l5.LadingLineItemNumber)
				suite.Equal(string(serviceCode), l5.LadingDescription)
				suite.Equal("TBD", l5.CommodityCode)
				suite.Equal("D", l5.CommodityCodeQualifier)
			})

			suite.T().Run("adds l0 service item segment", func(t *testing.T) {
				suite.IsType(&edisegment.L0{}, result.ServiceItems[segmentOffset+3])
				l0 := result.ServiceItems[segmentOffset+3].(*edisegment.L0)
				suite.Equal(hierarchicalNumberInt, l0.LadingLineItemNumber)
				if serviceCode == models.ReServiceCodeDSH {
					suite.Equal(float64(24245), l0.BilledRatedAsQuantity)
				} else {
					suite.Equal(float64(2424), l0.BilledRatedAsQuantity)
				}
				suite.Equal("DM", l0.BilledRatedAsQualifier)
				suite.Equal(float64(4242), l0.Weight)
				suite.Equal("B", l0.WeightQualifier)
				suite.Equal("L", l0.WeightUnitCode)
			})

			suite.T().Run("adds l3 service item segment", func(t *testing.T) {
				suite.IsType(&edisegment.L3{}, result.ServiceItems[segmentOffset+4])
				l3 := result.ServiceItems[segmentOffset+4].(*edisegment.L3)
				suite.Equal(float64(4242), l3.Weight)
				suite.Equal("B", l3.WeightQualifier)
				suite.Equal(paymentServiceItem.PriceCents.Int64(), l3.PriceCents)
			})
		}
	}
}

func (suite *GHCInvoiceSuite) TestOnlyMsandCsGenerateEdi() {
	generator := NewGHCPaymentRequestInvoiceGenerator(suite.DB())
	basicPaymentServiceItemParams := []testdatagen.CreatePaymentServiceItemParams{
		{
			Key:     models.ServiceItemParamNameContractCode,
			KeyType: models.ServiceItemParamTypeString,
			Value:   testdatagen.DefaultContractCode,
		},
	}
	mto := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{})
	paymentRequest := testdatagen.MakePaymentRequest(suite.DB(), testdatagen.Assertions{
		Move: mto,
		PaymentRequest: models.PaymentRequest{
			IsFinal:         false,
			Status:          models.PaymentRequestStatusPending,
			RejectionReason: nil,
		},
	})

	assertions := testdatagen.Assertions{
		Move:           mto,
		PaymentRequest: paymentRequest,
	}

	testdatagen.MakePaymentServiceItemWithParams(
		suite.DB(),
		models.ReServiceCodeMS,
		basicPaymentServiceItemParams,
		assertions,
	)
	testdatagen.MakePaymentServiceItemWithParams(
		suite.DB(),
		models.ReServiceCodeCS,
		basicPaymentServiceItemParams,
		assertions,
	)

	_, err := generator.Generate(paymentRequest, false)
	suite.NoError(err)
}

func (suite *GHCInvoiceSuite) TestNilValues() {
	currentTime := time.Now()
	basicPaymentServiceItemParams := []testdatagen.CreatePaymentServiceItemParams{
		{
			Key:     models.ServiceItemParamNameContractCode,
			KeyType: models.ServiceItemParamTypeString,
			Value:   testdatagen.DefaultContractCode,
		},
		{
			Key:     models.ServiceItemParamNameRequestedPickupDate,
			KeyType: models.ServiceItemParamTypeDate,
			Value:   currentTime.Format(dateFormat),
		},
		{
			Key:     models.ServiceItemParamNameWeightBilledActual,
			KeyType: models.ServiceItemParamTypeInteger,
			Value:   "4242",
		},
		{
			Key:     models.ServiceItemParamNameDistanceZip3,
			KeyType: models.ServiceItemParamTypeInteger,
			Value:   "2424",
		},
	}

	generator := NewGHCPaymentRequestInvoiceGenerator(suite.DB())
	nilMove := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{})

	nilPaymentRequest := testdatagen.MakePaymentRequest(suite.DB(), testdatagen.Assertions{
		Move: nilMove,
		PaymentRequest: models.PaymentRequest{
			IsFinal:         false,
			Status:          models.PaymentRequestStatusPending,
			RejectionReason: nil,
		},
	})

	assertions := testdatagen.Assertions{
		Move:           nilMove,
		PaymentRequest: nilPaymentRequest,
	}

	testdatagen.MakePaymentServiceItemWithParams(
		suite.DB(),
		models.ReServiceCodeDLH,
		basicPaymentServiceItemParams,
		assertions,
	)

	// This won't work because we don't have PaymentServiceItems on the PaymentRequest right now.
	// nilPaymentRequest.PaymentServiceItems[0].PriceCents = nil

	panicFunc := func() {
		generator.Generate(nilPaymentRequest, false)
	}

	suite.T().Run("nil TAC does not cause panic", func(t *testing.T) {
		oldTAC := nilPaymentRequest.MoveTaskOrder.Orders.TAC
		nilPaymentRequest.MoveTaskOrder.Orders.TAC = nil
		suite.NotPanics(panicFunc)
		nilPaymentRequest.MoveTaskOrder.Orders.TAC = oldTAC
	})

	suite.T().Run("nil country for NewDutyStation does not cause panic", func(t *testing.T) {
		oldCountry := nilPaymentRequest.MoveTaskOrder.Orders.NewDutyStation.Address.Country
		nilPaymentRequest.MoveTaskOrder.Orders.NewDutyStation.Address.Country = nil
		suite.NotPanics(panicFunc)
		nilPaymentRequest.MoveTaskOrder.Orders.NewDutyStation.Address.Country = oldCountry
	})

	suite.T().Run("nil country for OriginDutyStation does not cause panic", func(t *testing.T) {
		oldCountry := nilPaymentRequest.MoveTaskOrder.Orders.OriginDutyStation.Address.Country
		nilPaymentRequest.MoveTaskOrder.Orders.OriginDutyStation.Address.Country = nil
		suite.NotPanics(panicFunc)
		nilPaymentRequest.MoveTaskOrder.Orders.OriginDutyStation.Address.Country = oldCountry
	})

	suite.T().Run("nil reference ID does not cause panic", func(t *testing.T) {
		oldReferenceID := nilPaymentRequest.MoveTaskOrder.ReferenceID
		nilPaymentRequest.MoveTaskOrder.ReferenceID = nil
		suite.NotPanics(panicFunc)
		nilPaymentRequest.MoveTaskOrder.ReferenceID = oldReferenceID
	})

	// TODO: Needs some additional thought since PaymentServiceItems is loaded from the DB in Generate.
	//suite.T().Run("nil PriceCents does not cause panic", func(t *testing.T) {
	//	oldPriceCents := nilPaymentRequest.PaymentServiceItems[0].PriceCents
	//	nilPaymentRequest.PaymentServiceItems[0].PriceCents = nil
	//	suite.NotPanics(panicFunc)
	//	nilPaymentRequest.PaymentServiceItems[0].PriceCents = oldPriceCents
	//})
}