package main

import (
	"fmt"
	"reflect"
)

// func Changename(mir []int) {
// 	mir[1] = 12

// }

// func plusOne(digits []int) []int {
// 	for i := len(digits) - 1; i >= 0; i-- {
// 		if digits[i] < 9 {
// 			digits[i]++
// 			return digits
// 		}
// 		digits[i] = 0
// 	}

// 	// если все цифры были 9, например [9, 9] → [1, 0, 0]
// 	return append([]int{1}, digits...)
// }

// func addBinary(a string, b string) string {
// 	result_of_a := 0
// 	result_of_b := 0
// 	power_of_a := 1
// 	power_of_b := 1
// 	var chars []byte
// 	if a == "0" && b == "0" {
// 		return "0"
// 	}
// 	for i := len(a) - 1; i >= 0; i-- {
// 		if a[i] == '0' || a[i] == '1' {
// 			bit := int(a[i] - '0')
// 			result_of_a += power_of_a * bit
// 			power_of_a *= 2
// 		}
// 	}

// 	for i := len(b) - 1; i >= 0; i-- {
// 		if b[i] == '0' || b[i] == '1' {
// 			bit := int(b[i] - '0')
// 			result_of_b += power_of_b * bit
// 			power_of_b *= 2
// 		}
// 	}
// 	sum_of_var := result_of_a + result_of_b

// 	for sum_of_var > 0 {
// 		digit := (sum_of_var % 2)
// 		char := byte('0' + digit)
// 		chars = append([]byte{char}, chars...)
// 		sum_of_var /= 2
//  	}
// 	return string(chars)

// }
// func convertToTitle(columnNumber int) string {

// }

// func longestCommonPrefix(slice []string) string {
// 	letter_of_slice := ""

// 	if 1 <= len(slice) && len(slice) <= 200 {
// 		minSlice := slice[0]
// 		for _, s := range slice {
// 			if len(s) < len(minSlice) {
// 				minSlice = s
// 			}
// 		}

// 		if 0 <= len(minSlice) && len(minSlice) <= 200 {
// 			for i := 0; i < len(minSlice); i++ {
// 				if minSlice[i] >= 'A' && minSlice[i] <= 'Z' {
// 					return ""
// 				}

// 				for j := 0; j < len(slice); j++ { // <--- исправил здесь: len(slice), а не len(minSlice)
// 					slice_clone := slice[j]
// 					if slice_clone[i] == minSlice[i] {
// 						continue
// 					} else {
// 						return letter_of_slice
// 					}
// 				}

// 				letter_of_slice += string(minSlice[i])
// 			}
// 		} else {
// 			return ""
// 		}
// 	}

//		return letter_of_slice
//	}
// func isHappy(n int) bool {
// 	var sum_of_slice int
// 	slice2 := []int{n}
// 	slice1 := []int{}
// 	for true {
// 		if n == 1 {
// 			return true
// 		}

// 		for n > 0 {
// 			slice1 = append([]int{n % 10}, slice1...)
// 			n /= 10

// 		}

// 		for _, var_slice := range slice1 {
// 			sum_of_slice += (var_slice * var_slice)
// 		}

// 		for i := 0; i < len(slice2); i++ {
// 			if sum_of_slice == slice2[i] {
// 				return false
// 			}
// 		}

// 		slice2 = append(slice2, sum_of_slice)
// 		slice1 = []int{}
// 		n = sum_of_slice
// 		sum_of_slice = 0
// 	}
// 	return true

// }
//
//	func addDigits(num int) int {
//		sum_of_slice_var := 0
//		for true {
//			slice1 := []int{}
//			if 0 <= num && num < 10 {
//				return num
//			}
//			for num > 0 {
//				slice1 = append([]int{num % 10}, slice1...)
//				num /= 10
//			}
//			for _, var_slice := range slice1 {
//				sum_of_slice_var += (var_slice)
//			}
//			num = sum_of_slice_var
//			sum_of_slice_var = 0
//		}
//		return num
//	}

//	func convertToTitle(columnNumber int) string {
//		var (
//			a int
//			b int
//		)
//		alfabit := "AABCDEFGHIJKLMNOPQRSTUVWXYZ"
//		if columnNumber <= 26 {
//			return string(alfabit[columnNumber])
//		} else {
//			a = columnNumber / 26
//			b = columnNumber % 26
//		}
//		return string(alfabit[a]) + string(alfabit[b])
//	}
//
//	func isPowerOfTwo(n int) bool {
//		for true {
//			if n == 1 {
//				return true
//			}
//			if n%2 != 0 {
//				return false
//			}
//			n /= 2
//		}
//		return true
//	}
// func isUgly(n int) bool {
// 	int1 := 2
// 	int2 := 3
// 	int3 := 5
// 	slice := []int{n}
// 	for true {
// 		if n == 1 {
// 			return true
// 		}
// 		if n%int1 == 0 {
// 			n /= int1
// 		}
// 		if n%int2 == 0 {
// 			n /= int2
// 		}
// 		if n%int3 == 0 {
// 			n /= int3
// 		}
// 		for _, var_slice := range slice {
// 			if n == var_slice {
// 				return false
// 			}
// 		}
// 		slice = append([]int{n}, slice...)

//		}
//		return true
//	}

// func missingNumber(nums []int) int {
// 	sumExpected, sumActual := 0, 0
// 	for i := 0; i <= len(nums); i++ {
// 		sumExpected += i
// 	}
// 	for _, var_of_slice := range nums {
// 		sumActual += var_of_slice
// 	}
// 	sumExpected -= sumActual
// 	return sumExpected
// }

// func isPowerOfThree(n int) bool {
// 	if n < 1 {
// 		return false
// 	}

// 	for n%3 == 0 {
// 		n /= 3
// 	}

//		return n == 1
//	}
// func checkPerfectNumber(num int) bool {
// 	sum_of_slice := 1
// 	if num == 1 {
// 		return false
// 	}
// 	for i := 2; i*i <= num; i++ {
// 		if num%i == 0 {
// 			sum_of_slice += i
// 			if i != num/i {
// 				sum_of_slice += num / i
// 			}
// 		}
// 	}
// 	return sum_of_slice == num
// }

// func fib(n int) int {
// 	sum_fibo := 0
// 	var sum1 int = 0
// 	var sum2 int = 1
// 	if n == 0 {
// 		return 0
// 	}
// 	if n == 1 {
// 		return 0
// 	}
// 	if 0 <= n && n <= 30 {
// 		for i := 2; i < n; i++ {
// 			sum_fibo = sum1 + sum2
// 			sum1 = sum2
// 			sum2 = sum_fibo
// 		}
// 	}
// 	return sum1 + sum2
// }

// func selfDividingNumbers(left int, right int) []int {
// 	slice := []int{}
// 	var_of_slice := []int{}
// 	if 1 <= left && left <= right {
// 		for i := left; i <= right; i++ {
// 			isValid := true
// 			lens := i
// 			for lens > 0 {
// 				var_of_slice = append([]int{lens % 10}, var_of_slice...)
// 				lens /= 10
// 			}
// 			for _, digit := range var_of_slice {
// 				if digit == 0 || i % digit != 0{
// 					isValid = false
// 					break
// 				}
// 			}
// 			if isValid {
// 				slice = append(slice, i)
// 			}
// 			var_of_slice = []int{}
// 		}
// 	}
// 	return slice
// }

// func maximumProduct(nums []int) []int {
// 	// if len(nums) == 3 {
// 	// 	return nums[0] * nums[1] * nums[2]
// 	// }
// 	for i := 0; i < len(nums); i++ {
// 		j := i + 1
// 		if j < len(nums) && nums[i] > nums[j] {
// 			carry := nums[i]
// 			nums[i] = nums[j]
// 			nums[j] = carry
// 		}
// 	}
// 	return nums
// }

// func isAnagram(s string, t string) bool {
// 	sum_1 := make([]int, 26)
// 	for _, var_1 := range s {
// 		sum_1[int(var_1)-'a']++
// 	}
// 	for _, var_1 := range t {
// 		sum_1[int(var_1)-'a']--
// 	}
// 	for _, v := range sum_1 {
// 		if v != 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	list1
// 	list2
// }

// func subtractProductAndSum(n int) int {
// 	if 0 >= n {
// 		return 0
// 	}
// 	sum, multiplication := 0, 1
// 	for n > 0 {
// 		multiplication *= (n % 10)
// 		sum += n % 10
// 		n /= 10
// 	}
// 	return multiplication - sum
// }

func main() {
	p := map[string]int{
		"miras": 12,
		"dias":  34,
	}
	v := reflect.ValueOf(p)
	fmt.Println(v.Type())
	fmt.Println(v.Kind())
}
