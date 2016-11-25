#include <iostream>
using namespace std;

int main()
{
    int m,n,i,j;
    cin>>m;
    cin>>n;
    int mat[m][n];
    for(i=0;i<m;i++){
    	for(j=0;j<n;j++){
    		cin>>mat[i][j];
    	}
    }

    int cas;
    cin>>cas;
    int k;
    for(k=0;k<cas;k++){

    int ser;
    cin>>ser;
    int flag=0;
    for(i=0;i<m;i++){
    	for(j=0;j<n;j++)
    	{
    		if(ser==mat[i][j])
    		{
    			flag=1;
    			goto res;
    		}
    	}
    }
    res:
    if(flag==0) cout<<"-1 -1\n";
    else cout<<i<<" "<<j<<"\n";

    }
    return 0;
}