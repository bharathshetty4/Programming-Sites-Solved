#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>
using namespace std;


int main() {
    int n;
    cin>>n;
    int m,c;
    int t;
    while(n--){
       cin>>m;
        cin>>c;
        while(m--){
            cin>>t;
            if(t<=0) c--;
        }
        if(c<=0){
            cout<<"NO"<<"\n";
        }else cout<<"YES\n";
        
        
    }
    return 0;
}

