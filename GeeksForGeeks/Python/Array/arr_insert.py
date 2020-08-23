def insert_one(arr , val ):
    arr.append(val)

def sort_insert(arr, val):
    #append a calue to the last, which will append the new value
    arr.append(val)
    for i in range(len(arr)-2,0,-1):
        if(arr[i] < val):
            break
        arr[i+1] = arr[i]

    arr[i+1] = val
