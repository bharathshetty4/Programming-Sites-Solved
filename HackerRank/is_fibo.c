#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>
using namespace std;


int main() {
    /* Enter your code here. Read input from STDIN. Print output to STDOUT */   
  int n,i,flag;
    long double dig1,dig2,fibn,sum;
    cin>>n;
    for(i=0;i<n;i++) {
        dig1=0;
        dig2=1;
        sum=0;
        flag=0;
        cin>>fibn;
        while(sum<=fibn){
            if(sum==fibn) {
                flag=1;
            }
                sum=dig1+dig2;
                dig1=dig2;
                dig2=sum;
            }
           
    if(flag==1) cout<<"IsFibo\n";
        else cout<<"IsNotFibo\n";
    }
    
    
    return 0;
}

