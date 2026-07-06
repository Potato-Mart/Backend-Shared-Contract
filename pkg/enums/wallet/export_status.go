package walletenum

// WalletExportStatus is the lifecycle state of an async export request.
type WalletExportStatus string

const (
	WalletExportStatusPending   WalletExportStatus = "pending"
	WalletExportStatusRunning   WalletExportStatus = "running"
	WalletExportStatusCompleted WalletExportStatus = "completed"
	WalletExportStatusFailed    WalletExportStatus = "failed"
	WalletExportStatusExpired   WalletExportStatus = "expired"
)

func (s WalletExportStatus) IsValid() bool {
	switch s {
	case WalletExportStatusPending, WalletExportStatusRunning,
		WalletExportStatusCompleted, WalletExportStatusFailed,
		WalletExportStatusExpired:
		return true
	}
	return false
}

func (s WalletExportStatus) String() string { return string(s) }
