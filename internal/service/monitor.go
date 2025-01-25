package service

import (
	"log"
	"stk-monitor/internal/config"
	"stk-monitor/internal/models"
	"stk-monitor/internal/service/stock"
	"stk-monitor/internal/strategy"
	"time"
)

type Monitor struct {
	stock        *models.StockInfo
	grid         *strategy.GridStrategy
	martingale   *strategy.MartingaleStrategy
	stockService *stock.StockService
}

func NewMonitor(cfg *config.Config) *Monitor {
	return &Monitor{
		stock:        cfg.ToStockInfo(),
		stockService: stock.NewStockService(),
		grid:         strategy.NewGridStrategy(cfg),
		martingale:   strategy.NewMartingaleStrategy(cfg),
	}
}

func (m *Monitor) Start() {
	for {
		price, err := m.stockService.GetStockPrice(m.stock.Code)
		if err != nil || price == 0 {
			log.Printf("获取股票价格失败: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		// update stock info
		m.stock.CurrentPrice = price
		log.Printf("code: %v, currentPrice: %v", m.stock.Code, m.stock.CurrentPrice)

		// check strategies
		m.martingale.Check(m.stock)
		m.grid.Check(m.stock)

		time.Sleep(5 * time.Second)
	}
}
