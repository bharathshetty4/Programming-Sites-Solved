def array_left_rotation(a, n, k):
    #use .copy to avoid copy by reference
    b = a.copy()  # or b = a[:]
    for i in range(0,len(a)):
        b [i] = a[(i+k)%n]
    return b

n, k = map(int, input().strip().split(' '))
a = list(map(int, input().strip().split(' ')))
answer = array_left_rotation(a, n, k);
print(*answer, sep=' ')

