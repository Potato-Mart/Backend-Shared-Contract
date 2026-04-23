# Backend-Shared-Contract

土豆商城 後端生態系共用的契約模組。

本模組刻意設計為僅包含各後端服務共同使用的型別、列舉、DTO、錯誤碼，以及事件負載（event payload）。  
模組**不依賴任何資料庫驅動程式或 Web 框架的執行階段相依性**，因此無論是現行單體式架構（monolith）或未來的微服務（microservices），皆可安全地依賴此模組。

## 模組路徑

```text
github.com/Potato-Mart/Backend-Shared-Contract
````

## 目錄結構

```text
shared-contract-go/
├── README.md
├── go.mod
├── go.sum
├── /docs
│   ├── architecture.md
│   ├── versioning.md
│   └── error-handling.md
├── /pkg
│   ├── /common
│   │   ├── address.go
│   │   ├── audit.go
│   │   ├── money.go
│   │   ├── pagination.go
│   │   └── metadata.go
│   │
│   ├── /enums
│   │   ├── order.go
│   │   ├── customer.go
│   │   ├── warehouse.go
│   │   ├── cms.go
│   │   └── storage.go
│   │
│   ├── /errors
│   │   ├── code.go
│   │   ├── error.go
│   │   └── response.go
│   │
│   ├── /contracts
│   │   ├── /shared
│   │   │   ├── health.go
│   │   │   └── media.go
│   │   │
│   │   ├── /identity
│   │   │   └── user.go
│   │   │
│   │   ├── /customer
│   │   │   ├── customer.go
│   │   │   └── address.go
│   │   │
│   │   ├── /order
│   │   │   ├── order.go
│   │   │   ├── payment.go
│   │   │   ├── cart.go
│   │   │   └── history.go
│   │   │
│   │   ├── /promotion
│   │   │   ├── coupon.go
│   │   │   └── promotion.go
│   │   │
│   │   ├── /shipping
│   │   │   ├── zone.go
│   │   │   └── rate.go
│   │   │
│   │   ├── /warehouse
│   │   │   ├── depot.go
│   │   │   ├── inbound.go
│   │   │   ├── picking.go
│   │   │   └── shipment.go
│   │   │
│   │   └── /cms
│   │       ├── page.go
│   │       └── setting.go
│   │
│   └── /versioning
│       └── version.go
└── /examples
    ├── order_example.go
    └── error_example.go
```

## 使用方式

在使用方服務的 `go.mod` 中加入：

```go
require github.com/Potato-Mart/Backend-Shared-Contract v0.1.0
```

接著匯入所需套件，例如：

```go
import (
    "github.com/Potato-Mart/Backend-Shared-Contract/pkg/dto"
    "github.com/Potato-Mart/Backend-Shared-Contract/pkg/enums"
)
```

## 使用範例

以下示範另一個後端服務如何直接使用共用的 DTO、列舉型別與錯誤碼：

```go
package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "github.com/Potato-Mart/Backend-Shared-Contract/pkg/apierror"
    "github.com/Potato-Mart/Backend-Shared-Contract/pkg/dto"
    "github.com/Potato-Mart/Backend-Shared-Contract/pkg/enums"
)

func CreateOrder(c *gin.Context) {
    var req dto.CreateOrderRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, apierror.ErrInvalidRequest.ToResponse())
        return
    }

    if req.PaymentMethod != enums.PaymentMethodCard {
        c.JSON(http.StatusBadRequest, apierror.ErrInvalidPaymentMethod.ToResponse())
        return
    }

    resp := dto.CreateOrderResponse{
        OrderID: "ORD-20250101-0001",
        Status:  enums.OrderStatusPending,
    }

    c.JSON(http.StatusOK, resp)
}
```

## 版本管理

本模組一經發布後，遵循版本控制。
凡是對任何已匯出的內容有破壞性的改變，皆必須升版為新的主版本（major version）。
