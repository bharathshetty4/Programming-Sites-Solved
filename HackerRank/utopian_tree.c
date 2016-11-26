#include <iostream>
using namespace std;

int height(int n) {
    return 0;
}
int main() {
    int T;
    cin >> T;
    while (T--) {
        int n;
        long int count=1;
        cin >> n;
        for(int i=1;i<=n;i++){
            if(i%2!=0) count=count*2;
                else count++;
        }
        cout<<count<<"\n";
    }
}

