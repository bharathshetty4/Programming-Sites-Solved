#include <stdio.h>
#include <string.h>
#include <math.h>
#include <stdlib.h>
#include <assert.h>
int lonelyinteger(int a_size, int* a) {
int i,j,k,flag;
for(i=0;i<a_size;i++){
    flag=1;
    for(j=i+1;j<a_size;j++){
        if(a[i]==a[j]) {
            flag=0;
            for(k=j;k<a_size;k++) a[k]=a[k+1];
               a_size--;
        }
        }
    if(flag) printf("%d ",a[i]);
}
return 0;

}
int main() {
    int res;
    int _a_size, _a_i;
    scanf("%d", &_a_size);
    int _a[_a_size];
    for(_a_i = 0; _a_i < _a_size; _a_i++) {
        int _a_item;
        scanf("%d", &_a_item);

        _a[_a_i] = _a_item;
    }

    res = lonelyinteger(_a_size, _a);
    return 0;
}