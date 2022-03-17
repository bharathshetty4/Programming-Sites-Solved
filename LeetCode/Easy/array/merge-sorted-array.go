
/*
Link: https://leetcode.com/problems/merge-sorted-array
Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Sorted Array.
Memory Usage: 2.2 MB, less than 84.71% of Go online submissions for Merge Sorted Array.
*/

func merge(nums1 []int, m int, nums2 []int, n int)  {
    idxPtr := len(nums1)-1
    for m>0||n>0{
        if n<=0||(m>0&&nums1[m-1]>nums2[n-1]){
            nums1[idxPtr]=nums1[m-1]
            m--
            idxPtr--
        }else if m<=0||(n>0&&nums1[m-1]<=nums2[n-1]){
            nums1[idxPtr]=nums2[n-1]
            n--
            idxPtr--
        }
    }
}
