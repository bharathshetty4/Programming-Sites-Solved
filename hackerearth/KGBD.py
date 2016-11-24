def kc_gcd(a, b):
    n1 = a
    n2 = b
    while (a > 0 and b > 0):
        a = a - b
        a, b = b, a
    while (n1 != n2):
        if (n1 > n2):
            n1 = n1 - n2
        else:
            n2 = n2 - n1
    return (n1 == a + b)


n = int(input().strip())
while (n >= 1):
    count = 0
    data = int(input().strip())
    for i in range(1, data + 1):
        for j in range(1, data + 1):
            if (kc_gcd(i, j)):
                count = count + 1
    n = n - 1
    print(count)
