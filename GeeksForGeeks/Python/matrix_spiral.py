#https://www.geeksforgeeks.org/print-a-given-matrix-in-spiral-form/

#row_cloumn = input().strip().split(" ")

#row = int(row_cloumn[0])
#column = int(row_cloumn[1])

#print(row, column)

def print_one_spiral (matrix, count, size_row, size_col):
    print(count,matrix)
    #print top row


    #print right

    #print left


def print_spiral (matrix, size_row, size_col):
    if size_row % 2 == 1 :
        spiral_needed = (size_row/2) +1
    else:
        spiral_needed = (size_row / 2)

    for i in range(int(spiral_needed)):
        print_one_spiral(matrix, i, size_row, size_col)

matrix = [[1,2,3],
          [4,5,6],
          [1,2,3],
          [4,5,6]
          ]
print_spiral(matrix,4,3)

