#!/usr/bin/python
# TODO: TLE for 10^5

input_count = int(input())
while (input_count):
    input_data = input().strip().split()
    n = int(input_data[0])
    m = int(input_data[1])
    c = int(input_data[2])
    count = 0
    for i in range(1, m + 1):
        if (i <= c):
            for j in range(1, n + 1):
                if (j <= c):
                    if (i * j == c):
                        count = count + 1
    print(count)
    input_count = input_count - 1
