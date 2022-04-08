/*
Link: https://leetcode.com/problems/kth-largest-element-in-a-stream
Status: Time Limit Exceeded
*/
package main


type KthLargest struct {
    items []int
    index int
    streamLen int
}

func Constructor(k int, nums []int) KthLargest {
    streamLen := len(nums)
    if k > streamLen{
        streamLen=k
    }
    return KthLargest{items:nums,index:k,streamLen:streamLen}
}


func (this *KthLargest) Add(val int) int {
    this.items = append(this.items,val)
    sort.Ints(this.items) // TODO: use binary search to keep the new value in right place. Sort is causing TLE
    if len(this.items)>this.streamLen{
        this.items=this.items[1:]
    }
    if len(this.items)<=0{
        return val
    }
    ind := (len(this.items)-this.index)%len(this.items)
    return this.items[ind]
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
