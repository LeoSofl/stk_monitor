package strategy

import (
	"fmt"
	"math"
	"stk-monitor/internal/models"
)

type MartingaleStrategy struct {
	state *models.TradingState
}

func NewMartingaleStrategy() *MartingaleStrategy {
	return &MartingaleStrategy{
		state: &models.TradingState{},
	}
}

func (m *MartingaleStrategy) Check(stock *models.StockInfo) {
	basePrice := stock.InitialInvestment / 100
	// currentLevel := 0

	for i := 0; i < 5; i++ { // 最多5次加仓
		buyPrice := basePrice * (1 - float64(i+1)*0.1) // 每次下跌10%加仓
		amount := stock.InitialInvestment * math.Pow(stock.MartingaleMultiplier, float64(i))

		if stock.CurrentPrice <= buyPrice {
			fmt.Printf("马丁格尔策略提醒！等级: %d, 价格: %.2f, 建议加仓金额: %.2f\n",
				i+1, stock.CurrentPrice, amount)
		}
	}
}
