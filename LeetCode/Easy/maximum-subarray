/*
Link: https://leetcode.com/problems/maximum-subarray
Runtime: 167 ms, faster than 15.48% of Go online submissions for Maximum Subarray.
Memory Usage: 10 MB, less than 12.44% of Go online submissions for Maximum Subarray.
*/
package main

func maxSubArray(nums []int) int {
    if len(nums) <=0 {
        return 0
    }
    max := nums[0]
    currMax := nums[0]
    for i:=1;i<len(nums);i++ {
        if currMax > 0 {
            currMax = currMax+nums[i]
        }else{
            currMax = nums[i]
        }
        if max < currMax {
            max = currMax
        }
    }
    return max
}
