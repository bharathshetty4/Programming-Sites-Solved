#!/bin/python

import sys

def passhere(a):
    alen=len(a)
    count=0
    for i in range(alen):
        for j in range(alen-1):
            if(a[j]>a[j+1]):
                count = count + 1
                temp=a[j]
                a[j]=a[j+1]
                a[j+1]=temp
    print "Array is sorted in",count, "swaps."
    print "First Element:", a[0]
    print "Last Element:", a[len(a)-1]
    
n = int(raw_input().strip())
a = map(int,raw_input().strip().split(' '))
passhere(a)
