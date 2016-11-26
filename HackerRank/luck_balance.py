# Enter your code here. Read input from STDIN. Print output to STDOUT
new_dict = {}
inp = raw_input().split()
size = int(inp[0])
caseshouldpass = int(inp[1])
implist = []
notimplist = []

while(size):
    point,imp = raw_input().strip().split() 
    if(int(imp)==1):
        implist.insert(len(implist),int(point))
    else:
        notimplist.insert(len(notimplist),int(point))
    size = size - 1
    
casecanfail = len(implist)-caseshouldpass

implist.sort()

notimplist.sort()

implist.extend(notimplist)
i=0
totalluck = 0 
while(casecanfail>0):
    totalluck = totalluck - int(implist[0])
    implist.pop(0)
    casecanfail = casecanfail-1
for i in implist:
    totalluck = totalluck + int(i) 
print totalluck

