# Potato Portal — Supabase Schema & 交互作用總覽

> 本文件由 Claude 依據 `supabase/migrations/` 下 001–016 全部 migration 產生。
> 所有欄位皆標示完整 Postgres 型態（含 NULL / NOT NULL、DEFAULT、CHECK）。
>
> 格式慣例：
> - `col` `TYPE` · NOT NULL · DEFAULT `...` · CHECK (...) — 說明
> - FK 以 `→ table(col)` 表示
> - JSONB 欄位若有慣例 schema，列在下方

---

## 目錄

1. [整體架構](#整體架構)
2. [E-Commerce 訂單 Domain](#1-e-commerce-訂單-domain)
3. [促銷與折扣 Domain](#2-促銷與折扣-domain)
4. [運送 Domain](#3-運送-domain)
5. [顧客 Profile Domain](#4-顧客-profile-domain)
6. [倉儲管理 WMS / Depots Domain](#5-倉儲管理-wms--depots-domain)
7. [網站 CMS Domain](#6-網站-cms-domain)
8. [User / 權限 Domain](#7-user--權限-domain)
9. [Products / Pisell (pre-existing)](#8-products--pisell-pre-existing)
10. [RLS 政策總覽](#rls-政策總覽)
11. [函式與 Trigger](#函式與-trigger)
12. [Storage Buckets](#storage-buckets)
13. [Enums / Check Constraints](#enums--check-constraints)
14. [跨 Domain 交互作用與資料流](#跨-domain-交互作用與資料流)
15. [關鍵效能索引](#關鍵效能索引)

---

## 整體架構

Potato Portal 資料庫是統一的 **B2C + B2B + POS + Import** 商務平台，建構在 Supabase（Postgres + RLS + Storage + Auth）之上。八大 domain：

| Domain | 主要表 |
| --- | --- |
| E-Commerce 訂單 | `orders` / `order_items` / `payments` / `order_status_history` / `order_fulfillment` / `storefront_carts` |
| 促銷 | `coupons` / `promotions` |
| 運送 | `shipping_zones` / `shipping_rates` |
| 顧客 | `customers` / `customer_addresses` |
| 倉儲 WMS | `depots` / `depot_postcode_rules` / `depot_products` / `stock_locations` / `inbound_receipts` / `inbound_items` / `picking_lists` / `picking_list_items` / `outbound_shipments` / `packing_discrepancies` |
| 網站 CMS | `website_pages` / `website_settings` / `website_media` |
| User / 權限 | `user_profiles` / `user_roles` |
| Pre-existing | `products` / `pisell_sales_records`（僅被 ALTER，不在 migration 中建立） |

主要資料流：`products` → `storefront_carts` → `orders` (+ `order_items` / `payments`) → `picking_lists` → `outbound_shipments`。

---

## 1. E-Commerce 訂單 Domain

來源：`001_ecommerce_schema.sql`、`002_customers_schema.sql`（補 `customer_id`）、`004_promotions_schema.sql`（補 `applied_promotions`）、`008_orders_picking_packing.sql`、`009_order_fulfillment.sql`、`010_order_fulfillment_composite_key.sql`。

### 1.1 `orders` — 統一訂單主檔

| 欄位 | 型態 | 說明 |
| --- | --- | --- |
| `id` | `UUID` PK · NOT NULL · DEFAULT `gen_random_uuid()` | 訂單 UUID |
| `order_number` | `TEXT` UNIQUE · NOT NULL | `PT-YYYYMMDD-XXXX`，由 `generate_order_number()` 產生 |
| `channel` | `TEXT` · NOT NULL · DEFAULT `'online'` · CHECK IN (`'online'`,`'pos'`,`'b2b'`,`'relay'`,`'manual'`,`'import'`) | 通路（008 追加 `'import'`） |
| `customer_id` | `UUID` · NULLABLE · FK `→ customers(id) ON DELETE SET NULL` | 2 被 ALTER 改為 UUID+FK（原 001 為 TEXT） |
| `customer_email` | `TEXT` · NULLABLE | guest checkout 可 null |
| `customer_name` | `TEXT` · NULLABLE |  |
| `customer_phone` | `TEXT` · NULLABLE |  |
| `status` | `TEXT` · NOT NULL · DEFAULT `'pending'` · CHECK IN (`'pending'`,`'confirmed'`,`'processing'`,`'shipped'`,`'delivered'`,`'completed'`,`'cancelled'`,`'refunded'`) |  |
| `payment_status` | `TEXT` · NOT NULL · DEFAULT `'unpaid'` · CHECK IN (`'unpaid'`,`'pending'`,`'paid'`,`'partially_refunded'`,`'refunded'`) |  |
| `fulfillment_status` | `TEXT` · NOT NULL · DEFAULT `'unfulfilled'` · CHECK IN (`'unfulfilled'`,`'picking_printed'`,`'packing'`,`'packed'`,`'partial'`,`'fulfilled'`) | 008 追加 `picking_printed` / `packing` / `packed` |
| `subtotal` | `NUMERIC(10,2)` · NOT NULL · DEFAULT `0` | 未折扣小計 |
| `discount_amount` | `NUMERIC(10,2)` · NOT NULL · DEFAULT `0` |  |
| `shipping_amount` | `NUMERIC(10,2)` · NOT NULL · DEFAULT `0` |  |
| `tax_amount` | `NUMERIC(10,2)` · NOT NULL · DEFAULT `0` |  |
| `total` | `NUMERIC(10,2)` · NOT NULL · DEFAULT `0` |  |
| `currency` | `TEXT` · NOT NULL · DEFAULT `'AUD'` |  |
| `shipping_address` | `JSONB` · NULLABLE | schema: `{ name, phone, line1, line2, city, state, postcode, country }` |
| `billing_address` | `JSONB` · NULLABLE | 同上 |
| `shipping_method` | `TEXT` · NULLABLE |  |
| `shipping_zone_id` | `UUID` · NULLABLE · FK `→ shipping_zones(id)` |  |
| `shipping_rate_id` | `UUID` · NULLABLE · FK `→ shipping_rates(id)` |  |
| `tracking_number` | `TEXT` · NULLABLE |  |
| `tracking_url` | `TEXT` · NULLABLE |  |
| `shipped_at` | `TIMESTAMPTZ` · NULLABLE |  |
| `delivered_at` | `TIMESTAMPTZ` · NULLABLE |  |
| `coupon_code` | `TEXT` · NULLABLE | 冗餘保存使用時的 code（即便 coupon 後來刪除仍留紀錄） |
| `coupon_id` | `UUID` · NULLABLE · FK `→ coupons(id)` |  |
| `pos_terminal_id` | `TEXT` · NULLABLE | POS 專用 |
| `cashier_email` | `TEXT` · NULLABLE | POS 專用 |
| `cash_tendered` | `NUMERIC(10,2)` · NULLABLE | POS 現金收款金額 |
| `change_due` | `NUMERIC(10,2)` · NULLABLE | POS 找零 |
| `b2b_customer_id` | `TEXT` · NULLABLE | B2B 客戶 ID（非 FK，歷史欄位） |
| `is_invoiced` | `BOOLEAN` · NULLABLE · DEFAULT `false` |  |
| `customer_note` | `TEXT` · NULLABLE |  |
| `internal_note` | `TEXT` · NULLABLE |  |
| `tags` | `TEXT[]` · NULLABLE | 後台分類標籤 |
| `applied_promotions` | `JSONB` · NOT NULL · DEFAULT `'[]'` | 004 追加；套用 promotions 的快照 |
| `created_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` |  |
| `updated_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` | trigger `orders_updated_at` 自動更新 |
| `confirmed_at` | `TIMESTAMPTZ` · NULLABLE |  |
| `cancelled_at` | `TIMESTAMPTZ` · NULLABLE |  |
| `completed_at` | `TIMESTAMPTZ` · NULLABLE |  |
| `picking_printed_at` | `TIMESTAMPTZ` · NULLABLE | 008 追加 |
| `packed_at` | `TIMESTAMPTZ` · NULLABLE | 008 追加 |

**索引**：
- `idx_orders_number (order_number)`
- `idx_orders_status (status, created_at DESC)` / `idx_orders_status_created`（003 重建）
- `idx_orders_channel (channel, created_at DESC)` / `idx_orders_channel_created`（003）
- `idx_orders_customer (customer_email, created_at DESC)`
- `idx_orders_customer_id (customer_id, created_at DESC)`（003）
- `idx_orders_payment (payment_status)`
- `idx_orders_created (created_at DESC)`
- `idx_orders_fulfillment_created (fulfillment_status, created_at DESC)`（008）

---

### 1.2 `order_items` — 訂單明細

| 欄位 | 型態 | 說明 |
| --- | --- | --- |
| `id` | `UUID` PK · NOT NULL · DEFAULT `gen_random_uuid()` |  |
| `order_id` | `UUID` · NOT NULL · FK `→ orders(id) ON DELETE CASCADE` |  |
| `product_id` | `TEXT` · NULLABLE | 指向 `products.id`（非 FK，因 products.id 亦為 TEXT） |
| `sku` | `TEXT` · NULLABLE |  |
| `product_name` | `TEXT` · NOT NULL | 下單當下的 snapshot |
| `product_brand` | `TEXT` · NULLABLE |  |
| `product_image` | `TEXT` · NULLABLE |  |
| `variant_title` | `TEXT` · NULLABLE | 例：`Frozen 1kg` |
| `unit_price` | `NUMERIC(10,2)` · NOT NULL |  |
| `quantity` | `INTEGER` · NOT NULL · DEFAULT `1` |  |
| `discount_amount` | `NUMERIC(10,2)` · NOT NULL · DEFAULT `0` |  |
| `total` | `NUMERIC(10,2)` · NOT NULL | = `unit_price × quantity − discount_amount` |
| `carton_qty` | `INTEGER` · NULLABLE | B2B 整箱數 |
| `carton_size` | `INTEGER` · NULLABLE | B2B 每箱件數 |
| `properties` | `JSONB` · NULLABLE | 自訂屬性 / line-item notes |
| `created_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` |  |

**索引**：`idx_order_items_order (order_id)` / `idx_order_items_product (product_id)`。

---

### 1.3 `payments` — 金流交易

| 欄位 | 型態 | 說明 |
| --- | --- | --- |
| `id` | `UUID` PK · NOT NULL · DEFAULT `gen_random_uuid()` |  |
| `order_id` | `UUID` · NOT NULL · FK `→ orders(id) ON DELETE CASCADE` |  |
| `amount` | `NUMERIC(10,2)` · NOT NULL |  |
| `currency` | `TEXT` · NOT NULL · DEFAULT `'AUD'` |  |
| `method` | `TEXT` · NOT NULL · CHECK IN (`'card'`,`'cash'`,`'qr'`,`'bank_transfer'`,`'line_pay'`,`'ecpay'`,`'manual'`) |  |
| `gateway` | `TEXT` · NULLABLE | `stripe` / `ecpay` / `linepay` / `square` / `manual`（字串自由，無 CHECK） |
| `gateway_transaction_id` | `TEXT` · NULLABLE |  |
| `gateway_response` | `JSONB` · NULLABLE | gateway 原始回傳 |
| `status` | `TEXT` · NOT NULL · DEFAULT `'pending'` · CHECK IN (`'pending'`,`'processing'`,`'completed'`,`'failed'`,`'cancelled'`,`'refunded'`) |  |
| `paid_at` | `TIMESTAMPTZ` · NULLABLE |  |
| `refunded_at` | `TIMESTAMPTZ` · NULLABLE |  |
| `refund_amount` | `NUMERIC(10,2)` · NULLABLE |  |
| `refund_reason` | `TEXT` · NULLABLE |  |
| `metadata` | `JSONB` · NULLABLE |  |
| `created_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` |  |

**索引**：`idx_payments_order (order_id)` / `idx_payments_gateway (gateway, gateway_transaction_id)` / `idx_payments_status (status)`。

---

### 1.4 `order_status_history` — 狀態變更稽核

| 欄位 | 型態 | 說明 |
| --- | --- | --- |
| `id` | `UUID` PK · NOT NULL · DEFAULT `gen_random_uuid()` |  |
| `order_id` | `UUID` · NOT NULL · FK `→ orders(id) ON DELETE CASCADE` |  |
| `field` | `TEXT` · NOT NULL · DEFAULT `'status'` | 值：`status` / `payment_status` / `fulfillment_status`（文件慣例，無 CHECK） |
| `from_value` | `TEXT` · NULLABLE |  |
| `to_value` | `TEXT` · NOT NULL |  |
| `note` | `TEXT` · NULLABLE |  |
| `actor_email` | `TEXT` · NULLABLE |  |
| `created_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` |  |

**索引**：`idx_order_history_order (order_id, created_at)`。

---

### 1.5 `order_fulfillment` — Pisell/Excel 匯入訂單的側邊狀態

由 009 建立、010 改為複合 PK。

| 欄位 | 型態 | 說明 |
| --- | --- | --- |
| `order_number` | `TEXT` · NOT NULL · PK part 1 |  |
| `order_date` | `DATE` · NOT NULL · PK part 2 | 010 追加；Excel 訂單號可跨日重複 |
| `status` | `TEXT` · NOT NULL · DEFAULT `'unfulfilled'` · CHECK IN (`'unfulfilled'`,`'picking_printed'`,`'packing'`,`'packed'`) |  |
| `picking_printed_at` | `TIMESTAMPTZ` · NULLABLE |  |
| `packed_at` | `TIMESTAMPTZ` · NULLABLE |  |
| `updated_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` |  |

PK：`(order_number, order_date)`。索引：`idx_order_fulfillment_status_picked (status, picking_printed_at DESC)`。

---

### 1.6 `storefront_carts` — 前台購物車

| 欄位 | 型態 | 說明 |
| --- | --- | --- |
| `id` | `UUID` PK · NOT NULL · DEFAULT `gen_random_uuid()` |  |
| `session_id` | `TEXT` UNIQUE · NOT NULL | 瀏覽器 session token |
| `customer_id` | `TEXT` · NULLABLE | guest 為 null；登入後綁定 |
| `items` | `JSONB` · NOT NULL · DEFAULT `'[]'` | schema: `[{ product_id, sku, name, brand, image, price, quantity }]` |
| `coupon_code` | `TEXT` · NULLABLE |  |
| `subtotal` | `NUMERIC(10,2)` · NOT NULL · DEFAULT `0` |  |
| `expires_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now() + interval '30 days'` |  |
| `created_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` |  |
| `updated_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` | trigger `carts_updated_at` |

**索引**：`idx_carts_session (session_id)` / `idx_carts_customer (customer_id)` / `idx_carts_expires (expires_at)`。

---

## 2. 促銷與折扣 Domain

來源：`001_ecommerce_schema.sql`（`coupons`）、`004_promotions_schema.sql`（`promotions`）。

### 2.1 `coupons` — 折扣碼

| 欄位 | 型態 | 說明 |
| --- | --- | --- |
| `id` | `UUID` PK · NOT NULL · DEFAULT `gen_random_uuid()` |  |
| `code` | `TEXT` UNIQUE · NOT NULL | 使用者輸入的 code |
| `description` | `TEXT` · NULLABLE |  |
| `discount_type` | `TEXT` · NOT NULL · CHECK IN (`'percentage'`,`'fixed_amount'`,`'free_shipping'`) |  |
| `discount_value` | `NUMERIC(10,2)` · NOT NULL | percentage=10 代表 10% |
| `min_order_amount` | `NUMERIC(10,2)` · NULLABLE |  |
| `max_discount_amount` | `NUMERIC(10,2)` · NULLABLE | percentage 折扣上限 |
| `usage_limit` | `INTEGER` · NULLABLE | null = 不限 |
| `used_count` | `INTEGER` · NOT NULL · DEFAULT `0` |  |
| `per_customer_limit` | `INTEGER` · NULLABLE · DEFAULT `1` |  |
| `starts_at` | `TIMESTAMPTZ` · NULLABLE |  |
| `expires_at` | `TIMESTAMPTZ` · NULLABLE |  |
| `is_active` | `BOOLEAN` · NOT NULL · DEFAULT `true` |  |
| `applies_to` | `TEXT` · NOT NULL · DEFAULT `'all'` · CHECK IN (`'all'`,`'specific_products'`,`'specific_categories'`) |  |
| `product_ids` | `TEXT[]` · NULLABLE | 指定商品 SKU |
| `category_ids` | `TEXT[]` · NULLABLE | 指定 catalogue |
| `created_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` |  |
| `updated_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` | trigger `coupons_updated_at` |

**索引**：`idx_coupons_code (code)` / `idx_coupons_active (is_active, starts_at, expires_at)` / `idx_coupons_active_code (code) WHERE is_active`（003，熱路徑）。

---

### 2.2 `promotions` — 進階促銷引擎

| 欄位 | 型態 | 說明 |
| --- | --- | --- |
| `id` | `UUID` PK · NOT NULL · DEFAULT `gen_random_uuid()` |  |
| `name` | `TEXT` · NOT NULL |  |
| `description` | `TEXT` · NULLABLE |  |
| `type` | `TEXT` · NOT NULL · CHECK IN (`'auto_discount'`,`'spend_gift'`,`'addon_purchase'`,`'bogo'`,`'bundle'`) | 促銷類型 |
| `min_cart_amount` | `NUMERIC(10,2)` · NULLABLE | 觸發條件 |
| `min_cart_qty` | `INTEGER` · NULLABLE | 觸發條件 |
| `required_product_ids` | `TEXT[]` · NULLABLE | 必須含有的 SKU；bundle 類型也用此 |
| `required_qty_each` | `INTEGER` · NOT NULL · DEFAULT `1` | 每個 required product 需要的數量 |
| `discount_type` | `TEXT` · NULLABLE · CHECK IN (`'percentage'`,`'fixed_amount'`) | auto_discount |
| `discount_value` | `NUMERIC(10,2)` · NULLABLE |  |
| `max_discount` | `NUMERIC(10,2)` · NULLABLE |  |
| `discount_target` | `TEXT` · NOT NULL · DEFAULT `'cart'` · CHECK IN (`'cart'`,`'required_items'`) |  |
| `gift_product_id` | `TEXT` · NULLABLE | spend_gift / bogo |
| `gift_qty` | `INTEGER` · NOT NULL · DEFAULT `1` |  |
| `addon_product_id` | `TEXT` · NULLABLE | addon_purchase |
| `addon_price` | `NUMERIC(10,2)` · NULLABLE |  |
| `addon_max_qty` | `INTEGER` · NOT NULL · DEFAULT `1` |  |
| `bundle_price` | `NUMERIC(10,2)` · NULLABLE | bundle；若 null 則套用 `discount_type`/`discount_value` |
| `priority` | `INTEGER` · NOT NULL · DEFAULT `0` | 數字越大越先套用 |
| `is_stackable` | `BOOLEAN` · NOT NULL · DEFAULT `false` |  |
| `usage_limit` | `INTEGER` · NULLABLE |  |
| `used_count` | `INTEGER` · NOT NULL · DEFAULT `0` |  |
| `per_customer_limit` | `INTEGER` · NOT NULL · DEFAULT `1` |  |
| `starts_at` | `TIMESTAMPTZ` · NULLABLE |  |
| `expires_at` | `TIMESTAMPTZ` · NULLABLE |  |
| `is_active` | `BOOLEAN` · NOT NULL · DEFAULT `true` |  |
| `channels` | `TEXT[]` · NOT NULL · DEFAULT `'{online,pos}'` | 生效的 channel |
| `created_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` |  |
| `updated_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` | trigger `promotions_updated_at` |

**索引**：`idx_promotions_active (priority DESC, is_active) WHERE is_active`。

> 註：`orders.applied_promotions JSONB` 保存下單當下套用的 promotion 快照。

---

## 3. 運送 Domain

來源：`001_ecommerce_schema.sql`。

### 3.1 `shipping_zones`

| 欄位 | 型態 | 說明 |
| --- | --- | --- |
| `id` | `UUID` PK · NOT NULL · DEFAULT `gen_random_uuid()` |  |
| `name` | `TEXT` · NOT NULL | e.g. `"Melbourne Metro"` |
| `states` | `TEXT[]` · NULLABLE | e.g. `{'VIC','NSW'}` |
| `postcodes` | `TEXT[]` · NULLABLE | e.g. `{'3000-3210'}`（range 字串） |
| `is_active` | `BOOLEAN` · NOT NULL · DEFAULT `true` |  |
| `sort_order` | `INTEGER` · NOT NULL · DEFAULT `0` |  |
| `created_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` |  |

Seed：10 個澳洲 zone（Melbourne Metro / Regional VIC / Sydney Metro / Regional NSW / Queensland / South Australia / Western Australia / Tasmania / ACT / Northern Territory）。

---

### 3.2 `shipping_rates`

| 欄位 | 型態 | 說明 |
| --- | --- | --- |
| `id` | `UUID` PK · NOT NULL · DEFAULT `gen_random_uuid()` |  |
| `zone_id` | `UUID` · NOT NULL · FK `→ shipping_zones(id) ON DELETE CASCADE` |  |
| `name` | `TEXT` · NOT NULL | `"Standard"` / `"Express"` / `"Pickup"` |
| `price` | `NUMERIC(10,2)` · NOT NULL |  |
| `free_above` | `NUMERIC(10,2)` · NULLABLE | 免運門檻 |
| `estimated_days` | `TEXT` · NULLABLE | `"3-5 business days"` |
| `storage_type` | `TEXT` · NULLABLE | `'FROZEN'` / `'CHILLED'` / `'DRY'`（null=全部；無 CHECK） |
| `is_active` | `BOOLEAN` · NOT NULL · DEFAULT `true` |  |
| `sort_order` | `INTEGER` · NOT NULL · DEFAULT `0` |  |
| `created_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` |  |

Seed：Melbourne Metro 預設 3 rate（Standard / Express / Pickup）。

---

## 4. 顧客 Profile Domain

來源：`002_customers_schema.sql`。

### 4.1 `customers`

| 欄位 | 型態 | 說明 |
| --- | --- | --- |
| `id` | `UUID` PK · NOT NULL · DEFAULT `gen_random_uuid()` |  |
| `auth_user_id` | `UUID` · NULLABLE · FK `→ auth.users(id) ON DELETE SET NULL` | 匿名／import 顧客可 null |
| `email` | `TEXT` UNIQUE · NOT NULL |  |
| `phone` | `TEXT` · NULLABLE |  |
| `first_name` | `TEXT` · NULLABLE |  |
| `last_name` | `TEXT` · NULLABLE |  |
| `date_of_birth` | `DATE` · NULLABLE |  |
| `notes` | `TEXT` · NULLABLE |  |
| `points` | `INTEGER` · NULLABLE · DEFAULT `0` | 忠誠點數 |
| `tier` | `TEXT` · NULLABLE · DEFAULT `'standard'` · CHECK IN (`'standard'`,`'silver'`,`'gold'`,`'platinum'`) | 由 trigger 自動計算 |
| `tier_spend` | `NUMERIC(10,2)` · NULLABLE · DEFAULT `0` | 累計用於升等的金額 |
| `total_orders` | `INTEGER` · NULLABLE · DEFAULT `0` |  |
| `total_spend` | `NUMERIC(10,2)` · NULLABLE · DEFAULT `0` |  |
| `last_order_at` | `TIMESTAMPTZ` · NULLABLE |  |
| `tags` | `TEXT[]` · NULLABLE · DEFAULT `'{}'` |  |
| `default_shipping_address` | `JSONB` · NULLABLE |  |
| `accepts_marketing` | `BOOLEAN` · NULLABLE · DEFAULT `false` |  |
| `source` | `TEXT` · NULLABLE · DEFAULT `'online'` | `online` / `pos` / `import`（無 CHECK，文件慣例） |
| `is_active` | `BOOLEAN` · NULLABLE · DEFAULT `true` |  |
| `created_at` | `TIMESTAMPTZ` · NULLABLE · DEFAULT `now()` |  |
| `updated_at` | `TIMESTAMPTZ` · NULLABLE · DEFAULT `now()` | trigger `update_customers_updated_at` |

**Triggers**：
- `update_customers_updated_at` → `update_updated_at_column()`（預存於前置 migration）
- `trigger_customer_tier` BEFORE UPDATE OF `tier_spend` → `recalculate_customer_tier()`（≥5000 platinum / ≥2000 gold / ≥500 silver / else standard）

**索引**：`idx_customers_email (email)` / `idx_customers_phone (phone)` / `idx_customers_auth_user (auth_user_id)`（003 版加上 partial `WHERE is_active` / `WHERE auth_user_id IS NOT NULL`）。

---

### 4.2 `customer_addresses`

| 欄位 | 型態 | 說明 |
| --- | --- | --- |
| `id` | `UUID` PK · NOT NULL · DEFAULT `gen_random_uuid()` |  |
| `customer_id` | `UUID` · NULLABLE · FK `→ customers(id) ON DELETE CASCADE` |  |
| `label` | `TEXT` · NULLABLE · DEFAULT `'家'` |  |
| `first_name` | `TEXT` · NULLABLE |  |
| `last_name` | `TEXT` · NULLABLE |  |
| `phone` | `TEXT` · NULLABLE |  |
| `address_line1` | `TEXT` · NULLABLE |  |
| `address_line2` | `TEXT` · NULLABLE |  |
| `city` | `TEXT` · NULLABLE |  |
| `state` | `TEXT` · NULLABLE |  |
| `postcode` | `TEXT` · NULLABLE |  |
| `country` | `TEXT` · NULLABLE · DEFAULT `'AU'` |  |
| `is_default` | `BOOLEAN` · NULLABLE · DEFAULT `false` |  |
| `created_at` | `TIMESTAMPTZ` · NULLABLE · DEFAULT `now()` |  |

**索引**：`idx_customer_addresses_cid (customer_id)`。

---

## 5. 倉儲管理 WMS / Depots Domain

來源：`005_depots_wms_schema.sql`、`012_packing_discrepancies.sql`。

### 5.1 `depots` — 倉庫主檔

| 欄位 | 型態 | 說明 |
| --- | --- | --- |
| `id` | `UUID` PK · NOT NULL · DEFAULT `gen_random_uuid()` |  |
| `name` | `TEXT` · NOT NULL |  |
| `code` | `TEXT` UNIQUE · NOT NULL | `SYD` / `MEL` / `BNE` |
| `address` | `TEXT` · NULLABLE |  |
| `phone` | `TEXT` · NULLABLE |  |
| `is_active` | `BOOLEAN` · NOT NULL · DEFAULT `true` |  |
| `sort_order` | `INTEGER` · NOT NULL · DEFAULT `0` |  |
| `created_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` |  |
| `updated_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` | trigger `depots_updated_at` |

Seed：3 倉（Sydney / Melbourne / Brisbane）。

---

### 5.2 `depot_postcode_rules` — 郵遞區號 → 倉庫路由

| 欄位 | 型態 | 說明 |
| --- | --- | --- |
| `id` | `UUID` PK · NOT NULL · DEFAULT `gen_random_uuid()` |  |
| `depot_id` | `UUID` · NOT NULL · FK `→ depots(id) ON DELETE CASCADE` |  |
| `postcode` | `TEXT` · NOT NULL · UNIQUE | 支援 **前綴匹配**（如 `'3'` 對應 3xxx） |
| `priority` | `INTEGER` · NOT NULL · DEFAULT `0` | 數字越大越先檢查 |

**索引**：`idx_depot_postcode (postcode)`。Seed：VIC(3)→MEL、NSW(2)→SYD、QLD(4)→BNE。

---

### 5.3 `depot_products` — 各倉庫庫存 override

覆蓋 `products.current_stock` 作為前台 postcode-routed 的庫存顯示。

| 欄位 | 型態 | 說明 |
| --- | --- | --- |
| `depot_id` | `UUID` · NOT NULL · FK `→ depots(id) ON DELETE CASCADE` · PK part 1 |  |
| `product_id` | `TEXT` · NOT NULL · PK part 2 | 指向 `products.id` |
| `stock_qty` | `INTEGER` · NOT NULL · DEFAULT `0` |  |
| `is_available` | `BOOLEAN` · NOT NULL · DEFAULT `true` |  |
| `location_code` | `TEXT` · NULLABLE | 倉內儲位 |
| `updated_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` |  |

PK：`(depot_id, product_id)`。索引：`idx_depot_products_product (product_id)`。

---

### 5.4 `stock_locations` — 實體儲位

| 欄位 | 型態 | 說明 |
| --- | --- | --- |
| `id` | `UUID` PK · NOT NULL · DEFAULT `gen_random_uuid()` |  |
| `depot_id` | `UUID` · NOT NULL · FK `→ depots(id) ON DELETE CASCADE` |  |
| `code` | `TEXT` · NOT NULL | `A1-01` / `FROZEN-B3` |
| `name` | `TEXT` · NULLABLE |  |
| `zone` | `TEXT` · NULLABLE | `DRY` / `FROZEN` / `CHILLED`（無 CHECK） |
| `is_active` | `BOOLEAN` · NOT NULL · DEFAULT `true` |  |

Composite UNIQUE：`(depot_id, code)`。

---

### 5.5 `inbound_receipts` — 進貨收貨主檔

| 欄位 | 型態 | 說明 |
| --- | --- | --- |
| `id` | `UUID` PK · NOT NULL · DEFAULT `gen_random_uuid()` |  |
| `depot_id` | `UUID` · NOT NULL · FK `→ depots(id)` |  |
| `reference` | `TEXT` · NULLABLE | 供應商 PO / invoice |
| `supplier` | `TEXT` · NULLABLE |  |
| `eta` | `TIMESTAMPTZ` · NULLABLE |  |
| `operator` | `TEXT` · NULLABLE |  |
| `status` | `TEXT` · NOT NULL · DEFAULT `'draft'` · CHECK IN (`'draft'`,`'confirmed'`) |  |
| `note` | `TEXT` · NULLABLE |  |
| `confirmed_at` | `TIMESTAMPTZ` · NULLABLE |  |
| `created_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` |  |
| `updated_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` | trigger `inbound_updated_at` |

**索引**：`idx_inbound_depot (depot_id, status)` / `idx_inbound_status (status, created_at DESC)`。

---

### 5.6 `inbound_items` — 進貨明細

| 欄位 | 型態 | 說明 |
| --- | --- | --- |
| `id` | `UUID` PK · NOT NULL · DEFAULT `gen_random_uuid()` |  |
| `inbound_receipt_id` | `UUID` · NOT NULL · FK `→ inbound_receipts(id) ON DELETE CASCADE` |  |
| `product_id` | `TEXT` · NOT NULL |  |
| `barcode` | `TEXT` · NULLABLE |  |
| `product_name` | `TEXT` · NULLABLE |  |
| `storage` | `TEXT` · NULLABLE | `FROZEN`/`CHILLED`/`DRY`（無 CHECK） |
| `expected_qty` | `INTEGER` · NOT NULL · DEFAULT `0` |  |
| `received_qty` | `INTEGER` · NOT NULL · DEFAULT `0` |  |
| `location_code` | `TEXT` · NULLABLE |  |
| `created_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` |  |

**索引**：`idx_inbound_items_receipt (inbound_receipt_id)`。

---

### 5.7 `picking_lists` — 撿貨單

| 欄位 | 型態 | 說明 |
| --- | --- | --- |
| `id` | `UUID` PK · NOT NULL · DEFAULT `gen_random_uuid()` |  |
| `depot_id` | `UUID` · NULLABLE · FK `→ depots(id)` |  |
| `order_id` | `UUID` · NULLABLE · FK `→ orders(id)` |  |
| `order_number` | `TEXT` · NULLABLE | 冗餘欄位，支援 Pisell 匯入單（無 orders 記錄時） |
| `status` | `TEXT` · NOT NULL · DEFAULT `'pending'` · CHECK IN (`'pending'`,`'in_progress'`,`'complete'`,`'cancelled'`) |  |
| `assigned_to` | `TEXT` · NULLABLE |  |
| `note` | `TEXT` · NULLABLE |  |
| `created_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` |  |
| `updated_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` | trigger `picking_updated_at` |

**索引**：`idx_picking_depot (depot_id, status)` / `idx_picking_status (status, created_at DESC)` / `idx_picking_order (order_id)`。

---

### 5.8 `picking_list_items` — 撿貨明細

| 欄位 | 型態 | 說明 |
| --- | --- | --- |
| `id` | `UUID` PK · NOT NULL · DEFAULT `gen_random_uuid()` |  |
| `picking_list_id` | `UUID` · NOT NULL · FK `→ picking_lists(id) ON DELETE CASCADE` |  |
| `product_id` | `TEXT` · NOT NULL |  |
| `product_name` | `TEXT` · NOT NULL |  |
| `barcode` | `TEXT` · NULLABLE |  |
| `location` | `TEXT` · NULLABLE |  |
| `quantity_required` | `INTEGER` · NOT NULL |  |
| `quantity_picked` | `INTEGER` · NOT NULL · DEFAULT `0` |  |
| `status` | `TEXT` · NOT NULL · DEFAULT `'pending'` · CHECK IN (`'pending'`,`'partial'`,`'complete'`,`'skipped'`) |  |
| `created_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` |  |

**索引**：`idx_picking_items_list (picking_list_id)`。

---

### 5.9 `outbound_shipments` — 出貨記錄

| 欄位 | 型態 | 說明 |
| --- | --- | --- |
| `id` | `UUID` PK · NOT NULL · DEFAULT `gen_random_uuid()` |  |
| `depot_id` | `UUID` · NULLABLE · FK `→ depots(id)` |  |
| `picking_list_id` | `UUID` · NULLABLE · FK `→ picking_lists(id)` |  |
| `order_id` | `UUID` · NULLABLE · FK `→ orders(id)` |  |
| `order_number` | `TEXT` · NULLABLE |  |
| `customer_name` | `TEXT` · NULLABLE |  |
| `address` | `TEXT` · NULLABLE |  |
| `state` | `TEXT` · NULLABLE |  |
| `operator` | `TEXT` · NOT NULL |  |
| `status` | `TEXT` · NOT NULL · DEFAULT `'packed'` · CHECK IN (`'packed'`,`'dispatched'`) |  |
| `tracking_number` | `TEXT` · NULLABLE |  |
| `note` | `TEXT` · NULLABLE |  |
| `dispatched_at` | `TIMESTAMPTZ` · NULLABLE |  |
| `created_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` |  |

**索引**：`idx_outbound_depot (depot_id, status)` / `idx_outbound_order (order_id)`。

---

### 5.10 `packing_discrepancies` — 打包異常紀錄

來源：`012_packing_discrepancies.sql`。

| 欄位 | 型態 | 說明 |
| --- | --- | --- |
| `id` | `UUID` PK · NOT NULL · DEFAULT `gen_random_uuid()` |  |
| `order_number` | `TEXT` · NOT NULL |  |
| `order_date` | `DATE` · NOT NULL |  |
| `customer_name` | `TEXT` · NULLABLE |  |
| `sku` | `TEXT` · NOT NULL |  |
| `product_name` | `TEXT` · NULLABLE |  |
| `kind` | `TEXT` · NOT NULL · CHECK IN (`'shortage'`,`'overweight'`) | shortage=僅退款；overweight=退款+返庫存 |
| `ordered_qty` | `INTEGER` · NOT NULL |  |
| `scanned_qty` | `INTEGER` · NOT NULL |  |
| `diff_qty` | `INTEGER` · NOT NULL | shortage=ordered-scanned；overweight=scanned-ordered |
| `unit_price` | `NUMERIC(10,2)` · NULLABLE |  |
| `refund_amount` | `NUMERIC(10,2)` · NULLABLE |  |
| `return_to_stock` | `BOOLEAN` · NOT NULL · DEFAULT `false` | 僅 overweight 為 true |
| `notified` | `BOOLEAN` · NOT NULL · DEFAULT `false` | 是否已通知客戶 |
| `recorded_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` |  |
| `recorded_by` | `TEXT` · NULLABLE |  |

**索引**：`idx_packing_discrepancies_order (order_number, order_date)` / `idx_packing_discrepancies_recorded (recorded_at DESC)` / `idx_packing_discrepancies_kind (kind, recorded_at DESC)`。

---

## 6. 網站 CMS Domain

來源：`006_website_cms.sql`、`007_website_media.sql`、`008_website_cms_writes.sql`、`011_website_pages_authenticated_read.sql`。

### 6.1 `website_pages`

| 欄位 | 型態 | 說明 |
| --- | --- | --- |
| `id` | `TEXT` PK · NOT NULL | 頁面 ID（例如 `home` / `about`） |
| `title` | `TEXT` · NOT NULL |  |
| `slug` | `TEXT` UNIQUE · NOT NULL |  |
| `status` | `TEXT` · NOT NULL · DEFAULT `'draft'` · CHECK IN (`'draft'`,`'published'`) |  |
| `blocks` | `JSONB` · NOT NULL · DEFAULT `'[]'` | block-based editor content |
| `seo_title` | `TEXT` · NULLABLE |  |
| `seo_description` | `TEXT` · NULLABLE |  |
| `created_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` |  |
| `updated_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` |  |

**索引**：`idx_website_pages_slug_status (slug, status)`。

---

### 6.2 `website_settings` — KV 式全站設定

| 欄位 | 型態 | 說明 |
| --- | --- | --- |
| `key` | `TEXT` PK · NOT NULL | `header` / `footer` / `homepage` / ... |
| `value` | `JSONB` · NOT NULL · DEFAULT `'{}'` |  |
| `updated_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` |  |

---

### 6.3 `website_media` — 媒體 metadata

| 欄位 | 型態 | 說明 |
| --- | --- | --- |
| `id` | `UUID` PK · NOT NULL · DEFAULT `gen_random_uuid()` |  |
| `filename` | `TEXT` · NOT NULL |  |
| `storage_path` | `TEXT` UNIQUE · NOT NULL | Storage bucket 內路徑 |
| `url` | `TEXT` · NOT NULL | public URL |
| `mime_type` | `TEXT` · NOT NULL |  |
| `size_bytes` | `BIGINT` · NOT NULL |  |
| `width` | `INTEGER` · NULLABLE |  |
| `height` | `INTEGER` · NULLABLE |  |
| `created_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` |  |
| `created_by` | `UUID` · NULLABLE · FK `→ auth.users(id) ON DELETE SET NULL` |  |

**索引**：`idx_website_media_created_at (created_at DESC)`。

---

## 7. User / 權限 Domain

來源：`009_fix_user_roles_rls.sql`、`010_create_user_profiles.sql`。

### 7.1 `user_profiles` — 與 Auth 一對一的 profile

| 欄位 | 型態 | 說明 |
| --- | --- | --- |
| `id` | `UUID` PK · NOT NULL · FK `→ auth.users(id) ON DELETE CASCADE` |  |
| `email` | `TEXT` · NOT NULL |  |
| `display_name` | `TEXT` · NULLABLE |  |
| `status` | `TEXT` · NOT NULL · DEFAULT `'pending'` · CHECK IN (`'pending'`,`'approved'`,`'rejected'`) |  |
| `role` | `TEXT` · NOT NULL · DEFAULT `'user'` · CHECK IN (`'admin'`,`'user'`) |  |
| `created_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` |  |
| `updated_at` | `TIMESTAMPTZ` · NOT NULL · DEFAULT `now()` |  |

**索引**：`idx_user_profiles_email_lower (lower(email))`。

---

### 7.2 `user_roles`（pre-existing）

此表**不在 migration 中建立**，只在 `009_fix_user_roles_rls.sql` 重建 RLS。由 AuthProvider 推知的欄位至少有：
- `email` `TEXT`
- `role` `TEXT`（`'admin'` / `'user'` / ...）
- `approved` `BOOLEAN`

原 RLS 有遞迴問題：policy 在 USING 子句裡 SELECT `user_roles` 自身會觸發 42P17（infinite recursion）。解法：`is_current_user_admin()` 以 `SECURITY DEFINER` + owner=`postgres` 繞過 RLS。

---

## 8. Products / Pisell (pre-existing)

這兩張表**不在 migration 中建立**，只透過後續 migration 補欄位／索引。

### 8.1 `products` — 補強欄位

來自 `013_products_multi_image.sql` / `014_products_created_at.sql` / `015_products_pos_price.sql` 以及 `003_performance_indexes.sql` 推知：

| 欄位 | 型態 | 來源 / 說明 |
| --- | --- | --- |
| `image_urls` | `TEXT[]` · NULLABLE · DEFAULT `NULL` | 013；多圖 URL（保留 `image_url` 作為主圖） |
| `created_at` | `TIMESTAMPTZ` · NULLABLE · DEFAULT `now()` | 014；建立時間（舊資料 backfill `COALESCE(updated_at, now())`） |
| `pos_price` | `NUMERIC(10,2)` · NULLABLE · DEFAULT `NULL` | 015；POS 通路售價（null → fallback `price`） |

既有欄位（由 003 索引推知，非 migration 建立）：`id`、`name`、`brand`、`catalogue`、`en_name`、`storage`（`FROZEN`/`CHILLED`/`DRY`）、`price NUMERIC`、`status`（`'published'` 等）、`current_stock`、`avgweekly_sales`、`image_url`、`updated_at`。

**補強索引**（003 / 013 / 014）：
- `idx_products_storefront (status, price, created_at DESC) WHERE status='published' AND price>0`
- `idx_products_catalogue (catalogue) WHERE status='published'`
- `idx_products_storage (storage) WHERE status='published'`
- `idx_products_brand (brand) WHERE status='published'`
- `idx_products_fts` — GIN on `tsvector(name + en_name + brand)` `WHERE status='published'`
- `idx_products_sales (avgweekly_sales DESC NULLS LAST) WHERE status='published'`
- `idx_products_created_at (created_at DESC)`

---

### 8.2 `pisell_sales_records` — 補強欄位

來自 `011_pisell_delivery_address.sql`：

| 欄位 | 型態 | 說明 |
| --- | --- | --- |
| `delivery_time` | `TEXT` · NULLABLE | Excel `"Delivery time: DD/MM/YYYY HH:MM ~ HH:MM"` |
| `shipping_street` | `TEXT` · NULLABLE | 完整街道（與既有 `shipping_suburb` / `province` / `postcode` 並存） |

其他既有欄位（由相關程式碼推知）：`order_number`、`order_date`、`sku`、`product_name`、`quantity`、`shipping_suburb`、`province`、`postcode`、`customer_name`、`customer_phone` 等（line-item 結構，一張單多 row）。

---

## RLS 政策總覽

| 角色 | 可存取 |
| --- | --- |
| **anon / public** | `coupons` (active) / `promotions` (active) / `depots` (active) / `depot_postcode_rules` / `depot_products` (available) / `website_pages` (published) / `website_settings` / `website_media` |
| **authenticated** | 上述全部 + `website_pages` 全部（含 draft） + `website_*` / `website_media` 寫入 + `order_fulfillment` + `packing_discrepancies` + 自己的 `user_profiles` / `user_roles` + 自己的 `storefront_carts` + email 匹配的 `orders` / `order_items` + 自己的 `customers` / `customer_addresses` |
| **admin**（`is_current_user_admin()` = true） | `user_profiles` / `user_roles` 完整 CRUD |
| **service_role** | 所有表完整存取（多處 `service_*_all` policy） |

政策清單（節錄）：

| 表 | Policy | 角色 | 操作 | USING / WITH CHECK |
| --- | --- | --- | --- | --- |
| `coupons` | `coupons_public_read` | public | SELECT | `is_active = true` |
| `orders` | `orders_customer_read` | public | SELECT | `customer_email = auth.jwt()->>'email'` |
| `order_items` | `order_items_customer_read` | public | SELECT | `order_id IN (SELECT id FROM orders WHERE customer_email = auth.jwt()->>'email')` |
| `storefront_carts` | `carts_own` | public | ALL | `customer_id = auth.uid()::text OR session_id = current_setting('app.session_id', true)` |
| `customers` | `customers_own_select` / `customers_own_update` | public | SELECT / UPDATE | `auth_user_id = auth.uid()` |
| `customer_addresses` | `customer_addresses_own` | public | ALL | `customer_id IN (SELECT id FROM customers WHERE auth_user_id = auth.uid())` |
| `promotions` | `promotions_public_read` | public | SELECT | `is_active = true` |
| `depots` / `depot_postcode_rules` / `depot_products` | `*_public_read` | public | SELECT | 參考各表 |
| `website_pages` | `website_pages_public_read` | public | SELECT | `status = 'published'` |
| `website_pages` | `website_pages_authenticated_read` | authenticated | SELECT | `true` |
| `website_pages` / `website_settings` | `*_authenticated_insert` / `_update` / `_delete` | authenticated | INSERT / UPDATE / DELETE | `true` |
| `website_media` | `website_media_authenticated_insert` / `_delete` | authenticated | INSERT / DELETE | `true` |
| `order_fulfillment` | `order_fulfillment_read` / `_write` | authenticated | SELECT / ALL | `true` |
| `packing_discrepancies` | `packing_discrepancies_read` / `_write` | authenticated | SELECT / ALL | `true` |
| `user_profiles` | `user_profiles_self_read` | authenticated | SELECT | `id = auth.uid()` |
| `user_profiles` | `user_profiles_admin_all` | authenticated | ALL | `is_current_user_admin()` |
| `user_roles` | `user_roles_self_read` | authenticated | SELECT | `lower(email) = lower(auth.jwt()->>'email')` |
| `user_roles` | `user_roles_admin_*` | authenticated | SELECT/INSERT/UPDATE/DELETE | `is_current_user_admin()` |
| 多表 | `service_*_all` | （condition） | ALL | `auth.role() = 'service_role'` |

顧客訂單隔離透過 `customer_email` 比對（非 `customer_id`，因 guest order 可能沒 `customer_id`）。

---

## 函式與 Trigger

### 函式

| 函式 | 回傳 | Language | Security | 用途 |
| --- | --- | --- | --- | --- |
| `generate_order_number()` | `TEXT` | plpgsql | INVOKER（預設） | 產生 `PT-YYYYMMDD-XXXX`，基於 `order_number_seq` sequence |
| `update_updated_at()` | `TRIGGER` | plpgsql | INVOKER | 通用 `NEW.updated_at = now()` |
| `update_updated_at_column()` | `TRIGGER` | plpgsql | INVOKER | 同上（於 002 由前置 migration 提供，給 `customers` 使用） |
| `recalculate_customer_tier()` | `TRIGGER` | plpgsql | INVOKER | 依 `NEW.tier_spend` 改寫 `NEW.tier`（≥5000 platinum / ≥2000 gold / ≥500 silver / else standard） |
| `is_current_user_admin()` | `BOOLEAN` | sql | **DEFINER**（owner=postgres）· STABLE · `search_path=public` | 避免 RLS 遞迴；判斷 `email LIKE '%@potato.com'` **或** `user_roles.role='admin' AND approved=true` |

### Triggers

| Trigger | 表 | 時機 | 函式 |
| --- | --- | --- | --- |
| `orders_updated_at` | `orders` | BEFORE UPDATE | `update_updated_at()` |
| `coupons_updated_at` | `coupons` | BEFORE UPDATE | `update_updated_at()` |
| `carts_updated_at` | `storefront_carts` | BEFORE UPDATE | `update_updated_at()` |
| `promotions_updated_at` | `promotions` | BEFORE UPDATE | `update_updated_at()` |
| `depots_updated_at` | `depots` | BEFORE UPDATE | `update_updated_at()` |
| `inbound_updated_at` | `inbound_receipts` | BEFORE UPDATE | `update_updated_at()` |
| `picking_updated_at` | `picking_lists` | BEFORE UPDATE | `update_updated_at()` |
| `update_customers_updated_at` | `customers` | BEFORE UPDATE | `update_updated_at_column()` |
| `trigger_customer_tier` | `customers` | BEFORE UPDATE OF `tier_spend` | `recalculate_customer_tier()` |

### Sequences

- `order_number_seq` START 1 — 提供 `generate_order_number()` 的序號。

---

## Storage Buckets

來源：`007_website_media.sql`、`016_import_documents_bucket.sql`。

| Bucket | public | file_size_limit | allowed_mime_types | 用途 |
| --- | --- | --- | --- | --- |
| `website-media` | `true` | `52428800`（50 MB） | `image/jpeg`, `image/png`, `image/webp`, `image/gif`, `image/svg+xml`, `video/mp4`, `video/webm` | CMS 媒體庫 |
| `import-documents` | `true` | `20971520`（20 MB） | `application/pdf`, `image/jpeg`, `image/png`, `image/webp`, `image/gif`, `image/heic`, `image/heif` | Declaration Manager 的 BOL / Arrival Notice 附件 |

Storage RLS（`storage.objects`）：
- `*_public_select` — public 可 SELECT 指定 bucket
- `*_authenticated_insert` / `_update`（僅 import-documents） / `_delete` — 已登入使用者可寫

---

## Enums / Check Constraints

以 Postgres CHECK 實作的字串列舉（非 ENUM 型別）：

| 欄位 | CHECK 允許值 |
| --- | --- |
| `orders.channel` | `online` / `pos` / `b2b` / `relay` / `manual` / `import` |
| `orders.status` | `pending` / `confirmed` / `processing` / `shipped` / `delivered` / `completed` / `cancelled` / `refunded` |
| `orders.payment_status` | `unpaid` / `pending` / `paid` / `partially_refunded` / `refunded` |
| `orders.fulfillment_status` | `unfulfilled` / `picking_printed` / `packing` / `packed` / `partial` / `fulfilled` |
| `order_fulfillment.status` | `unfulfilled` / `picking_printed` / `packing` / `packed` |
| `coupons.discount_type` | `percentage` / `fixed_amount` / `free_shipping` |
| `coupons.applies_to` | `all` / `specific_products` / `specific_categories` |
| `payments.method` | `card` / `cash` / `qr` / `bank_transfer` / `line_pay` / `ecpay` / `manual` |
| `payments.status` | `pending` / `processing` / `completed` / `failed` / `cancelled` / `refunded` |
| `promotions.type` | `auto_discount` / `spend_gift` / `addon_purchase` / `bogo` / `bundle` |
| `promotions.discount_type` | `percentage` / `fixed_amount` |
| `promotions.discount_target` | `cart` / `required_items` |
| `inbound_receipts.status` | `draft` / `confirmed` |
| `picking_lists.status` | `pending` / `in_progress` / `complete` / `cancelled` |
| `picking_list_items.status` | `pending` / `partial` / `complete` / `skipped` |
| `outbound_shipments.status` | `packed` / `dispatched` |
| `packing_discrepancies.kind` | `shortage` / `overweight` |
| `website_pages.status` | `draft` / `published` |
| `customers.tier` | `standard` / `silver` / `gold` / `platinum` |
| `user_profiles.status` | `pending` / `approved` / `rejected` |
| `user_profiles.role` | `admin` / `user` |

**無 CHECK 但有文件慣例**（程式端需自行保護）：
- `customers.source`：`online` / `pos` / `import`
- `payments.gateway`：`stripe` / `ecpay` / `linepay` / `square` / `manual`
- `shipping_rates.storage_type`：`FROZEN` / `CHILLED` / `DRY`
- `inbound_items.storage`、`stock_locations.zone`：同上
- `order_status_history.field`：`status` / `payment_status` / `fulfillment_status`

---

## 跨 Domain 交互作用與資料流

### A. 線上 Checkout
1. 瀏覽 `products`（或 `depot_products` 取當倉庫存）。
2. 加入 `storefront_carts`（session-based，登入後綁 `customer_id`）。
3. 套 `coupons`（檢查 `is_active` / `min_order_amount` / `per_customer_limit`）與 `promotions`（依 `priority` + `is_stackable`）。
4. 依 `shipping_address.postcode` → `depot_postcode_rules`（prefix match）→ 決定 `depot_id`。
5. 依 state / postcode 選 `shipping_zones` → `shipping_rates`，`shipping_amount` 寫入 `orders`。
6. 建 `orders`（`status='pending'`, `payment_status='unpaid'`）+ `order_items`；套用的 promotions 快照寫 `orders.applied_promotions`。
7. 付款 gateway 回傳寫入 `payments`；成功後 `orders.payment_status='paid'`。
8. 狀態流轉 `pending → confirmed → processing → shipped → delivered → completed`，每次變更寫 `order_status_history`。

### B. POS Checkout
- `orders.channel='pos'`、`pos_terminal_id` / `cashier_email` / `cash_tendered` / `change_due`。
- 現金付款 `payments.method='cash'`、`paid_at` 立即寫入。
- POS 使用 `products.pos_price`（若 null 則退回 `products.price`）。
- 綁 `customer_id` 時，更新 `customers.tier_spend` 觸發 `trigger_customer_tier` 自動升降 tier。

### C. B2B 訂單
- `orders.channel='b2b'`、`b2b_customer_id`、`is_invoiced`。
- `order_items.carton_qty` / `carton_size` 記大箱單位。

### D. Pisell / Excel 匯入訂單
- 資料來源：`pisell_sales_records`（line-item 扁平結構）。
- `order_number` 會跨日重複 → `order_fulfillment` 用複合 PK `(order_number, order_date)`。
- `delivery_time` + `shipping_street` 於撿貨／包裝標籤上顯示。
- 撿貨狀態 `picking_printed → packing → packed` 同步 `orders.fulfillment_status`（若有對應 order）。

### E. 倉庫作業（WMS）
1. 訂單由 postcode 路由到 `depot_id` → 建 `picking_lists` + `picking_list_items`。
2. 撿貨人員 (`assigned_to`) 掃 `barcode`、更新 `quantity_picked` / `location`。
3. 包裝掃描異常：
   - **shortage**：寫 `packing_discrepancies(kind='shortage')`，`return_to_stock=false`，僅退款。
   - **overweight**：寫 `kind='overweight'`，`return_to_stock=true`，退款 + 回補 `depot_products.stock_qty`。
4. 建 `outbound_shipments`（`tracking_number`、`dispatched_at`），`orders.status='shipped'` 同步。

### F. 進貨（Inbound）
`inbound_receipts` (`status='draft'`) → 逐筆 `inbound_items.received_qty` → `status='confirmed'`、`confirmed_at=now()` → 更新 `depot_products.stock_qty`。

### G. 顧客忠誠
- 完成訂單 → `customers.total_orders++`、`total_spend += total`、`tier_spend` 累計 → trigger 自動升降 `tier`。
- `customers.points` 由後端邏輯維護（migration 未定義 trigger）。

### H. 促銷引擎
- Checkout 一次讀所有 `is_active=true` 且 `channels[]` 符合的 `promotions`，依 `priority DESC` 排序。
- `is_stackable=false` 的 promotion 套用後阻擋後續同類。
- 計算結果寫回 `orders.discount_amount`、`order_items.discount_amount`、`orders.applied_promotions`。

### I. 管理 / CMS
- 編輯 `website_pages.blocks` JSONB → `status='published'` → public 可讀。
- 上傳媒體到 Storage `website-media` → 同步寫 `website_media` metadata。
- Declaration Manager 上傳 BOL / Arrival Notice 到 `import-documents` bucket。

### J. 權限審核
- 新使用者 signup → `user_profiles.status='pending'`。
- Admin（`is_current_user_admin()`=true）approve 後 `status='approved'`，才具後台操作權限。

---

## 關鍵效能索引

### 熱路徑（高併發）
- `idx_coupons_active_code` — checkout 折扣碼查找
- `idx_products_storefront` — 前台列表（partial `WHERE status='published' AND price>0`）
- `idx_depot_postcode` — postcode 即時路由
- `idx_orders_customer` / `idx_orders_status` / `idx_orders_channel` — 後台訂單多種篩選
- `idx_picking_status` — 倉庫「待撿貨」清單

### 全文搜尋
- `idx_products_fts` — GIN on `tsvector(name + en_name + brand)` `WHERE status='published'`

### 排序 / 分頁
- `idx_products_created_at` / `idx_products_sales` / `idx_website_media_created_at` / `idx_order_fulfillment_status_picked`

### 複合索引（「狀態 + 時間」雙欄位排序）
- `idx_orders_fulfillment_created` / `idx_orders_status_created` / `idx_orders_channel_created`

---

_本文件由 Claude 依 `supabase/migrations/` 001–016 產生，日期 2026-04-20。_
_每個欄位均標注 Postgres 型態、NULL / NOT NULL、DEFAULT 與 CHECK（若有），並列出索引與 FK。_
_若 migration 有更新，請同步維護。_

