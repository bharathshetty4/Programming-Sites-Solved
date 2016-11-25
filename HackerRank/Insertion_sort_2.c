#include <stdio.h>
#include <string.h>
#include <math.h>
#include <stdlib.h>
#include <assert.h>
int main(void) {
   
   int ar_size;
scanf("%d", &ar_size);
int ar[ar_size], _ar_i;
for(_ar_i = 0; _ar_i < ar_size; _ar_i++) { 
   ar[_ar_i]=0; 
	}
for(_ar_i = 0; _ar_i < ar_size; _ar_i++) { 
   scanf("%d", &ar[_ar_i]); 
	}

int n=ar_size,i,j;
int temp=0,temp1;
for(i=1;i<n;i++){
        
temp=i;        
while(ar[temp]<ar[temp-1]){     
     temp1=ar[temp];
     ar[temp]=ar[temp-1];
     ar[temp-1]=temp1;
     temp--;
     if(temp<0) break;         
     }
        
for(j=0;j<n;j++){
    printf("%d ",ar[j]);
    }
printf("\n");
    }
    
return 0;
}
