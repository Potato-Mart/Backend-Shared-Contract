package payments

// Adyen/nexo constants used by services when translating the shared
// contract to github.com/adyen/adyen-go-api-library request models.
const (
	AdyenProtocolVersion = "3.0"

	AdyenTenderOptionMOTO        = "MOTO"
	AdyenTenderOptionAskGratuity = "AskGratuity"

	AdyenResponseResultSuccess = "Success"
	AdyenResponseResultFailure = "Failure"
	AdyenResponseResultPartial = "Partial"
)
