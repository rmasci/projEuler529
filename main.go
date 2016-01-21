package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// initialize arg as a string.
	var arg string
	fmt.Println("Please input a numbers, ex: 19028037046055")
	for {
		fmt.Printf("Numbers: ")
		// User types in a number...hopefully more than 2 digits
		fmt.Scanln(&arg)
		if len(arg) <= 1 {
			fmt.Printf("C'mon, give me a number I can work with here!\n")
		} else {
			fmt.Printf("Great! you've given me the number %v that has %v digits!\n", arg, len(arg))
			break
		}
	}
	// how long is the string
	lArg := len(arg)
	//split the string into an slice we can loop through.
	a := strings.Split(arg, "")
	// make another slice
	arg2 := make([]string, lArg)
	// We need two main loops, Loop 1 loops through the string. Loop 2 also
	// loops through the string but starts at the i1 position, and once it has
    // a substring that equals 10, it goes to loop3. Loop 3 looks for any '0' that
    // might follow.  5500055. [55]00055 equals 10, but so does [550]0055, and 
    // [5500]055. If we exited loop2 without loop3 we'd loose the '0's 
    fmt.Println("Substrings that equal 10:")
	for i1 := 0; i1 <= lArg-1; i1++ {
		// ct is used for the addition of the substring.
		ct := 0
		// Loop through a substring of the original string.
		for i2 := i1; i2 <= lArg-1; i2++ {
			// use strconv to convert the string a[i2] to an integer
			n, _ := strconv.Atoi(a[i2])
			// add ct + n
			ct = ct + n
			// Check if we've reached 10. If so we can stop and go to the
			// next subsstring
			if ct == 10 {
				// print that we've reached 10
				fmt.Println("\t", arg[i1:i2+1])
				// loop through and assign to proper position in
				// arg2. arg[i1:i2+1] is a another way to slice just
				// the substring from the original arg.
				for ary, aStr := range arg[i1 : i2+1] {
					arg2[ary+i1] = string(aStr)
				}
				// zero is hard to catch, go into loop 3.
				i3 := i2 + 1
				i := 1
				for i3 <= lArg-1 {
					if a[i3] == "0" {
						fmt.Println("\t", arg[i1:i2+i]+"0")
						arg2[i2+i] = "0"
						//i2++
						i3++
						i++
					} else {
						// if it doesn't equal 0, get out of the loop
						i3 = lArg + 1
					}
				}
				//break out of l2 and go back to l1
				break
				// if the substring is greater than 10 such as 549
				// 5+4 < 10, but 5+4+9 = 18
			} else if ct > 10 {
				break
			}
		}
	}
	// Check the original string input by the user against the
	// string we built that equal 10.
	if strings.Join(arg2, "") == arg {
		fmt.Printf("\nAll the substrings in the number %v equal 10", arg)
	} else {
		fmt.Printf("\nAll the substrings in the number %v do not equal 10\n", arg)
	}
}


/*
https://projecteuler.net/problem=529

10-substrings
Problem 529
A 10-substring of a number is a substring of its digits that sum to 10. For example, the 10-substrings of the number 3523014 are:

3523014
3523014
3523014
3523014
A number is called 10-substring-friendly if every one of its digits belongs to a 10-substring. For example, 3523014 is 10-substring-friendly, but 28546 is not.

Let T(n) be the number of 10-substring-friendly numbers from 1 to 10n (inclusive).
For example T(2) = 9 and T(5) = 3492.

Find T(1018) mod 1 000 000 007.
*/