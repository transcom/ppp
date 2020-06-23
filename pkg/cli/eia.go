package cli

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	// EIAKeyFlag is the EIA Key Flag
	EIAKeyFlag string = "eia-key"
	// EIAURLFlag is the EIA URL Flag
	EIAURLFlag string = "eia-url"
)

// InitEIAFlags initializes EIA command line flags
func InitEIAFlags(flag *pflag.FlagSet) {
	flag.String(EIAKeyFlag, "", "Key for Energy Information Administration (EIA) Open Data API")
	flag.String(EIAURLFlag, "", "URL for Energy Information Administration (EIA) Open Data API")
}

// CheckEIA validates EIA command line flags
func CheckEIA(v *viper.Viper) error {
	eiaKey := v.GetString(EIAKeyFlag)
	if len(eiaKey) != 32 {
		return fmt.Errorf("expected EIA Open Data API key to be 32 characters long; key is %d chars", len(eiaKey))
	}

	eiaURL := v.GetString(EIAURLFlag)
	if eiaURL != "https://api.eia.gov/series/" {
		return fmt.Errorf("invalid EIA Open Data URL %s, expecting https://api.eia.gov/series/", eiaURL)
	}
	return nil
}
