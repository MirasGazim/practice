package main

import "fmt"

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

func longestCommonPrefix(slice []string) string {
	var letter_of_slice string = ""
	minSlice := slice[0]
	for _, s := range slice {
		if len(s) < len(minSlice) {
			minSlice = s
		}
	}

	if 1 <= len(slice) && len(slice) <= 200 && 0 <= len(minSlice) && len(minSlice) <= 200 {
		for i := 0; i < len(minSlice); i++ {
			if minSlice[i] >= 'A' && minSlice[i] <= 'Z' {
				return ""
			}

			for j := 0; j < len(minSlice); j++ {
				slice_clone := slice[j]
				if slice_clone[i] == minSlice[i] {
					continue
				} else {
					return letter_of_slice
				}
			}
			letter_of_slice += string(minSlice[i])

		}
	} else {
		return ""
	}

	return letter_of_slice
}

func main() {
	slice := []string{}
	fmt.Println(longestCommonPrefix(slice))

}
