# Enter your code here. Read input from STDIN. Print output to STDOUT

def factorial(n):
    if(n<=1):
        return 1
    else:
        num=n*factorial(n-1)
        return num
    
n=int(raw_input())
factn=factorial(n)
print factn
    
