#include <stdio.h>
#include <string.h>
#include <math.h>
#include <stdlib.h>
long long int a,b,j;
int main() {
scanf("%d");
   while(scanf("%ld%ld",&a,&b)>0)
   printf("%d\n",(int)(floor(sqrt(b))-ceil(sqrt(a))+1));
    return 0;
}

  

