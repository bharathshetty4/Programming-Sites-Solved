import java.io.*;
import java.util.*;
import java.text.*;
import java.math.*;
import java.util.regex.*;

public class Solution {

    public static void main(String[] args) {
        Scanner myScan = new Scanner(System.in);
        
       int j;
        int a[]=new int[26];
        int b[]=new int[26];
        int sum=0;
        for(j=0;j<26;j++){
                a[j]=0;
            b[j]=0;
                          }
        
        String inputString1 = myScan.next();
        int len1=inputString1.length();
        String inputString2 = myScan.next();
        int len2=inputString2.length();
        
        for(j=0;j<len1;j++)
            {
                a[inputString1.charAt(j)-'a']++;
            }
        for(j=0;j<len2;j++)
            {
                b[inputString2.charAt(j)-'a']++;
            }
        
        for(j=0;j<26;j++){
            
            if(a[j]!=b[j]){
            if(a[j]>b[j]) sum+=a[j]-b[j];
                else sum+=b[j]-a[j];
            }
            
        }
        
        System.out.println(sum);
        
    }
    
    
}
