#include <stdio.h>

int main()
{
  unsigned  int count,i,a,b,c,n,j;
    scanf("%d",&n);
    for(j=0;j<n;j++)
    {
        scanf("%u%u%u",&a,&b,&c);
        count=0;
        for(i=a;i<b;i++)
        {
            if(i%c==0)
            {
                count++;
                break;
            }
        }
        while(1)
        {
            i=i+c;
            if(i>b)
            break;
            count++;
        }
        printf("%u\n",count);
    }
}