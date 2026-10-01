# Delivery company model

`pkg/contracts/supply/courier` defines Supply-owned delivery company records. Supply owns validation, authorized CRUD, encrypted persistence, secret masking and live API execution; Shared Contract contains data models only.

`DeliveryCompany.code` is an immutable company identity. Its record includes name, enabled state, revision, country coverage, default instructions, service areas, configuration and capabilities. Revision identifies the effective configuration, service areas and metadata.

Configuration is an array of arbitrary company-scoped keys, typed JSON values and `sensitive` flags. Supply validates duplicate keys and the normalized delivery protocol. Sensitive values require encrypted storage and masked read projections: the model does not mask values during JSON encoding.

Capabilities describe booking, tracking, proof of delivery, refrigerated delivery, provider coverage and provider slots. They indicate configured support, not verified readiness, live availability or coverage. False never implies support. Configured APIs use the shared delivery request, schedule and selection models.

Custom metadata contains non-secret typed JSON values. It is inert and must never affect API requests; consumers preserve unknown value types and values.

Provider slots are transient authoritative selectable windows. Orders retains each accepted delivery selection and window; company records contain no stored slot catalogue.
