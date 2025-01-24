package strategy

import (
	"fmt"
	"stk-monitor/internal/models"
)

type GridStrategy struct {
	state *models.TradingState
}

func NewGridStrategy() *GridStrategy {
	return &GridStrategy{
		state: &models.TradingState{},
	}
}

func (g *GridStrategy) Check(stock *models.StockInfo) {
	basePrice := stock.InitialInvestment / 100
	for i := 0; i < stock.GridLevels; i++ {
		buyPrice := basePrice * (1 - float64(i+1)*stock.GridSize)
		sellPrice := basePrice * (1 + float64(i+1)*stock.GridSize)

		if stock.CurrentPrice <= buyPrice {
			fmt.Printf("网格策略买入提醒！等级: %d, 价格: %.2f\n", i+1, stock.CurrentPrice)
		}
		if stock.CurrentPrice >= sellPrice {
			fmt.Printf("网格策略卖出提醒！等级: %d, 价格: %.2f\n", i+1, stock.CurrentPrice)
		}
	}
}
