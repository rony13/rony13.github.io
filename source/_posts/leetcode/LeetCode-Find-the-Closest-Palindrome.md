---
title: 'LeetCode: Find the Closest Palindrome'
date: 2020-03-15 14:58:18
tags:
---
```
func nearestPalindromic(n string) string {
	bytes := []rune(n)
	for i := 0; i < len(bytes)/2; i++ {
		bytes[len(bytes)-1-i] = bytes[i]
	}

	result := string(bytes)

	left := []rune(n)[0 : len(n)/2]
	odd := len(n)%2 == 1
	if odd {
		a := bytes[len(bytes)/2]
		if a == '0' {
			result = closer(n, result, string(left)+"1"+reverse(string(left)))
		}
		if a == '9' {
			result = closer(n, result, string(left)+"8"+reverse(string(left)))
		}
		if a > '0' && a < '9' {
			bigger := string(left) + string(a+1) + reverse(string(left))
			smaller := string(left) + string(a-1) + reverse(string(left))
			result = closer(n, result, closer(n, bigger, smaller))
		}
	} else {
		a := left[len(left)-1]
		if a == '0' {
			left[len(left)-1] = '1'
			result = closer(n, result, string(left)+reverse(string(left)))
		}
		if a == '9' {
			left[len(left)-1] = '8'
			result = closer(n, result, string(left)+reverse(string(left)))
		}
		if a > '0' && a < '9' {
			left[len(left)-1] = a + 1
			bigger := string(left) + reverse(string(left))
			left[len(left)-1] = a - 1
			smaller := string(left) + reverse(string(left))
			result = closer(n, result, closer(n, bigger, smaller))
		}
	}

	minimum := make([]rune, len(n)-1)
	if len(n) > 1 {
		for i := 0; i < len(n)-1; i++ {
			minimum[i] = '9'
		}
	} else {
		minimum = []rune{'0'}
	}
	maximum := make([]rune, len(n)+1)
	for i := 0; i < len(n)+1; i++ {
		if i == 0 || i == len(n) {
			maximum[i] = '1'
		} else {
			maximum[i] = '0'
		}
	}
	return closer(n, closer(n, result, string(minimum)), string(maximum))

}

func closer(target, a, b string) string {
	ti, _ := strconv.ParseInt(target, 10, 64)
	ai, _ := strconv.ParseInt(a, 10, 64)
	bi, _ := strconv.ParseInt(b, 10, 64)
	if ti == ai {
		return b
	}
	if ti == bi {
		return a
	}
	if abs(ti-ai) < abs(ti-bi) {
		return a
	}
	if abs(ti-ai) > abs(ti-bi) {
		return b
	}

	if ai < bi {
		return a
	}
	return b
}

func abs(a int64) int64 {
	if a > 0 {
		return a
	} else {
		return -a
	}

}

func reverse(input string) string {
	if input == "" {
		return input
	}
	output := make([]rune, len(input))
	for i, c := range input {
		output[len(input)-1-i] = c
	}
	return string(output)

}
```
