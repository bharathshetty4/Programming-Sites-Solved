import java.io.*;
import java.util.*;
import java.text.*;
import java.math.*;
import java.util.regex.*;

public class Solution {

    public static void main(String[] args) {
        
                Scanner in=new Scanner(System.in);
        int loop,j,count=0;
        String input;
        int a[]=new int[26];
            
         for(j=0;j<26;j++){
                a[j]=0;
                
                            }
        loop=in.nextInt();
        for(int i=1;i<=loop;i++)
              {
            int flag=0;
            //System.out.println("HE");
            input=in.next();
            int len=input.length();
            for(j=0;j<len;j++){
                if(a[input.charAt(j)-'a']==(i-1))
                a[input.charAt(j)-'a']=i;
            }
           }
        
        for(j=0;j<26;j++){
               if(a[j]==loop){
                   count++;
                   //System.out.println(j);
               }
                
          }
        System.out.print(count);
    }
}
    
