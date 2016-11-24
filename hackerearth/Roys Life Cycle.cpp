#include <iostream>
using namespace std;

int main()
{
   double n;
   double i=0;
   cin>>n;
   static double num1=0,num2=0;
   for(i=0;i<n;i++){
   string input;
   cin>>input;
   	double j=0;
   	double rowcount=0,count=0;
   	while(input[j]!='\0'){

   		if(input[j]=='C'){
   		count=count+1;

   		rowcount=rowcount+1;
   			if(count>num1) num1=count;
   			if(rowcount>num2) num2=rowcount;
   		}
   		else rowcount=0;
   		j++;
   	}
   }
   cout<<num2<<" "<<num1;
}