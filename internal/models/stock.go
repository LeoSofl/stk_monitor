package models

type StockInfo struct {
	Code                 string
	CurrentPrice         float64
	InitialInvestment    float64
	GridSize             float64
	GridLevels           int
	MartingaleMultiplier float64
	StopLoss             float64
	TakeProfit           float64
}
