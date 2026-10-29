package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v52/pkg/contracts/common/measurement/measurement_enums"
)

func TestMeasurementEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "measurement.NetContentUnit", valid: []stringEnum{measurement_enums.NetContentUnitGram, measurement_enums.NetContentUnitKilogram, measurement_enums.NetContentUnitMillilitre, measurement_enums.NetContentUnitLitre, measurement_enums.NetContentUnitMetre, measurement_enums.NetContentUnitEach}, invalid: measurement_enums.NetContentUnit("G")},
	})
}
