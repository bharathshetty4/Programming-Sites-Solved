# TODO: TLE for 10^8
from fractions import gcd
input_count = int(input().strip().split()[0])

while(input_count):
    min_lcm = 100000000
    junk_numbers = input().strip().split()
    input_digit = input().strip().split()
    input_digit = sorted(list(map(int, input_digit)))
    for i in range (0,len(input_digit)-1):
        for j in range (i+1,len(input_digit)):
            temp_lcm = (input_digit[i]*input_digit[j])/gcd(input_digit[i],input_digit[j])
            if(temp_lcm<=min_lcm):
                min_lcm=temp_lcm
    print (int(min_lcm))
    input_count = input_count - 1

