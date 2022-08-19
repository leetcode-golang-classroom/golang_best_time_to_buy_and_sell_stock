package sol

func maxProfit(prices []int) int {
	buy, max_profit := 0, 0
	n := len(prices)
	var max = func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	for sell := 1; sell < n; sell++ {
		if prices[sell] > prices[buy] {
			max_profit = max(max_profit, prices[sell]-prices[buy])
		} else { // prices[sell] is current smallest
			buy = sell
		}
	}
	return max_profit
}
