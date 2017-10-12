import java.io.*;
import java.util.*;
import java.text.*;
import java.math.*;
import java.util.regex.*;

public class Solution {

    public static void main(String[] args) {
        Scanner in=new Scanner(System.in);
        int loop,j,flag=1;
        String input;
        loop=in.nextInt();
        for(int i=0;i<loop;i++){
            flag=1;
    input=in.next();        
    int len=input.length();
            if(len==1){ System.out.print("-1\n");
                       flag=0;
                      break;
                      }
    
            for(j=0;j<len/2;j++)
               {
       
                   if(input.charAt(j)!=input.charAt(len-1-j))
                    {
                      int next=j+1;
                      if(input.charAt(next)==input.charAt(len-1-j))
                            {
                              System.out.print(j+"\n");
                              flag=0;
                              break; 
                            }else  {
                                      System.out.print((len-1-j)+"\n");
                                      flag=0; break;}
            
        } 
        
                }
            if(flag==1)System.out.print("-1\n"); 
               
                 
        }  
    }
}

