package main

func Changename(mir []int) {
	mir[1] = 12

}

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
func main() {
	slice := []int{0, 3, 5, 6, 4, 2}
	for i := 0; i <= len(slice); i++ {
		for j := 0; i < len(slice); j++ {
			if i == slice[j] {
				break
			} else {
				fmt.println(i)
			}
		}
	}
}
