#include <stdio.h>
#include <string.h>
#include <math.h>
#include <stdlib.h>
#include <assert.h>
void insertionSort(int ar_size, int *  ar) {
int n=ar_size,i,j;
    int temp=ar[n-1];
    for(i=n-1;i>=0;i--){
        if(temp>ar[i-1]){
            ar[i]=temp;
            break;
        }else ar[i]=ar[i-1];
        
        for(j=0;j<n;j++){
            printf("%d ",ar[j]);
        }
        printf("\n");
    }
    for(j=0;j<n;j++){
            printf("%d ",ar[j]);
    }
}
int main(void) {
   
   int _ar_size;
scanf("%d", &_ar_size);
int _ar[_ar_size], _ar_i;
for(_ar_i = 0; _ar_i < _ar_size; _ar_i++) { 
   scanf("%d", &_ar[_ar_i]); 
}

insertionSort(_ar_size, _ar);  
   return 0;
}
