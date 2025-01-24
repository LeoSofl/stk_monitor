package main

import (
	"log"
	"stk-monitor/internal/config"
	"stk-monitor/internal/service"
)

func main() {
	// 加载配置
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 启动监控服务
	monitor := service.NewMonitor(cfg)
	monitor.Start()
}
