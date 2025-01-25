package config

import (
	"os"
	"stk-monitor/internal/models"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Stock struct {
		Code              string  `yaml:"code"`
		InitialInvestment float64 `yaml:"initial_investment"`
		BasePrice         float64 `yaml:"base_price"`

		Martingale struct {
			Multiplier float64 `yaml:"multiplier"`
			GridSize   float64 `yaml:"grid_size"`
			Levels     int     `yaml:"levels"`
			StopLoss   float64 `yaml:"stop_loss"`
		} `yaml:"martingale"`

		Grid struct {
			UpperPrice      float64 `yaml:"upper_price"`
			LowerPrice      float64 `yaml:"lower_price"`
			GridCount       int     `yaml:"grid_count"`
			GridSize        float64 `yaml:"grid_size"`
			AmountPerGrid   float64 `yaml:"amount_per_grid"`
			StopLossPrice   float64 `yaml:"stop_loss_price"`
			StopProfitPrice float64 `yaml:"stop_profit_price"`
		} `yaml:"grid"`
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
		Code:              c.Stock.Code,
		InitialInvestment: c.Stock.InitialInvestment,
		BasePrice:         c.Stock.BasePrice,

		GridUpperPrice:      c.Stock.Grid.UpperPrice,
		GridLowerPrice:      c.Stock.Grid.LowerPrice,
		GridSize:            c.Stock.Grid.GridSize,
		GridCount:           c.Stock.Grid.GridCount,
		GridAmount:          c.Stock.Grid.AmountPerGrid,
		GridStopLossPrice:   c.Stock.Grid.StopLossPrice,
		GridStopProfitPrice: c.Stock.Grid.StopProfitPrice,

		MartingaleMultiplier: c.Stock.Martingale.Multiplier,
		MartingaleGridSize:   c.Stock.Martingale.GridSize,
		MartingaleLevels:     c.Stock.Martingale.Levels,
		MartingaleStopLoss:   c.Stock.Martingale.StopLoss,
	}
}
