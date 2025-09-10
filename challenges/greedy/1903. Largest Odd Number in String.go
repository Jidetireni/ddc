package greedy

/*
The question we were ask to return the largest-valued odd integer
(as a string) for the iven string

To solve this, we can think greedily. If an odd number exists,
where must the number start from?
we can only know if a number is odd or not through the last digit

so what I did was to first convert the string to an array of bytes then
iterated from the end of the string toward the beginning.
As soon as I found an odd digit, I returned the
substring up to that digit.

*/

func LargestOddNumber(num string) string {
	l := len(num) - 1
	byteInt := []byte(num)
	for i := l; i >= 0; i-- {
		b := byteInt[i]
		if b != 0 && b%2 == 1 {
			return num[:i+1]
		}
	}

	return ""
}

// idomatic version
func largestOddNumber(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		ch := num[i]
		if (ch-'0')%2 == 1 {
			return num[:i+1]
		}
	}
	return ""
}
