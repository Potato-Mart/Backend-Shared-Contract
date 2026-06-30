# Backend-Shared-Contract

土豆商城 後端生態系共用的契約模組。

本模組刻意設計為僅包含各後端服務共同使用的 domain contract、value struct、列舉、常數、錯誤碼，以及事件負載（event payload）。
模組**不依賴任何資料庫驅動程式或 Web 框架的執行階段相依性**，因此無論是現行單體式架構（monolith）或未來的微服務（microservices），皆可安全地依賴此模組。
只有合約的結構會被寫入本資料夾中

## 最新版本

### V10.0.0

## 模組路徑

```text
github.com/Potato-Mart/Backend-Shared-Contract/v10
```

## 目錄結構

```text
在後續的更新後會加上
```

## 使用方式

在使用方服務的 `go.mod` 中加入：

```go
require github.com/Potato-Mart/Backend-Shared-Contract/v10 v10.0.0
```

接著匯入所需套件，例如：

```go
import (
    "github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/enums"
)
```

## 版本管理

本模組一經發布後，遵循版本控制。
凡是對任何已匯出的內容有破壞性的改變，皆必須升版為新的主版本（major version）
