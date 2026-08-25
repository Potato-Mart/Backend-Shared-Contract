package payment

type Mx51PaymentReference struct {
	TransactionID            string `json:"transaction_id,omitempty"`
	ReferenceID              string `json:"reference_id,omitempty"`
	RetrievalReferenceNumber string `json:"retrieval_reference_number,omitempty"`
	SystemTraceAuditNumber   string `json:"system_trace_audit_number,omitempty"`
	AuthorizationCode        string `json:"authorization_code,omitempty"`
	MerchantID               string `json:"merchant_id,omitempty"`
	TerminalID               string `json:"terminal_id,omitempty"`
	SettlementID             string `json:"settlement_id,omitempty"`
}
