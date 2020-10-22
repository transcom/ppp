package migrate

import (
	"fmt"
	"log"
	"testing"

	"go.uber.org/zap"

	"github.com/stretchr/testify/suite"

	"github.com/transcom/mymove/pkg/testingsuite"
)

type MigrateSuite struct {
	testingsuite.PopTestSuite
	logger Logger
}

func TestMigrateSuite(t *testing.T) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Panic(err)
	}

	ms := &MigrateSuite{
		PopTestSuite: testingsuite.NewPopTestSuite("migrate"),
		logger:       logger,
	}
	suite.Run(t, ms)
	ms.PopTestSuite.TearDown()
}

func (suite *MigrateSuite) TestCopyStdinPattern() {
	tests := []struct {
		name          string
		tableName     string
		listOfColumns string
		shouldMatch   bool
	}{
		{
			"table name with prefix, letters, and underscores",
			"public.transportation_service_provider_performances",
			"id, performance_period_start, performance_period_end, traffic_distribution_list_id, quality_band, offer_count, best_value_score, transportation_service_provider_id, created_at, updated_at, rate_cycle_start, rate_cycle_end, linehaul_rate, sit_rate",
			true,
		},
		{
			"table name with prefix, letters, numbers, and underscores",
			"public.tariff400ng_full_unpack_rates",
			"id, schedule, rate_millicents, effective_date_lower, effective_date_upper, created_at, updated_at",
			true,
		},
		{
			"bad table name",
			"public.tari+ff400ng_full_unpack_rates",
			"id, schedule, rate_millicents",
			false,
		},
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			// This is sql formatting only for testing
			stmtString := fmt.Sprintf("COPY %s (%s) FROM stdin;", test.tableName, test.listOfColumns) //  G201
			match := copyStdinPattern.FindStringSubmatch(stmtString)

			if !test.shouldMatch {
				suite.Nil(match, "Match found, but wasn't expecting to match")
			} else if suite.NotNil(match, "Match not found when expecting one") {
				expectedMatchArray := []string{
					stmtString,
					"",
					"COPY",
					" ",
					test.tableName,
					" ",
					test.listOfColumns,
					" ",
					"FROM",
					" ",
					"stdin",
					"",
					";",
					"",
				}

				for i, matchItem := range expectedMatchArray {
					suite.Equalf(matchItem, match[i], "Match array item %d not equal", i)
				}

				suite.Len(match, len(expectedMatchArray))
			}
		})
	}
}
