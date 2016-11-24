in1 = (raw_input().strip().split())
in2 = (raw_input().strip().split())
in3 = (raw_input().strip().split())

n = int(in1[0])
m = int(in1[1])
c = []
for i in in2:
    c.append(int(i))

num = 1

for i in in3:
    flag = 0
    while (flag == 0):
        if (num % c[int(i) - 1] == 0):
            flag = 1
            num = num + 1
        else:
            num = num + 1
            flag = 0
        for val in c:
            if (val != c[int(i) - 1]):
                if (num % val == 0):
                    flag = 0

print
num - 1
