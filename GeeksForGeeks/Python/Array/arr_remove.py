from Array import arr_search

def remove_last(arr):
    try:
        elem = arr.pop()
        print("Successfully deleted the last element %s"% elem)
    except Exception as e:
        print("Error deleting the last element of an array")

def remove_key(arr,key):
    try:
        arr.remove(key)
        print("Successfully deleted the key %s"% key)
    except Exception as e:
        print("Key %s not found in array to delete"% key)

def sort_remove(arr,key):
    pos = arr_search.binary_search(arr,0,len(arr)-1,key)
    if (pos == -1 ):
        print("Key %s is not found to delete\n" % key)
    else:
        for i in range (pos,len(arr)-1):
            arr[i]=arr[i+1]