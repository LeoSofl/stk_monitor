package utils

import "time"

func IsTradeTime() bool {
	now := time.Now()

	// 如果是周末，不交易
	if now.Weekday() == time.Saturday || now.Weekday() == time.Sunday {
		return false
	}

	// 获取当前时间的小时和分钟
	hour := now.Hour()
	minute := now.Minute()
	currentTime := hour*100 + minute // 转换为类似 930、1500 的格式

	// 上午交易时段：9:30 - 11:30
	// 下午交易时段：13:00 - 15:00
	return (currentTime >= 930 && currentTime <= 1130) ||
		(currentTime >= 1300 && currentTime <= 1500)
}
