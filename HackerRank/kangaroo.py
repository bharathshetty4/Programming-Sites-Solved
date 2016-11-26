#!/bin/python

import sys


x1,v1,x2,v2 = raw_input().strip().split(' ')
x1,v1,x2,v2 = [int(x1),int(v1),int(x2),int(v2)]
#print x1,v1,x2,v2
if(x1<x2):
    posless=x1
    speedless=v1
    posmore=x2
    speedmore=v2
else:
    posless=x2
    speedless=v2
    posmore=x1
    speedmore=v1
#print posless,posmore
flag=0
if(speedless<speedmore):
    flag=1
while(posless<posmore and flag ==0):
    posless = posless+speedless
    posmore = posmore+speedmore
if(posless==posmore):
    print "YES"
else:
    print "NO"

