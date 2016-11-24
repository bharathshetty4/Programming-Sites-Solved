using System;
class MyClass {
    static void Main(string[] args) {
        // Read input from stdin and provide input before running
        var line1 = System.Console.ReadLine().Trim();
        var N = Int32.Parse(line1);
        for (var i = 0; i < N; i++) {
            var line2 = System.Console.ReadLine().Trim();
            var digit = Int32.Parse(line2);
            var answer = ((digit*6)+(digit*2)+((digit-1)*6));
            System.Console.WriteLine(answer);
        }
    }
}