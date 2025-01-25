# stk-monitor dir structure

project-root/
├── cmd/                    # 主要的应用程序入口
│   └── stock-monitor/
│       └── main.go        # 程序入口文件
│
├── internal/              # 私有应用程序和库代码
│   ├── config/           # 配置文件处理
│   │   └── config.go
│   ├── models/           # 数据模型
│   │   ├── stock.go
│   │   └── trading.go
│   ├── strategy/         # 交易策略
│   │   ├── grid.go
│   │   ├── martingale.go
│   │   └── common.go
│   └── service/          # 业务逻辑
│       └── monitor.go
│
├── pkg/                  # 可以被外部应用程序使用的库代码
│   └── utils/
│       ├── time.go
│       └── http.go
│
├── configs/              # 配置文件目录
│   └── config.yaml
│
├── api/                  # API 相关定义
│   └── api.go
│
├── test/                 # 测试文件
│   └── strategy_test.go
│
├── go.mod               # Go 模块文件
├── go.sum               # Go 模块依赖版本锁定文件
├── README.md            # 项目说明文档
└── Makefile             # 项目管理工具