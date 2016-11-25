import java.io.BufferedReader;
import java.io.InputStreamReader;


class TestClass {
    public static void main(String args[] ) throws Exception {

        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        String basef = br.readLine();
        int base1 = Integer.parseInt(basef);
        int totalf=0;
        int base2;
        String line = br.readLine();
        int N = Integer.parseInt(line);
          //String bases = br.readLine();
          String []bases=br.readLine().split(" ");
    for(int i=0;i<bases.length;i++)
    {
       base2=Integer.parseInt(bases[i]);
      // System.out.println(base1+" "+base2);
      int min=base1;
      if(base2<base1) min=base2;
      int valid=0;
      for(int j=2;j<min;j++){
      	if(base1%j==0&&base2%j==0) valid=1;
      }
      if(valid==1){
      	totalf++;
      }
    }

System.out.println(totalf);


    }
}