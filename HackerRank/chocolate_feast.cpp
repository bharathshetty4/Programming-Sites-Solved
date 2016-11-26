#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>
using namespace std;


int main() {
    int t,n,c,m,extra;
    cin>>t;
    while(t--){
        cin>>n>>c>>m;
        int answer=0;
        int mid;
        mid=n/c;
        answer=answer+mid;
       
        while(mid>=m)
        {
            
            answer=answer+(mid/m);
            
            extra=mid%m;
            mid=mid/m;
            
            mid=mid+extra;
           
        }
        
 
        cout<<answer<<endl;
    }
    return 0;
}

