package main

func main() {

	println(myAtoi("-13-2"))
	println(myAtoi("1-2"))
	println(myAtoi("- 912"))
	println(myAtoi("-91283472332"))
	println(myAtoi("-2147483648"))
	println(myAtoi("-2147483647"))
	println(myAtoi("2147483647"))
	println(myAtoi("2147483648"))
	println(myAtoi("9223372036854775808"))
	println(myAtoi("  987"))
	println(myAtoi("-3 1"))
	println(myAtoi("-0 1"))
	println(myAtoi("-42.11"))
	println(myAtoi("-42w11"))
	println(myAtoi("+-987-"))
	println(myAtoi("+987-"))
	println(myAtoi("-4193 with words"))
	println(myAtoi("words and 987"))
}

/**
Implement atoi which converts a string to an integer.
The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.
The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.
If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.
If no valid conversion could be performed, a zero value is returned.
Note:
Only the space character ' ' is considered as whitespace character.
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. If the numerical value is out of the range of representable values, INT_MAX (231 − 1) or INT_MIN (−231) is returned.
Example 1:
Input: "42"
Output: 42
Example 2:
Input: "   -42"
Output: -42
Explanation: The first non-whitespace character is '-', which is the minus sign.
             Then take as many numerical digits as possible, which gets 42.
Example 3:
Input: "4193 with words"
Output: 4193
Explanation: Conversion stops at digit '3' as the next character is not a numerical digit.
Example 4:
Input: "words and 987"
Output: 0
Explanation: The first non-whitespace character is 'w', which is not a numerical
             digit or a +/- sign. Therefore no valid conversion could be performed.
Example 5:
Input: "-91283472332"
Output: -2147483648
Explanation: The number "-91283472332" is out of the range of a 32-bit signed integer.
             Thefore INT_MIN (−231) is returned.
*/
const maxUint = ^uint32(0)
const maxInt = int(maxUint >> 1)
const minInt = -maxInt - 1

func myAtoi(str string) int {
	i := 0
	var result int
	digital := false
	space := false
	nodigital_char := false
	sign := true
	signed := false
	exced := false
	for ; i < len(str); i++ {
		value := int(str[i])
		if value == 32 && (digital || signed) {
			space = true
			continue
		}
		if value == 43 || value == 45 {
			if digital {
				break
			}
			if signed && !digital {
				return 0
			}
			if !signed && value == 45 {
				sign = false
			}
			signed = true
			continue
		}
		if value > 47 && value < 58 {
			if !space && !nodigital_char {
				result *= 10
				result += value - 48
				if result > maxInt {
					exced = true
				}
			}
			digital = true
			continue
		}
		if digital {
			nodigital_char = true
		} else if value != 32 {
			return 0
		}
	}
	if !sign {
		result = -result
	}
	if exced && sign {
		return maxInt
	}
	if exced && !sign {
		return minInt
	}
	return result
}
