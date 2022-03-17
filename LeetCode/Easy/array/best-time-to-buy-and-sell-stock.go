package main

/*
URL: https://leetcode.com/problems/best-time-to-buy-and-sell-stock
Runtime: 278 ms, faster than 9.83% of Go online submissions for Best Time to Buy and Sell Stock.
Memory Usage: 8.7 MB, less than 64.30% of Go online submissions for Best Time to Buy and Sell Stock.
*/

func maxProfit(prices []int) int {
    lPrices := len(prices)
    if lPrices <= 0{
        return 0
    }
    maxToSell := prices[lPrices-1]
    maxProfit := 0
    
    for i:=lPrices-2;i>=0;i--{
        if maxToSell-prices[i]>maxProfit{
            maxProfit = maxToSell-prices[i]
        }
        if prices[i] > maxToSell{
            maxToSell=prices[i]
        }
    }
    return maxProfit
}
