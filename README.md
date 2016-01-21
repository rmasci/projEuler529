Richard Masci -- Woodstock, GA
I started out on Apple ][+ in 1980, and now work as a Team Lead at AT&T. I use Go all the time as a
replacement to Bash, PHP, Perl, TCL, Ruby... Once I found that Go could parse through a 500mb csv file
in 2 seconds vs. the 9 seconds in Ruby... I was hooked because I had 2000+ files to process in 5 locations a day.

Best feature of Go is that I can write it on a Mac. It then runs on Linux, Windows, Raspberry PI. Crosscompile is 
VERY easy, and I don't have to worry about what version of perl, or what extensions are installed. 

# projEuler529
This is project 529 from https://projecteuler.net/problem=529

# Description
What this progam does is ask the user to input any number they desire... Then look to see if there are substrings
in that file that equal 10.  For example if the user entered:<pre>
    Please input some numbers, ex: 19028037046055
    Numbers: 2
    C'mon, give me a number I can work with here!
    Numbers:
    </pre> 
It looks to see if the user entered a number that even has a substring.  If they entered only one digit, try again.

Now if they enter number with two digits, then it will try to find if there is a substring in that number that
equals 10.  For example:a<pre>
    Please input some numbers, ex: 19028037046055
    Numbers: 55
    Substrings that equal 10:
	   55

    All the substrings in the number 55 equal 10
 </pre>
 We need two main loops, Loop 1 (i1) loops through the string. Loop 2 (i2) also loops through the string but starts at the 
 i1 position, and once it has a substring that equals 10, it goes to loop3 (i3). Loop 3 looks for any '0' that might follow.
 <pre>Example: 5500055 
          [55]00055 equals 10 
          [550]0055 also equals 10
          [5500]055 equals 10 as well. If we exited loop2 when we reached 10 without loop3 that checks for a '0' after 
          loop2 hits 10... we'd loose the '0's
            </pre>
 Here's what that looks like when we run it:
    <pre>Please input some numbers, ex: 19028037046055
    Numbers: 5500055
    Great! you've given me the number 5500055 that has 7 digits!
    Substrings that equal 10:
	   55
	   550
	   5500
	   55000
	   50005
       00055
	   0055
	   055
	   55
     
    All the substrings in the number 5500055 equal 10 </pre>