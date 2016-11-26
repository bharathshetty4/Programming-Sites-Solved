s=input();
for i in s:
    if(i.isupper()):
        print (i.lower(),end="")
    elif(i.islower()):
        print (i.upper(),end="")
    else :
        print (i,end="")
