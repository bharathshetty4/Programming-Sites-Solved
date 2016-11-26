#!/bin/python3

import sys

one =input().strip().split(' ')
two =input().strip().split(' ')
count = 0 
count1 = 0

len1 = len(one)
while(len1):
    if(int(one[len1-1])>int(two[len1-1])):
        count = count + 1
    elif(int(one[len1-1])<int(two[len1-1])):
        count1= count1 + 1 
    len1= len1 - 1

print (count,count1)
