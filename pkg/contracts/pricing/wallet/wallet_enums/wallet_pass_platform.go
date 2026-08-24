package wallet_enums

// WalletPassPlatform identifies an external membership-pass wallet provider.
type WalletPassPlatform string

const (
	WalletPassPlatformGoogle WalletPassPlatform = "google_wallet"
	WalletPassPlatformApple  WalletPassPlatform = "apple_wallet"
)

func (p WalletPassPlatform) IsValid() bool {
	return p == WalletPassPlatformGoogle || p == WalletPassPlatformApple
}
func (p WalletPassPlatform) String() string { return string(p) }
