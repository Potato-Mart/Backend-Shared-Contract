package purchase

import "time"

type Supplier struct {
	SupplierID    string    `json:"supplier_id,omitempty"`
	SupplierName  string    `json:"supplier_name,omitempty"`
	SupplierPhone string    `json:"supplier_phone,omitempty"`
	SupplierEmail string    `json:"supplier_email,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
