package config

import (
	"os"
	"stk-monitor/internal/models"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Stock struct {
		Code                 string  `yaml:"code"`
		InitialInvestment    float64 `yaml:"initial_investment"`
		GridSize             float64 `yaml:"grid_size"`
		GridLevels           int     `yaml:"grid_levels"`
		MartingaleMultiplier float64 `yaml:"martingale_multiplier"`
		StopLoss             float64 `yaml:"stop_loss"`
		TakeProfit           float64 `yaml:"take_profit"`
	} `yaml:"stock"`
}

// 加载配置文件
func Load() (*Config, error) {
	cfg := &Config{}

	// 读取配置文件
	file, err := os.ReadFile("configs/config.yaml")
	if err != nil {
		return nil, err
	}

	// 解析YAML
	err = yaml.Unmarshal(file, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// 转换为StockInfo
func (c *Config) ToStockInfo() *models.StockInfo {
	return &models.StockInfo{
		Code:                 c.Stock.Code,
		InitialInvestment:    c.Stock.InitialInvestment,
		GridSize:             c.Stock.GridSize,
		GridLevels:           c.Stock.GridLevels,
		MartingaleMultiplier: c.Stock.MartingaleMultiplier,
		StopLoss:             c.Stock.StopLoss,
		TakeProfit:           c.Stock.TakeProfit,
	}
}
