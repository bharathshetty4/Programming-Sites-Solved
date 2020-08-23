from Array import arr_insert,arr_printer,arr_remove,arr_search

#Initialize an array variable
arr = []

#Insert an element
print("Insert some elements to the array")
for i in range (15) :
    arr_insert.insert_one(arr,(5*i+9))
arr_printer.print_arr(arr)

#Remove the last element
print("Remove the last element of an array")
arr_remove.remove_last(arr)
arr_printer.print_arr(arr)

#1.http://www.geeksforgeeks.org/search-insert-and-delete-in-an-unsorted-array/
#Remove a key from array
print("Remove a key from an array")
delete_element = 34
arr_remove.remove_key(arr,delete_element)
arr_printer.print_arr(arr)

#Search an element
print("search for a key in an array")
search_element = 39
if(arr_search.linear_search(arr,search_element)):
    print("Found the key %s in the array\n"% search_element)
else:
    print("%s: Key not found\n" % search_element)


#2.http://www.geeksforgeeks.org/search-insert-and-delete-in-a-sorted-array/
#insert an element into sorted array
print("Insert an element to a sorted array")
insert_element = 43
arr_insert.sort_insert(arr,insert_element)
arr_printer.print_arr(arr)


#serach an element from sorted array
print("[Binary] search an element from a sorted array")
search_element = 54
n = len(arr) - 1
pos = arr_search.binary_search(arr,0,int(n),search_element)
if (pos == -1 ):
    print("%s: Key not found\n" % search_element)
else:
    print("Found the key %s in the array in position \n"% search_element)

#delete an element from sorted array

print("Delete a given key from an array")
delete_element = 59
arr_remove.sort_remove(arr,delete_element)
arr_printer.print_arr(arr)