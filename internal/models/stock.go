package models

type StockInfo struct {
	Code         string
	CurrentPrice float64

	BasePrice         float64
	InitialInvestment float64

	MartingaleMultiplier float64
	MartingaleGridSize   float64
	MartingaleLevels     int
	MartingaleStopLoss   float64

	GridUpperPrice      float64
	GridLowerPrice      float64
	GridSize            float64
	GridCount           int
	GridAmount          float64
	GridStopLossPrice   float64
	GridStopProfitPrice float64
}
