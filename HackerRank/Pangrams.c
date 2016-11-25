#include <stdio.h>
#include <string.h>
#include <math.h>
#include <stdlib.h>

int main() {
char in_string[4000];
    int present[27],i;
    int cur_str;
    char cur_string;
    for(i=1;i<=26;i++) present[i]=0;
    gets(in_string);
    int in_len=strlen(in_string);
    for(i=0;i<in_len;i++){
        
       cur_string=tolower(in_string[i]);
        
        if(isalpha(cur_string)){
            cur_str=cur_string-'a'+1;
            
            present[cur_str]=1;
                   }
    }
   
    int flag=1;
    for(i=1;i<=26;i++){
        if(present[i]==0){
        flag=0;
                         }       
                       }
    if(flag==1){printf("pangram");}
    else printf("not pangram");   
    return 0;
}
