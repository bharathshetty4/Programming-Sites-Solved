num = int(raw_input())
for x in range(0, num):
    next = 1
    pn = int(raw_input())
    for y in range(2, pn + 1):
        next = next * y
    count = 0
    while (next >= 0 and next % 10 == 0):
        count = count + 1
        next = next / 10
    print
    count
