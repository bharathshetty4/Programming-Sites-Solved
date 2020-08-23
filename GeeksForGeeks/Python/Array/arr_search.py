def linear_search(arr,element):
    for i in arr:
        if(i == element):
            return 1
    return 0

def binary_search(arr, low, high, element):
    if(high<low):
        return -1
    mid = (low+high) // 2
    if (arr[mid] == element):
        return mid
    if (arr[mid]<element):
        return binary_search(arr,mid+1,high,element)
    return binary_search(arr,low,mid-1,element)