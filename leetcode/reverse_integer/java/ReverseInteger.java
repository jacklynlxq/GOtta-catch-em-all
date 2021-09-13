// Compile with javac ReverseInteger.java
// Run with java -ea ReverseInteger
// Solution from leetcode

class ReverseInteger {
    public static int reverse(int x) {
        // variable to store the reversed number
        int rev = 0;
        while (x != 0) {
            // to get the number
            // x % 10 since we want to get the number of the smallest unit
            // popping the last digit and pushing it to the back of rev
            int pop = x % 10;

            // to get to the position (ie: unit, tenth, hundreth, ...)
            // x /= 10 because we want to get one tenth (0.1) of the number from x
            // akin to moving the decimal point by one significant number
            x /= 10;

            // condition to check where reversing x falls out of the
            // signed 32-bit integer range [-231, 231 - 1], return 0.
            // signed integer range: -2147483648 to 2147483647
            if (rev > Integer.MAX_VALUE/10 || (rev == Integer.MAX_VALUE / 10 && pop > 7)) return 0;
            if (rev < Integer.MIN_VALUE/10 || (rev == Integer.MIN_VALUE / 10 && pop < -8)) return 0;

            // pushing the pop to the rev
            // since everytime we push the pop to the rev
            // we increase the tenth of the previous numbers (base 10),
            // hence we have to do rev*10 (previous numbers * 10)
            rev = rev * 10 + pop;
        }
        return rev;
    }

    public static void main(String[] args){
        assert reverse(123) == 321 : "Failed";
        assert reverse(-123) == -321 : "Failed";
        assert reverse(120) == 21 : "Failed";
        assert reverse(0) == 0 : "Failed";

        // function execution time
        long startTime = System.nanoTime();
        reverse(123);
        long endTime = System.nanoTime();
        System.out.printf("Time elapsed: %s nanoseconds\n", endTime - startTime);
    }

}

