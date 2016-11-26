#!/bin/python

import sys


n = int(raw_input().strip())
count=0
max=0
while(n>0):
    digit=n%2
    if(digit==1):
        count = count + 1
        if(count > max):
            max=count
    else:
        count=0
    n=n/2

print max
    
    

