# golang_best_time_to_buy_and_sell_stock

You are given an array `prices` where `prices[i]` is the price of a given stock on the `ith` day.

You want to maximize your profit by choosing a **single day** to buy one stock and choosing a **different day in the future** to sell that stock.

Return *the maximum profit you can achieve from this transaction*. If you cannot achieve any profit, return `0`.

## Examples

**Example 1:**

```
Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

```

**Example 2:**

```
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.

```

**Constraints:**

- `1 <= prices.length <= 105`
- `0 <= prices[i] <= 104`

## 解析

給定一個正整數陣列 prices 

每個 prices[i] 代表第 i 日股票的價錢

假設必須要選某一天 假設是 j 買入 還有另一天 k 賣出 且 j < k

則獲利 = height[k] - height[j]

要求寫出一個演算法算出最大的獲利

已知如果要找出最大獲利

就必須在每個 i  找到一個 k 使得 k > i 且 height[k] > height[i] 讓 height[k] 最大化

![](https://i.imgur.com/BimdqXp.png)

![](https://i.imgur.com/sWZooAr.png)



這時可以透過 slide window 的技巧計算每個 slide window 內的最大值

逐步比較出整體的最大值

因為知道 買入日比需要在賣出日之前

所以初始化 buy_min = 0, sell_max = 1, max_profit  = 0

sell_max 從 1.. len(prices) - 1 做以下運算

如果 prices[sell_max] > prices[buy_min] 則更新 max_profit = max(max_profit, prices[sell_max] - prices[buy_min])

如果 prices[sell_max] ≤ prices[buy_min] 代表 當下 sell_max 可能成為下個最小 buy_min 所以更新 buy_min = sell_max

當所有值比完 回傳 max_profit


![](https://i.imgur.com/lIthXp7.png)
  

## 程式碼
```go
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
```
## 困難點

1. 要看出 profit 的buy 與 sell 的關係
2. 要看出 max profit 的形成條件

## Solve Point

- [x]  初始化 buy =0, sell = 1, maxProfit = 0
- [x]  從 sell = 1..len(prices) 做以下運算
- [x]  當 prices[sell] > prices[buy]  時，更新 maxProfit = max(maxProfit, prices[sell] - prices[buy])
- [x]  當 prices[sell] ≤ prices[buy] 時， 當下 prices[sell] 是最小值 。所以更新 buy = sell
- [x]  當比完所有 sell 時，回傳 maxProfit