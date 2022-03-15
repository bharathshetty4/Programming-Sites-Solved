package main

/*
URL: https://leetcode.com/problems/sort-array-by-parity
Status: Success
Runtime: 5 ms, faster than 88.56% of Go online submissions for Sort Array By Parity.
Memory Usage: 5.3 MB, less than 12.94% of Go online submissions for Sort Array By Parit
*/

func sortArrayByParity(nums []int) []int {
    if len(nums) <=1{
        return nums
    }
    cnt := 0
    var oddArray []int
     for idx:=0;idx<len(nums);idx++{
        if nums[idx] %2 !=0 {
            oddArray=append(oddArray,nums[idx])
            continue
        }
        nums[cnt] = nums[idx]
        cnt++
    }
    return append(nums[:cnt],oddArray...)
}
