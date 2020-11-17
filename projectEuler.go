package main

import (
	"fmt"
	"strings"
	"strconv"
)


func main() {
	fmt.Println(problem_01(1000))
	fmt.Println(problem_02(4_000_000))
	fmt.Println(problem_03(600851475143))
	fmt.Println(problem_04(999))
	fmt.Println(problem_05(20))
	fmt.Println(problem_06(100))
	fmt.Println(problem_07(10001))
	fmt.Println(problem_08(13))
	fmt.Println(problem_09(1000))
	fmt.Println(problem_10(2_000_000))
	fmt.Println(problem_11())
}


// Returns the sum of all the multiples of 3 or 5 below maxnum
func problem_01(maxnum int) int {
	var sum int = 0
	for i := 0; i < maxnum; i++ {
		if i % 3 == 0 || i % 5 == 0 { sum += i }
	}
	return sum
}


// Returns the sum of even Fibonacci numbers under maxnum
func problem_02(maxnum int) int {
	bigFib, smallFib := 1, 1
	var sum int = 0
	for bigFib < maxnum {
		if bigFib % 2 == 0 { sum += bigFib }
		smallFib, bigFib = bigFib, bigFib + smallFib
	}
	return sum								   
}


// Returns the largest factor of numToFactor
func problem_03(numToFactor int) int {
	var largestFactor int = 1
	for i := 2; i <= numToFactor; i++ {
		for numToFactor % i == 0{
			largestFactor = i
			numToFactor /= i 
		}
	}
	return largestFactor
}


// Returns the largest palindrome from two numbers below factorCap
func problem_04(largestFactor int) int {
	// reverses an integer, 123 -> 321 ect.
	reverse_int := func(num int) int {
		reversed := 0
		for num > 0 {
			reversed = (reversed * 10) + num % 10 
			num /= 10
		}
		return reversed 
	}
	var i int = largestFactor // set to 999 for the problem_0
	var palindrome int = 0
	/* This function is structured to test all combinations
	 * 999*999, 999*988 ... but in a heavily optimised manner
	  
	  1 No need to search 99*99, 99*98... if the largest palindrome
	    is larger than 99*99, so we loop while i * i > palindrome
	  
	  2 Why search continue the search 999*91, 999*90 if the largest
	    palindrome is larger than 999*91
	    Also, since i >= j we half the number of cases since we do not
	    consider both 111*999 and 999*111 ect.
	 */

	for i * i > palindrome{                                 // 1
		for j := i; i * j > palindrome; j-- {               // 2
			num := i * j
			if num == reverse_int(num) && num > palindrome{
				// fmt.Println("Palindrome spotted: ", num)
				palindrome = num
			}
		}
		i -= 1
	}
	return palindrome
}


// Returns smallest number that is divisible by all of 1 to maxFactor
func problem_05(maxFactor int) int{
	var veryDivisible int = 1
	for i := 2; i < maxFactor + 1; i++{
		if veryDivisible % i != 0 {
			j := i
			for j < maxFactor { j *= i }
			veryDivisible *= j / i
		}
	}
	return veryDivisible
}


// Returns the difference between the sum of the squares of 1-n 
// and the sum of 1-n squared
func problem_06(n int) int {
	m := n * (n + 1) / 2
	return m * m - m * ( 2 * n + 1) / 3
}


// Returns the Nth prime
func problem_07(n int) int{
	// sieve of Eratosthenes, indexed to start with 1
	var list [120_000]bool
	list[0] = true
	var primesCaught = 0

	for i := 0; i < len(list); i++ {
		if ! list[i] { 
			for j := i; j < len(list); {
				list[j] = true
				j += i + 1
			}
			primesCaught += 1
			if primesCaught == n{
				return i + 1
			}
		}
	}
	fmt.Println("sieve too small:", primesCaught, "primes caught")
	return 0
}

// Returns largest product of numDig successive digits in given string
func problem_08(numDigits int) int{
	var bigNum string =
   `73167176531330624919225119674426574742355349194934
	96983520312774506326239578318016984801869478851843
	85861560789112949495459501737958331952853208805511
	12540698747158523863050715693290963295227443043557
	66896648950445244523161731856403098711121722383113
	62229893423380308135336276614282806444486645238749
	30358907296290491560440772390713810515859307960866
	70172427121883998797908792274921901699720888093776
	65727333001053367881220235421809751254540594752243
	52584907711670556013604839586446706324415722155397
	53697817977846174064955149290862569321978468622482
	83972241375657056057490261407972968652414535100474
	82166370484403199890008895243450658541227588666881
	16427171479924442928230863465674813919123162824586
	17866458359124566529476545682848912883142607690042
	24219022671055626321111109370544217506941658960408
	07198403850962455444362981230987879927244284909188
	84580156166097919133875499200524063689912560717606
	05886116467109405077541002256983155200055935729725
	71636269561882670428252483600823257530420752963450`
	bigNum = strings.Replace(bigNum, "\n", "", -1)
	bigNum = strings.Replace(bigNum, "\t", "", -1)
	bigNum = strings.Replace(bigNum,  " ", "", -1)

	// Returns the product of all digits in a string 
	// eg. 555 -> 625, 1234 -> 24 ect.
	getProduct := func(numStr string) (prod int) {
		num, error := strconv.Atoi(numStr)
		prod = 1
		for num != 0 && error == nil{
			prod *= num % 10
			num /= 10
		}
		return
	}
	// By removing zeros we eliminate 730 test cases
	var SplitNums = strings.Split(bigNum, "0")
	var largestProd = 0

	for _, numStr := range SplitNums{
		for i:= 0; i + numDigits <= len(numStr); i++ {
			prod := getProduct(numStr[i:i + numDigits])
			if prod > largestProd {
				largestProd = prod
			}
		}
	}
	return largestProd
}

// Returns the product of the pythagorean triplet that sums to TripletSum
func problem_09(TripletSum int) int {
	b := TripletSum / 2 + 1
	for a := 2; a < TripletSum;{
		continue_loop := true
		for continue_loop {
			c := TripletSum - a - b
			outcome := a * a + b * b - c * c 
			if outcome < 0 { continue_loop = false
			} else if outcome == 0 { return a * b * (TripletSum - a - b)
			} else if outcome > 0 { b -= 2 }
		}
		a += 2
	}
	return 0
}

// Returns the sum of all primes under maxnum
func problem_10(maxnum int) int{
	// sieve of Eratosthenes, indexed to start with 1
	list := make([]bool, maxnum)
	list[0] = true
	sum := 0
	for i := 0; i < len(list); i++ {
		if ! list[i] { 
			sum += i + 1
			for j := 2 * i + 1; j < len(list); {
				list[j] = true
				j += i + 1
			}
		}
	}
	return sum
}


// Returns the greatest product of four adjacent numbers in Arr
func problem_11() int{
var Arr = [20][20]int{
	[20]int{ 8, 2,22,97,38,15,00,40,00,75, 4, 5, 7,78,52,12,50,77,91, 8},
	[20]int{49,49,99,40,17,81,18,57,60,87,17,40,98,43,69,48, 4,56,62,00},
	[20]int{81,49,31,73,55,79,14,29,93,71,40,67,53,88,30, 3,49,13,36,65},
	[20]int{52,70,95,23, 4,60,11,42,69,24,68,56, 1,32,56,71,37, 2,36,91},
	[20]int{22,31,16,71,51,67,63,89,41,92,36,54,22,40,40,28,66,33,13,80},
	[20]int{24,47,32,60,99, 3,45, 2,44,75,33,53,78,36,84,20,35,17,12,50},
	[20]int{32,98,81,28,64,23,67,10,26,38,40,67,59,54,70,66,18,38,64,70},
	[20]int{67,26,20,68, 2,62,12,20,95,63,94,39,63, 8,40,91,66,49,94,21},
	[20]int{24,55,58, 5,66,73,99,26,97,17,78,78,96,83,14,88,34,89,63,72},
	[20]int{21,36,23, 9,75,00,76,44,20,45,35,14,00,61,33,97,34,31,33,95},
	[20]int{78,17,53,28,22,75,31,67,15,94, 3,80, 4,62,16,14, 9,53,56,92},
	[20]int{16,39, 5,42,96,35,31,47,55,58,88,24,00,17,54,24,36,29,85,57},
	[20]int{86,56,00,48,35,71,89, 7, 5,44,44,37,44,60,21,58,51,54,17,58},
	[20]int{19,80,81,68, 5,94,47,69,28,73,92,13,86,52,17,77, 4,89,55,40},
	[20]int{ 4,52, 8,83,97,35,99,16, 7,97,57,32,16,26,26,79,33,27,98,66},
	[20]int{88,36,68,87,57,62,20,72, 3,46,33,67,46,55,12,32,63,93,53,69},
	[20]int{ 4,42,16,73,38,25,39,11,24,94,72,18, 8,46,29,32,40,62,76,36},
	[20]int{20,69,36,41,72,30,23,88,34,62,99,69,82,67,59,85,74, 4,36,16},
	[20]int{20,73,35,29,78,31,90, 1,74,31,49,71,48,86,81,16,23,57, 5,54},
	[20]int{ 1,70,54,71,83,51,54,69,16,92,33,48,61,43,52, 1,89,19,67,48},
}
	prod := 1
	// horizontal and verticals
	for i := 0; i < 20 - 3; i++ {
		for j := 0; j < 20; j++ {
			x := Arr[i][j]*Arr[i+1][j]*Arr[i+2][j]*Arr[i+3][j]
			y := Arr[j][i]*Arr[j][i+1]*Arr[j][i+2]*Arr[j][i+3]
			if y > prod{
				prod = y
			}
			if x > prod{
				prod = x
			}
		}
	}

	// Diagonals
	for i := 0; i < 20 - 3; i++ {
		for j := 0; j < 20 - 3; j++ {
			x := Arr[i][j]*Arr[i+1][j+1]*Arr[i+2][j+2]*Arr[i+3][j+3]
			y := Arr[i][j+3]*Arr[i+1][j+2]*Arr[i+2][j+1]*Arr[i+3][j]
			if y > prod{
				prod = y
			}
			if x > prod{
				prod = x
			}
		}
	}
	return prod
}