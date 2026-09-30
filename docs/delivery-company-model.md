# Delivery company model

The current v39 model uses arbitrary administrator-authored configuration and capability flags. See [contract adoption](configurable-contract-adoption.md) for exact fields, masking, protocol and migration requirements. Older release documents describe historical shapes only; preset credentials, connection checks, provider settings, request mappings and company slot catalogues no longer exist.

Supply owns encrypted persistence, authorized CRUD, secret masking and live API execution. The contract contains models only. Orders retains each accepted delivery selection/window; transient available slots are never stored as a company catalogue.
