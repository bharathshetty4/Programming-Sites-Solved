#include <stdio.h>
#include <string.h>
#include <math.h>
#include <stdlib.h>

int main() {
int n,c,d,i;
char str[10000];
char str1[10000]="";
int size;
scanf("%d",&size);
while(size--){
   scanf("%s",str);
   n = strlen(str);
 // copy string
    for (c = n - 1, d = 0; c >= 0; c--, d++)
      str1[d] = str[c];
      str1[d] = '\0';
    int flag=1;
    for(i=1;i<n;i++){
       int one=abs(str[i]-str[i-1]);
       int two=abs(str1[i]-str1[i-1]);
    if(one!=two){
        flag=0;
    }
}
        if(flag==1){
            printf("Funny\n");
        }else{
            printf("Not Funny\n");
        }
}
    return 0;
}

