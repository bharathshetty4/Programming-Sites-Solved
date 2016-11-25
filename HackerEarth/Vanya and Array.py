# TODO : Passing only one test case :(
line1 = input().strip().split(" ")
line2 = input().strip().split(" ")
lenline1 = int(line1[0])
funcline = line2

for i in range(lenline1):
    count = 0
    for j in range(i + 1, lenline1):
        if (int(line2[i]) < int(line2[j])):
            count = count + 1
    funcline[i] = count

newcount = 0
for i in range(lenline1):
    for j in range(i + 1, lenline1):
        if (funcline[i] + funcline[j] >= 10):
            newcount = newcount + 1

print(newcount)
