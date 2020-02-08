/**
For a positive integer n, let f(n) denote the sum
of the digits of n when represented in base 10.
It is easy to see that the sequence of numbers
n, f(n), f(f(n)), f(f(f(n))), . . . eventually becomes a
single digit number that repeats forever. Let this single digit be denoted g(n).
For example, consider n = 1234567892. Then:
f(n) = 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 2 = 47
f(f(n)) = 4 + 7 = 11
f(f(f(n))) = 1 + 1 = 2
Therefore, g(1234567892) = 2.
 */

package main

import "fmt"

func main() {
	in := 1234567892
	for {
		in = f(in)
		if in < 10 {
			break
		}
	}
	fmt.Println(in)
}

func f(n int) int {
	s := 0
	var i int
	for i = n; i > 10; i = i / 10 {
		s += i%10
	}
	return s+i
}
