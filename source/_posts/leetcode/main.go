package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(strongPasswordChecker("aaa111"))

}

func strongPasswordChecker(s string) int {
	var hasLowerCase bool
	var hasUpperCase bool
	var hasDigit bool

	requireContains := 0

	for i, _ := range s {
		if s[i] >= '0' && s[i] <= '9' {
			hasDigit = true
		}
		if s[i] >= 'a' && s[i] <= 'z' {
			hasLowerCase = true
		}

		if s[i] >= 'A' && s[i] <= 'Z' {
			hasUpperCase = true
		}
	}

	if !hasDigit {
		requireContains++
	}
	if !hasLowerCase {
		requireContains++
	}
	if !hasUpperCase {
		requireContains++
	}

	var availableDeletes, requiredDeletes int
	var availableAdditions, requiredAdditions int
	if len(s) > 6 {
		availableDeletes = len(s) - 6
	} else {
		requiredAdditions = 6 - len(s)
	}
	if len(s) < 20 {
		availableAdditions = 20 - len(s)
	} else {
		requiredDeletes = len(s) - 20
	}

	// f[i, remove, add, replace]
	f := make([][][][]bool, len(s)+1)
	for i := 0; i <= len(s); i++ {
		f[i] = make([][][]bool, availableDeletes+1)

		for d := 0; d <= availableDeletes; d++ {
			f[i][d] = make([][]bool, availableAdditions+1)

			for a := 0; a <= availableAdditions; a++ {
				f[i][d][a] = make([]bool, len(s)+1)

				for r := 0; r <= len(s); r++ {
					f[i][d][a][r] = false

					if i == 0 {
						if d == 0 && r == 0 {
							f[i][d][a][r] = true
						}
						continue
					}
					if d+r > i {
						continue
					}

					if d != 0 {
						f[i][d][a][r] = f[i][d][a][r] || f[i-1][d-1][a][r]
					}
					if a != 0 {
						f[i][d][a][r] = f[i][d][a][r] || f[i-1][d][a-1][r]
					}
					if r != 0 {
						f[i][d][a][r] = f[i][d][a][r] || f[i-1][d][a][r-1]
					}
					if i > 0 {
						f[i][d][a][r] = f[i][d][a][r] || f[i-1][d][a][r]
					}

					lastCharIdx := i - 1

					for k := i - 1; k > 0; k-- {
						needToBreak := false
						if s[k-1] == s[lastCharIdx] {
							repeatingLength := lastCharIdx - (k - 1) + 1
							if repeatingLength == 3 {
								needToBreak = true
							}
						} else {
							lastCharIdx = k - 1
						}
						f[i][d][a][r] = f[i][d][a][r] || f[k][d][a][r]
						if d != 0 {
							f[i][d][a][r] = f[i][d][a][r] || f[k][d-1][a][r]
						}
						if a != 0 {
							f[i][d][a][r] = f[i][d][a][r] || f[k][d][a-1][r]
						}
						if r != 0 {
							f[i][d][a][r] = f[i][d][a][r] || f[k][d][a][r-1]
						}
						if needToBreak {
							f[i][d][a][r] = false
							for j := k; j < i; j++ {
								if a != 0 {
									f[i][d][a][r] = f[i][d][a][r] || f[j-1][d][a-1][r]
								}
								if r != 0 {
									f[i][d][a][r] = f[i][d][a][r] || f[j-1][d][a][r-1]
								}
							}

							idx := lastCharIdx
							for idx >= 0 && s[idx] == s[lastCharIdx] {
								idx--
							}
							idx++
							repeatingCount := lastCharIdx - idx + 1
							deleting := repeatingCount - 2
							if d-deleting >= 0 {
								f[i][d][a][r] = f[i][d][a][r] || f[idx][d-deleting][a][r]
							}
							break
						}
					}
				}
			}
		}
	}
	result := math.MaxInt64
	for d := requiredDeletes; d <= availableDeletes; d++ {
		for a := requiredAdditions; a <= availableAdditions; a++ {
			for r := 0; r <= len(s); r++ {
				if f[len(s)][d][a][r] && a+r >= requireContains {
					if d+a+r < result {
						result = d + a + r
					}
				}
			}
		}
	}

	return result
}
