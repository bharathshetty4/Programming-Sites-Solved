package main

/*
URL: https://leetcode.com/problems/maximum-units-on-a-truck/
Status: Success
Runtime: 32 ms, faster than 49.07% of Go online submissions for Maximum Units on a Truck.
Memory Usage: 7 MB, less than 8.33% of Go online submissions for Maximum Units on a Truck.
*/
import (
	"fmt"
)

func maximumUnits(boxTypes [][]int, truckSize int) int {
    var unitNums []int
    boxMap := make(map[int]int)
    for _, box := range boxTypes{
        // first index is number of boxes
        // second index is size of thar unit
        if boxMap[box[1]] == 0 {
            // get the unique qty numbers
            unitNums = append(unitNums,box[1])
        }
        // get the total number boxes available for the given qty
        boxMap[box[1]]+= box[0]
    }
    sort.Ints(unitNums)
    numsLen := len(unitNums)
    numsUnit := 0 
    for i:=numsLen-1;i>=0;i--{
        if truckSize <=0 {
            return numsUnit
        }
        if boxMap[unitNums[i]] < truckSize {
            numsUnit+=(boxMap[unitNums[i]]*unitNums[i])
            truckSize-=boxMap[unitNums[i]]
        }else{
            numsUnit+=(unitNums[i]*truckSize)
            truckSize-=boxMap[unitNums[i]]
        }
    }
    return numsUnit
}
