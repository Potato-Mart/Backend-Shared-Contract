package product_enums

// BarcodeFormat identifies the symbol format of a manufacturer barcode.
type BarcodeFormat string

const (
	BarcodeFormatEAN8    BarcodeFormat = "EAN_8"
	BarcodeFormatEAN13   BarcodeFormat = "EAN_13"
	BarcodeFormatUPCA    BarcodeFormat = "UPC_A"
	BarcodeFormatUPCE    BarcodeFormat = "UPC_E"
	BarcodeFormatCode128 BarcodeFormat = "CODE_128"
	BarcodeFormatQRCode  BarcodeFormat = "QR_CODE"
)

// IsValid reports whether f is a known BarcodeFormat value.
func (f BarcodeFormat) IsValid() bool {
	switch f {
	case BarcodeFormatEAN8, BarcodeFormatEAN13, BarcodeFormatUPCA,
		BarcodeFormatUPCE, BarcodeFormatCode128, BarcodeFormatQRCode:
		return true
	}
	return false
}

func (f BarcodeFormat) String() string { return string(f) }
