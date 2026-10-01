# Capture timing evidence

`payments/payment.CaptureTimingEvidence` records optional exact `captured_at` and optional `confirmation`. Missing evidence means unknown timing; `paid_at`, envelope occurrence time and local processing time are not capture proof.

`CaptureConfirmationEvidence` records `upper_bound_at`, `source`, `provider` and `evidence_reference`. Sources are `provider_event` and `provider_status_observation`. Payments owns authenticated provenance and payment/provider/account/amount correlation; raw provider responses and credentials are excluded.

`pubsub/payments.PaymentCapturedEventV2` carries required `capture_timing` for the captured amount. Payments validates that at least one form of timing evidence is present and checks consistency when both exist.

An authoritative exact capture time at or before a deadline proves timely capture; an exact time after it proves late capture. A valid confirmation upper bound at or before a deadline proves timely capture. An upper bound after the deadline is ambiguous, since capture may have happened earlier.

Orders owns deadline decisions. Replays preserve immutable evidence; later observations do not erase stronger evidence. This module implements no provider mapping or deadline workflow.
