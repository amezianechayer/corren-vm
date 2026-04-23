package vm

import (
	"strings"

	"github.com/amezianechayer/corren-vm/core"
)

type NoRibaConstraint struct{}

func (c *NoRibaConstraint) Name() string { return "NO_RIBA" }

func (c *NoRibaConstraint) Validate(funding core.Funding, dest core.Account) error {
	name := strings.ToUpper(string(funding.Asset))
	if strings.Contains(name, "RIBA") || strings.Contains(name, "INTEREST") {
		return &core.ShariaViolation{
			Constraint: c.Name(),
			Reason:     string(funding.Asset) + " is a riba-bearing asset",
		}
	}
	return nil
}

type NoGhararConstraint struct{}

func (c *NoGhararConstraint) Name() string { return "NO_GHARAR" }

func (c *NoGhararConstraint) Validate(funding core.Funding, dest core.Account) error {
	if funding.Infinite {
		return nil
	}
	var total uint64
	for _, part := range funding.Parts {
		total += part.Amount
	}
	if total == 0 {
		return &core.ShariaViolation{
			Constraint: c.Name(),
			Reason:     "transfer of zero amount represents gharar (prohibited uncertainty)",
		}
	}
	return nil
}

type AssetBackedConstraint struct {
	AllowedAssets map[string]bool
}

func NewAssetBackedConstraint(assets []string) *AssetBackedConstraint {
	m := make(map[string]bool, len(assets))
	for _, a := range assets {
		m[strings.ToUpper(a)] = true
	}
	return &AssetBackedConstraint{AllowedAssets: m}
}

func (c *AssetBackedConstraint) Name() string { return "ASSET_BACKED" }

func (c *AssetBackedConstraint) Validate(funding core.Funding, dest core.Account) error {
	if !funding.Infinite || len(c.AllowedAssets) == 0 {
		return nil
	}
	asset := strings.ToUpper(string(funding.Asset))
	if !c.AllowedAssets[asset] {
		return &core.ShariaViolation{
			Constraint: c.Name(),
			Reason:     string(funding.Asset) + " is not in the approved asset registry",
		}
	}
	return nil
}
