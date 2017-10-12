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
        for(j=0;j<loop;j++)
              {
            input=in.next();
            int count=0;
            int len=input.length();
            for(int i=0;i<len-1;i++){
                if(input.charAt(i)==input.charAt(i+1)){
                    count++;
                }
                
            
            }
            System.out.println(count);
        
            }
        
        
    }
}
