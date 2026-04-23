package shared

type Health struct {
	Status  string `json:"status"`
	Service string `json:"service"`
	Version string `json:"versioning"`
}
