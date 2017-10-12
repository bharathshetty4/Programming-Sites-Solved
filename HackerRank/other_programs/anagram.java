import java.io.*;
import java.util.*;
import java.text.*;
import java.math.*;
import java.util.regex.*;

public class Solution {

    public static void main(String[] args) {
        
        Scanner in=new Scanner(System.in);
        int loop,j;
        String input;
        loop=in.nextInt();
        for(int i=0;i<loop;i++)
              {
            int flag=0;
            //System.out.println("HE");
            int a[]=new int[26];
            int b[]=new int[26];
            
            for(j=0;j<26;j++){
                a[j]=0;
                b[j]=0;
                            }
            
            input=in.next();        
             int len=input.length();
            int chang=0;
            if(len%2==1)
                  { System.out.print("-1\n");
                    
                  }else
                       { 
    
            for(j=0;j<len/2;j++)
            {
                a[input.charAt(j)-'a']++;
            }
                for(j=len/2;j<len;j++)
                {
                    b[input.charAt(j)-'a']++;
                
                }
                
                
                for(j=0;j<26;j++)
                  {
                  if(b[j]>a[j]){
                      chang=chang+b[j]-a[j];
                      
                  }
              }
                
                System.out.print(chang+"\n");
                
                         }
               
        
        
    }
    }
}
