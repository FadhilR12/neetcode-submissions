func maxProfit(prices []int) int {
	sell := 0
	buy := 0
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[buy] {
			buy = i
		} else if prices[i]-prices[buy] > profit {
			sell = i
			profit = prices[sell] - prices[buy]
		}
	}
	if profit < 0 {
		profit = 0
	}
	return profit
}