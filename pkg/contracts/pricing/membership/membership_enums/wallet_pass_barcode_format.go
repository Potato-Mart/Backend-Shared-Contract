package membership_enums

// WalletPassBarcodeFormat identifies a canonical membership-pass barcode.
type WalletPassBarcodeFormat string

const WalletPassBarcodeFormatCode128 WalletPassBarcodeFormat = "code_128"

func (f WalletPassBarcodeFormat) IsValid() bool  { return f == WalletPassBarcodeFormatCode128 }
func (f WalletPassBarcodeFormat) String() string { return string(f) }
