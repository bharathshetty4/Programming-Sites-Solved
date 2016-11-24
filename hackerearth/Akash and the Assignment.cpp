#include <iostream>
using namespace std;

int main()
{
    int m,n,i;
    int f,s,d;
    cin>>m;
    cin>>n;
    char a[100000];
    for(i=0;i<m;i++)
    {
    cin>>a[i];
    }

    for(i=0;i<n;i++){
    cin>>f;
    cin>>s;
    cin>>d;
    int num=f-2+d;
    cout<<a[num]<<"\n";
    }
    return 0;
}