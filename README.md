# Backend-Shared-Contract

土豆商城 後端生態系共用的契約模組。

本模組刻意設計為僅包含各後端服務共同使用的型別、列舉、DTO、錯誤碼，以及事件負載（event payload）。  
模組**不依賴任何資料庫驅動程式或 Web 框架的執行階段相依性**，因此無論是現行單體式架構（monolith）或未來的微服務（microservices），皆可安全地依賴此模組。

## 最新版本

### V5.4.0

## 模組路徑

```text
github.com/Potato-Mart/Backend-Shared-Contract/v5
````

## 目錄結構

```text
在後續的更新後會加上
```

## 使用方式

在使用方服務的 `go.mod` 中加入：

```go
require github.com/Potato-Mart/Backend-Shared-Contract/v5 v5.4.0
```

接著匯入所需套件，例如：

```go
import (
    "github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/enums"
)
```

## 版本管理

本模組一經發布後，遵循版本控制。
凡是對任何已匯出的內容有破壞性的改變，皆必須升版為新的主版本（major version）。

自 V3.0.0 開始，所有的發布內容會自動化生成

自 V4.0.0 開始，付費接口已經泛化，可用任何第三方進行合作

自 V5.0.0 開始

- `pkg/contracts/stockops`：新增內部庫存端點路徑常數 `PathReserve`、`PathCommit`、`PathRelease`（提供方 Backend-Operations）。
- `pkg/serviceauth`：新增服務權杖端點路徑常數 `PathToken`（提供方 Backend-Management）。
- `pkg/contracts/pricing`：新增跨服務報價契約 `QuoteRequest` / `QuoteResponse` 與路徑常數 `PathQuote`（提供方 Backend-Management、使用方 Backend-Commerce，scope `pricing:quote`）。
- `pkg/apiresponse`：於 `APIResponse` 文件註解明示內部服務間端點（token、stockops、pricing）同樣使用此回應信封。