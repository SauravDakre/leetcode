package BestTimeToBuySellStock2

//  simply go on crawling over the slope and keep on adding the profit obtained from every consecutive transaction. In the end,we will be using the peaks and valleys effectively,
func maxProfit(prices []int) int {
	s := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			s += (prices[i] - prices[i-1])
		}
	}
	return s
}
