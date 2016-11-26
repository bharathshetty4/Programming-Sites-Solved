#!/bin/python3

import sys


t = int(input().strip())
for a0 in range(t):
    maxnum=0
    n,k = input().strip().split(' ')
    n,k = [int(n),int(k)]
    print(k-1 if ((k-1) | k) <= n else k-2)
            

