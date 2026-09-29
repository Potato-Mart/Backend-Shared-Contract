package product_enums

// BarcodeFormat identifies the symbol format of a product barcode; CODE_128 is the only supported format.
type BarcodeFormat string

const (
	BarcodeFormatCode128 BarcodeFormat = "CODE_128"
)

// IsValid reports whether f is a known BarcodeFormat value.
func (f BarcodeFormat) IsValid() bool {
	switch f {
	case BarcodeFormatCode128:
		return true
	}
	return false
}

func (f BarcodeFormat) String() string { return string(f) }
