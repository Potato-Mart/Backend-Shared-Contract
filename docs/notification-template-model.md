# Notification template model

The `notification/template` package defines reusable authoring records for
Notification and its API consumers. The owning service supplies validation,
authorization, storage, rendering, translation and workflow. These JSON structs
do not enforce their invariants or implement an API. Notification exposes the
records through its own DTOs and derives frontend schemas from its OpenAPI.

## Identity and independent versions

| Record or value | Identity and meaning |
| --- | --- |
| `NotificationTemplate` | Mutable definition/current content, keyed by `(country_code, code)`; `id` is opaque. Country and code are immutable. `revision` is positive optimistic concurrency for accepted definition/content writes. |
| `NotificationTemplateVersion` | Immutable `(country_code, template_code, published_version)`; `source_revision` identifies the definition revision published. `published_at` is UTC and `published_by` an opaque actor ID. |
| `NotificationTemplateReference` | Required country, code and exact positive published version. Zero never means latest. |
| `NotificationTemplateBinding` | At most one active record per `(country_code, topic_code, channel, purpose_code)`; its reference must match country and channel. `revision` protects binding updates. |
| `NotificationContent.schema_version` | Positive document/token interpretation version, initially `1`; independent of every record revision and the Go module release. Unsupported versions fail closed in the service. |
| Source/target fingerprints | Opaque server-calculated content evidence; not revision counters or client-generated authority. |

Country uses existing `geography.CountryCode` (canonical ISO alpha-2). The owner
rejects empty/global/unsupported country codes and mismatched references. There
is no cross-country fallback. Same template codes can exist independently in
different countries. Copying across countries creates a new identity and new
review/publication evidence. A restore copies an old publication into a new
mutable revision and passes normal publication checks; old versions never change.

Topic and purpose codes remain open, backend-managed strings. An omitted purpose
means the ordinary purpose, not a wildcard. A security purpose never grants
authorization. Bindings and template activity alone cannot authorize delivery.
The content channel remains fixed after first publication.

Every lookup/cache/reference includes country. Publication caches also include
the exact version; localized/rendered caches include locale and applicable
content/schema identity. Business-send country comes from verified owning-domain
context, not recipient language, email or merely the editor's country.

## Localized channel content

`NotificationContent` is one channel's unrendered content with schema version,
source language/fingerprint, placeholder declarations and localized variants.
The enclosing template or service-owned manual draft supplies country identity.
Schema 1 supports `email`, `sms` and `push` from the existing channel enum; the
other existing enum values are not authoring capabilities in this schema.

Language strings use the existing BCP 47 representation: `en`, `zh-TW`, `zh-CN`.
Any one can be the source; the other two are translated directly from that
source. Country is independent of language. Identity's existing preferred-language
values remain unchanged; Notification maps them explicitly.

Each variant has exactly one content arm matching the enclosing channel. Email
uses `EmailTemplateContent`; SMS/push reuse `SMSNotification`/`PushNotification`.
Their fields hold unrendered copy in this context. Source origin is `authored`;
target origin is `translated`, including human-refined translation. Origin does
not establish approval. The source has no `review`; both targets require current
`TranslationReview` evidence. Arrays are non-null; `placeholders` may be empty.

For the first authoring release, completed Save requires all three complete
locales and current confirmation of both targets. Incomplete edits stay local;
Translate returns a nonpersisting proposal and modal Confirm only stages edits.
Save performs final validation and atomically stamps review. Send/publish remain
separate owner-authorized operations. Service DTOs, manual-message aggregate,
recipient selection, send commands and idempotency are not shared contracts.

## Translation evidence and privacy

`TranslationReview` records source and final target fingerprints plus UTC review
time and opaque reviewer ID. The owner stamps these values and binds the evidence
to country/resource identity; it never accepts a forged client review record.
Clients treat fingerprints as opaque. Notification defines deterministic hashing
and canonicalization in its service contract; fixture digests are synthetic.

Source fingerprint input covers country/resource context, schema, channel,
source language, ordered stable block/field identity and kind, protected section
codes, translatable leaves, and placeholder declarations/token structure. Target
evidence covers target language and final translated leaves/token/structure in
that context. Neither fingerprint includes counters, audit timestamps, purely
visual values, URL values or media references. Those values participate in the
separate full-render preview/content digest maintained by Notification.

- Source language/text, semantic structure or placeholder changes invalidate
  both target reviews. The owner compares current source and target fingerprints.
- Editing one target invalidates its review until final confirmation. An absent
  review remains distinct from a stale retained review; service DTOs derive status.
- Visual/link/media edits preserve otherwise-current translation evidence but
  invalidate render preview. They still require authorization, safety validation
  and the normal save/publication checks. Changing a system-section code changes
  semantic structure and requires review again.

Translation receives only allowed textual leaves from unrendered authoring
content. Preserve field paths, block IDs/order/kinds, token occurrences, URLs,
media references and protected sections. No actual customer data, recipient list,
OTP, reset token, claim material or rendered personal content goes to a translator
or synthetic preview. The owner must also screen free text; a placeholder's
`sensitive` flag is not a guarantee that surrounding text is safe.

## Typed placeholders

`TemplatePlaceholder` has `key`, `value_type` and required boolean `sensitive`.
Keys are unique and match `[a-z][a-z0-9_]*(\.[a-z][a-z0-9_]*)*`. Schema 1 uses
literal scalar tokens `{{key}}`, with no functions, loops, includes or evaluation.
Every referenced token must be declared and supplied with a correctly typed value
for rendering. Declarations contain no values, defaults or customer examples.

| Value type | Runtime meaning, resolved by Notification |
| --- | --- |
| `text` | Plain text, including protected text where `sensitive` is true; escaped for its output context. |
| `integer` | Integral value, formatted by locale without losing precision. |
| `money` | Existing common minor-unit amount/currency model; never floating-point money. |
| `timestamp` | UTC instant, formatted by locale and the owner's time policy. |
| `url` | Validated URL only in a nontranslatable URL slot; never resolved for translation. |

Missing, extra, moved or changed tokens in translated textual leaves are rejected
by the owner. Field-level token occurrences are preserved. Links and other
nontranslatable token slots remain unchanged. Runtime token values, transport
types, URL allowlists and escaping implementations stay service-owned.

## Email blocks and presentation

Email subject, preview text and the allowed text/alt leaves are translatable.
Each block has a nonempty document-unique stable `id`; IDs remain stable across
all locales. There is no raw HTML/CSS authoring field.

| Kind | Payload |
| --- | --- |
| `heading` | `text` and `level` (1..3). |
| `paragraph` | `text`. |
| `button` | Translatable `text` and nontranslatable `action_url`. |
| `image` | Nontranslatable `media` (`ObjectMedia`) and `alt_text`. Explicit empty alt denotes decorative imagery; absent alt is different. |
| `divider` | No payload. |
| `system_section` | Only an open `section_code`; no runtime payload or user-controlled style. |

Pointer leaves retain omitted versus explicitly empty values. Notification
rejects irrelevant payload fields, unsupported block kinds, invalid links and
unsafe media. The owner defines allowed protected sections and required
placement/count/order for order details, authentication/expiry, claims and
unsubscribe. These appear locked in the editor and are rendered from protected
owner data and reviewed system locale copy. Their identifiers do not embed data.

One optional `email_theme` on the channel document applies across locales. The
owner resolves explicit theme values at publication; it rejects themes on SMS
or push. `EmailTemplateTheme` contains four `#RRGGBB` colors (`background_color`,
`content_color`, `text_color`, `accent_color`) and these bounded tokens:

| Field | Values |
| --- | --- |
| `font_family` | `sans`, `serif` |
| `text_size` | `small`, `normal`, `large` |
| `spacing` | `compact`, `normal`, `relaxed` |
| `button_shape` | `square`, `rounded` |

Editable blocks may carry `EmailBlockStyle` with optional `alignment`
(`start`, `center`, `end`) and `spacing` from the same scale. Omitted tokens
inherit owner defaults. All variants retain identical presentation for each
block ID. The renderer owns actual sizes, font stacks, contrast/accessibility
checks and HTML generation; remote fonts, free CSS and arbitrary style maps
are not represented.

## Ownership, compatibility and verification

Existing role/country primitives suffice. Notification implements a template
capability allowing country admins and marketing staff to edit every topic and
email/SMS/push channel within their validated country. Security-sensitive topics
still permit copy editing while system sections/runtime data stay protected.
Superadmin visibility spans countries. Publication/send authorization remains
separate; market-scoped recipient/campaign permissions do not broaden. Generic
market filtering must not silently restrict country-level template records.

All additions are under `notification/template` and its leaf enum package.
Existing delivery, marketing, country, locale and role models are unchanged.
No shared manual-send snapshot is introduced: the owner pins country, publication,
safe content and resolved locale at acceptance, with credentials handled through
its protected runtime mechanisms. Provider configuration, APIs, MongoDB mappings,
hashing and validation are absent from this module.

`testdata/notification_content.json` covers all nine source/channel combinations;
`testdata/country_template_records.json` covers independent AU/TW identities,
publications and bindings. Both live beside the package tests and contain
synthetic content. Tests lock JSON round trips, presence semantics, country
references, independent counters, enums, package/audit policy and model boundaries.
They do not prove authorization, translation, hashing or delivery behavior.

Adoption is sequential: actual Contract publication first; backend confirms the
released model and both module/CI pins before implementation; completed backend
API handoff precedes frontend implementation. Rich/security/unsubscribe paths
retain current behavior until protected-section parity is verified. After each
consumer cutover and rollback window, the owner removes its superseded paths,
flags, fixtures and stale docs. This additive release introduces no deprecated
aliases and removes no still-consumed model. Exported removals require a separate
major release with consumer evidence.
