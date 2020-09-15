package edisegment

import (
	"testing"
)

func (suite *SegmentSuite) TestValidateL3() {
	validWeightL3 := L3{
		Weight:          300.0,
		WeightQualifier: "B",
		PriceCents:      100,
	}

	suite.T().Run("validate success weight", func(t *testing.T) {
		err := suite.validator.Struct(validWeightL3)
		suite.NoError(err)
	})

	suite.T().Run("validate failure 1", func(t *testing.T) {
		l3 := L3{
			Weight:     300.0,
			PriceCents: 100,
		}

		err := suite.validator.Struct(l3)
		suite.ValidateError(err, "WeightQualifier", "required_with")
		suite.ValidateErrorLen(err, 1)
	})

	suite.T().Run("validate failure 2", func(t *testing.T) {
		l3 := L3{
			Weight:          300.0,
			WeightQualifier: "INVALID",
			PriceCents:      100,
		}

		err := suite.validator.Struct(l3)
		suite.ValidateError(err, "WeightQualifier", "eq")
		suite.ValidateErrorLen(err, 1)
	})
}
