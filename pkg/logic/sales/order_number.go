// Package saleslogic holds derived helpers for sales contract structs. It
// lives outside pkg/contracts so the contract files stay pure data shapes.
package saleslogic

import "regexp"

// OrderNumberPrefix is the mandatory prefix of every sales order number.
const OrderNumberPrefix = "MAMA"

// OrderNumberPattern is the canonical order-number shape:
// "MAMA" + yymmdd + 6 uppercase alphanumerics, e.g. MAMA260703ABC123.
// Consumers that cannot import this package (e.g. TypeScript frontends)
// mirror the literal and point back at this file.
const OrderNumberPattern = `^MAMA\d{6}[A-Z0-9]{6}$`

var orderNumberRe = regexp.MustCompile(OrderNumberPattern)

// IsValidOrderNumber reports whether s matches OrderNumberPattern. The date
// digits are deliberately not checked for calendar validity so externally
// minted numbers with unusual dates still pass.
func IsValidOrderNumber(s string) bool { return orderNumberRe.MatchString(s) }
