package event

// SKUCodeEventVersion is the envelope event_version for every V30 payload
// whose catalogue identity changed from the former database identifier to
// sku_code. V30 publishers
// must not emit those payloads under version 2.
const SKUCodeEventVersion = "3"
