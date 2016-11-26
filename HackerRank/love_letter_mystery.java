import java.io.*;
import java.util.*;
import java.text.*;
import java.math.*;
import java.util.regex.*;

public class Solution {

    public static void main(String[] args) {
        
        Scanner in=new Scanner(System.in);
            
        int no=in.nextInt();
        int i;
        int j;
        int len;
        String str;
        for(j=0;j<no;j++){
           
            int ans=0,finanswer=0;
          str=in.next();
            
            len=0;
            len=str.length();
            
            
            for(i=0;i<(len/2);i++){
               
                if(str.charAt(i)==str.charAt(len-1-i)){
                }else{
                    ans=str.charAt(i)-str.charAt(len-1-i);
                    if(ans<0) ans=ans*(-1);
                    
                    finanswer=finanswer+ans;
                }
                
            }
            
          System.out.println(finanswer);  
            
            
        }
        
    }
}
