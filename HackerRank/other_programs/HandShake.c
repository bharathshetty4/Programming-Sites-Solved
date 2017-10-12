#include <stdio.h>
#include <string.h>
#include <math.h>
#include <stdlib.h>

int main() {
unsigned int a,b,i,shake;
scanf("%d",&a);
for(i=0;i<a;i++){
  scanf("%d",&b); 
  shake=(b*(b-1))/2;
   b<2?printf("0\n"):printf("%d\n",shake);        
    }
return 0;
}
