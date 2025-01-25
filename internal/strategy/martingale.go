package strategy

import (
	"log"
	"math"
	"stk-monitor/internal/config"
	"stk-monitor/internal/models"
)

type MartingaleStrategy struct {
	gridPrices []float64 // 存储固定的马丁格尔买入价格点
	amounts    []float64 // 存储对应的买入金额
}

func NewMartingaleStrategy(cfg *config.Config) *MartingaleStrategy {
	// 计算所有价格点和对应的买入金额

	gridPrices := make([]float64, cfg.Stock.Martingale.Levels)
	amounts := make([]float64, cfg.Stock.Martingale.Levels)

	for i := 0; i < cfg.Stock.Martingale.Levels; i++ {
		// 计算每个级别的价格点
		gridPrices[i] = cfg.Stock.BasePrice * (1 - cfg.Stock.Martingale.GridSize*float64(i+1))
		// 计算每个级别的买入金额
		amounts[i] = cfg.Stock.InitialInvestment * math.Pow(cfg.Stock.Martingale.Multiplier, float64(i))
	}

	return &MartingaleStrategy{
		gridPrices: gridPrices,
		amounts:    amounts,
	}
}

func (m *MartingaleStrategy) Check(stock *models.StockInfo) {
	currentPrice := stock.CurrentPrice
	stopLossPrice := stock.BasePrice * (1 - stock.MartingaleStopLoss)

	// 检查是否触发止损
	if currentPrice <= stopLossPrice {
		log.Printf("警告！当前价格 %.2f 已触发止损价格 %.2f，建议清仓止损！\n",
			currentPrice, stopLossPrice)
		return
	}

	// 从高级别往低级别检查，找到第一个符合条件的级别
	for i := len(m.gridPrices) - 1; i >= 0; i-- {
		targetPrice := m.gridPrices[i]

		// 当前价格低于目标价格，且在1%误差范围内
		if currentPrice <= targetPrice &&
			currentPrice > targetPrice*(1-0.01) {

			// 计算累计跌幅
			totalDrop := (stock.BasePrice - currentPrice) / stock.BasePrice * 100

			log.Printf("马丁格尔策略提醒！\n")
			log.Printf("当前等级: %d, 累计跌幅: %.1f%%\n", i+1, totalDrop)
			log.Printf("当前价格: %.2f, 目标价格: %.2f\n", currentPrice, targetPrice)
			log.Printf("建议买入金额: %.2f\n", m.amounts[i])

			// 给出风险提示
			if i >= len(m.gridPrices)/2 {
				log.Printf("风险提示：已触发第%d档，请注意仓位控制！\n", i+1)
			}
			return
		}

		// 止盈
		if currentPrice >= stock.BasePrice*(1+stock.MartingaleStopLoss) {
			log.Printf("止盈提醒！当前价格 %.2f 已达到止盈价格 %.2f，建议清仓止盈！\n",
				currentPrice, stock.BasePrice*(1+stock.MartingaleStopLoss))
			return
		}
	}
}
