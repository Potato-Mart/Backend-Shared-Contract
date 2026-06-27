package product

import (
	"encoding/json"
	"strings"

	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/common"
)

const defaultLanguage = "default"

// UnmarshalJSON accepts the pre-migration product display shape while keeping
// the public Product fields canonical. Legacy string description/brand values
// are normalised into localized slices and identifiers.vendor backfills
// identifiers.supplier when the canonical field is absent.
func (p *Product) UnmarshalJSON(data []byte) error {
	fields, err := normalizeProductDisplayJSON(data)
	if err != nil {
		return err
	}
	normalised, err := json.Marshal(fields)
	if err != nil {
		return err
	}
	type alias Product
	var out alias
	if err := json.Unmarshal(normalised, &out); err != nil {
		return err
	}
	*p = Product(out)
	p.normaliseLegacyDisplay()
	return nil
}

// UnmarshalJSON accepts legacy snapshots that carried product id/vendor and a
// flat brand string while new snapshots use sku_code/supplier/localized brand.
func (s *Snapshot) UnmarshalJSON(data []byte) error {
	fields, err := normalizeProductDisplayJSON(data)
	if err != nil {
		return err
	}
	normalised, err := json.Marshal(fields)
	if err != nil {
		return err
	}
	type alias Snapshot
	var out alias
	if err := json.Unmarshal(normalised, &out); err != nil {
		return err
	}
	*s = Snapshot(out)
	if s.Supplier == "" {
		s.Supplier = s.Vendor
	}
	return nil
}

func (p *Product) normaliseLegacyDisplay() {
	if len(p.Description) == 0 && len(p.Localization.Descriptions) > 0 {
		p.Description = p.Localization.Descriptions
	}
	if len(p.Brand) == 0 && len(p.Localization.BrandNames) > 0 {
		p.Brand = p.Localization.BrandNames
	}
	if len(p.Brand) == 0 && strings.TrimSpace(p.BrandKey) != "" {
		p.Brand = []common.LocalizedName{{Language: defaultLanguage, Name: strings.TrimSpace(p.BrandKey)}}
	}
	if p.Identifiers.Supplier == "" {
		p.Identifiers.Supplier = p.Identifiers.Vendor
	}
}

func normalizeProductDisplayJSON(data []byte) (map[string]json.RawMessage, error) {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if raw, ok := fields["description"]; ok {
		normalised, err := normalizeDescriptionJSON(raw)
		if err != nil {
			return nil, err
		}
		fields["description"] = normalised
	}
	if raw, ok := fields["brand"]; ok {
		normalised, err := normalizeBrandJSON(raw)
		if err != nil {
			return nil, err
		}
		fields["brand"] = normalised
	}
	return fields, nil
}

func normalizeDescriptionJSON(raw json.RawMessage) (json.RawMessage, error) {
	if len(raw) == 0 || string(raw) == "null" {
		return raw, nil
	}
	var legacy string
	if err := json.Unmarshal(raw, &legacy); err == nil {
		legacy = strings.TrimSpace(legacy)
		if legacy == "" {
			return json.RawMessage("[]"), nil
		}
		return json.Marshal([]common.LocalizedDescription{{
			Language:    defaultLanguage,
			Description: legacy,
		}})
	}
	var canonical []common.LocalizedDescription
	if err := json.Unmarshal(raw, &canonical); err != nil {
		return nil, err
	}
	return raw, nil
}

func normalizeBrandJSON(raw json.RawMessage) (json.RawMessage, error) {
	if len(raw) == 0 || string(raw) == "null" {
		return raw, nil
	}
	var legacy string
	if err := json.Unmarshal(raw, &legacy); err == nil {
		legacy = strings.TrimSpace(legacy)
		if legacy == "" {
			return json.RawMessage("[]"), nil
		}
		return json.Marshal([]common.LocalizedName{{
			Language: defaultLanguage,
			Name:     legacy,
		}})
	}
	var canonical []common.LocalizedName
	if err := json.Unmarshal(raw, &canonical); err != nil {
		return nil, err
	}
	return raw, nil
}
