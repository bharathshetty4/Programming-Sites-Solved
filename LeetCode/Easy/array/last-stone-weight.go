/*
Link: https://leetcode.com/problems/last-stone-weight
Runtime: 0 ms, faster than 100.00% of Go online submissions for Last Stone Weight.
Memory Usage: 2 MB, less than 82.05% of Go online submissions for Last Stone Weight.
*/
package main

func lastStoneWeight(stones []int) int {
    if len(stones) == 0{
        return 0
    }
    for len(stones) > 1{
        sort.Ints(stones)
        l := stones[len(stones)-1]
        l2 := stones[len(stones)-2]
        stones[len(stones)-2]=l-l2
        stones = stones[:len(stones)-1]
    }
    return stones[0]
}
