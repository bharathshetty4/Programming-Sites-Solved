#include <iostream>
using namespace std;

int main()
{
   int n,row,col;
   cin>>n;
   int i,j,k;
   int a[1000][1000];

   for(int i=0;i<n;i++){
   	cin>>row;

int sum=0;
   	//input to array
   for(j=0;j<=row-1;j++){
   	for(k=0;k<=row-1;k++)
   	{

   	  	a[j][k]=abs(j-k);

   	  }
   }
   for(j=0;j<row;j++){
   	for(k=0;k<row;k++)
   	{

   	  sum+=a[j][k];

   	  }

   }

   	cout<<sum<<"\n";


   }
    return 0;
}