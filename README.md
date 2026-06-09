# mongox

MongoDB 客户端管理工具，基于 mongo-driver/v2，支持多数据源。

## 安装

```bash
go get github.com/go-xuan/mongox
```

## 快速开始

在 `conf/mongo.yaml` 中配置：

```yaml
source: "default"
enable: true
uri: "mongodb://admin:admin@127.0.0.1:27017/demo?authSource=admin"
```

```go
import "github.com/go-xuan/mongox"

func main() {
    mongox.Initialize()
    client := mongox.GetClient("default")
    db := mongox.GetDatabase("default")
    collection := db.Collection("users")
}
```

## 主要功能

- **多数据源** — 支持同时连接多个 MongoDB 实例
- **连接池管理** — MaxPoolSize / MinPoolSize / MaxIdleTime
- **配置驱动** — 配合 configx 自动从 nacos / 本地文件加载
