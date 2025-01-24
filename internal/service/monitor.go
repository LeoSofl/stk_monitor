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
		grid:         strategy.NewGridStrategy(),
		martingale:   strategy.NewMartingaleStrategy(),
	}
}

func (m *Monitor) Start() {
	for {
		// if !utils.IsTradeTime() {
		// 	log.Println("非交易时间")
		// 	time.Sleep(time.Minute)
		// 	continue
		// }

		price, err := m.stockService.GetStockPrice(m.stock.Code)
		if err != nil {
			log.Printf("获取股票价格失败: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}
		if price > 0 {
			m.stock.CurrentPrice = price
			m.grid.Check(m.stock)
			m.martingale.Check(m.stock)
		}
		time.Sleep(5 * time.Second)
	}
}
