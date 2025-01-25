package strategy

import (
	"log"
	"math"
	"sort"
	"stk-monitor/internal/config"
	"stk-monitor/internal/models"
)

type GridStrategy struct {
	gridPrices []float64 // 存储固定的网格价格点
}

func NewGridStrategy(cfg *config.Config) *GridStrategy {
	// 计算固定的网格价格点
	upperPrice := cfg.Stock.Grid.UpperPrice
	lowerPrice := cfg.Stock.Grid.LowerPrice
	basePrice := cfg.Stock.BasePrice

	// prefer gridSize
	gridSize := cfg.Stock.Grid.GridSize

	// 从基准价格开始，向上和向下分别计算网格
	upGridCount := int((upperPrice - basePrice) / (basePrice * gridSize))
	downGridCount := int((basePrice - lowerPrice) / (basePrice * gridSize))
	totalGridCount := upGridCount + downGridCount

	log.Printf("网格策略初始化:")
	log.Printf("基准价格: %.2f", basePrice)
	log.Printf("网格大小: %.1f%%", gridSize*100)
	log.Printf("上边界: %.2f (+%.1f%%)", upperPrice, (upperPrice-basePrice)/basePrice*100)
	log.Printf("下边界: %.2f (-%.1f%%)", lowerPrice, (basePrice-lowerPrice)/basePrice*100)
	log.Printf("向上网格数: %d", upGridCount)
	log.Printf("向下网格数: %d", downGridCount)

	// 计算所有网格价格点
	gridPrices := make([]float64, totalGridCount+1)
	// 从基准价格向下计算网格
	for i := 0; i <= downGridCount; i++ {
		price := basePrice * (1 - float64(i)*gridSize)
		gridPrices[i] = price
		log.Printf("下方网格%d: %.2f (-%.1f%%)",
			i, price, (basePrice-price)/basePrice*100)
	}

	// 从基准价格向上计算网格
	for i := 1; i <= upGridCount; i++ {
		price := basePrice * (1 + float64(i)*gridSize)
		gridPrices[downGridCount+i] = price
		log.Printf("上方网格%d: %.2f (+%.1f%%)",
			i, price, (price-basePrice)/basePrice*100)
	}

	// 排序确保价格从低到高
	sort.Float64s(gridPrices)

	log.Printf("gridPrices: %v", gridPrices)

	return &GridStrategy{
		gridPrices: gridPrices,
	}
}

func (g *GridStrategy) Check(stock *models.StockInfo) {
	currentPrice := stock.CurrentPrice
	basePrice := stock.BasePrice

	// 检查是否触及止盈止损
	if currentPrice >= stock.GridStopProfitPrice {
		log.Printf("触及止盈价格 %.2f，建议大比例卖出！", stock.GridStopProfitPrice)
		return
	}
	if currentPrice <= stock.GridStopLossPrice {
		log.Printf("触及止损价格 %.2f，建议止损！", stock.GridStopLossPrice)
		return
	}

	// 检查是否触及任何网格线
	for i, gridPrice := range g.gridPrices {
		// 允许0.1%的误差范围
		if math.Abs(currentPrice-gridPrice) < gridPrice*0.001 {
			log.Printf("网格策略提醒！触及第%d条网格线\n", i+1)
			log.Printf("当前价格: %.2f, 网格价格: %.2f\n", currentPrice, gridPrice)
			// 计算相对于基准价格的偏离百分比
			deviation := (currentPrice - basePrice) / basePrice

			if currentPrice < basePrice {
				// 买入区域
				shares := stock.GridAmount
				// 确保是100的整数倍
				shares = (shares / 100) * 100
				log.Printf("建议买入股数: %d (约%.2f万元)\n",
					shares, float64(shares)*currentPrice/10000)

			} else {
				// 卖出区域，根据偏离度调整卖出比例
				sellRatio := getSellRatio(deviation)
				shares := int(float64(stock.GridAmount) * sellRatio)
				// 确保是100的整数倍
				shares = (shares / 100) * 100
				log.Printf("建议卖出股数: %d (卖出比例: %.1f%%, 约%.2f万元)\n",
					shares, sellRatio*100, float64(shares)*currentPrice/10000)
			}
			return
		}
	}
}

// 根据偏离度计算卖出比例
func getSellRatio(deviation float64) float64 {
	// 基础卖出比例为1.0（与买入量相同）
	baseRatio := 1.0

	// 根据偏离度调整卖出比例
	if deviation <= 0.015 { // 1.5%以内
		return baseRatio * 0.8 // 卖出80%
	} else if deviation <= 0.03 { // 1.5%-3%
		return baseRatio * 1.0 // 卖出100%
	} else if deviation <= 0.045 { // 3%-4.5%
		return baseRatio * 1.2 // 卖出120%
	} else if deviation <= 0.06 { // 4.5%-6%
		return baseRatio * 1.5 // 卖出150%
	} else { // 6%以上
		return baseRatio * 2.0 // 卖出200%
	}
}
