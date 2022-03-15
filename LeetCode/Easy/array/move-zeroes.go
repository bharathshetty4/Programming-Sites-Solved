package main

/*
URL: https://leetcode.com/problems/move-zeroes
Status: Success
Runtime: 20 ms, faster than 92.93% of Go online submissions for Move Zeroes.
Memory Usage: 6.6 MB, less than 86.77% of Go online submissions for Move Zeroes.
*/

func moveZeroes(nums []int)  {
    if len(nums) <=1{
        return
    }
    cnt := 0
    for idx:=0;idx<len(nums);idx++{
        if 0 == nums[idx]{
            continue
        }
        nums[cnt] = nums[idx]
        cnt++
    }
    for idx:=cnt;idx<len(nums);idx++{
        nums[idx]=0
    }
}
