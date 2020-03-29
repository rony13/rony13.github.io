---
title: 'LeetCode: Wildcard Matching'
date: 2020-03-25 13:48:16
tags:
---
```
func isMatch(s string, p string) bool {
	f := make([][]bool, len(p)+1)
	f[0] = make([]bool, len(s)+1)
	f[0][0] = true
	for i := 1; i <= len(p); i++ {
		f[i] = make([]bool, len(s)+1)
		if p[i-1] == '*' {
			f[i][0] = f[i][0] || f[i-1][0]
		}
		for j := 1; j <= len(s); j++ {
			if p[i-1] == '?' {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
			if p[i-1] == '*' {
				for k := 0; k <= j; k++ {
					f[i][j] = f[i][j] || f[i-1][k]
				}
			}
			if p[i-1] == s[j-1] {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[len(p)][len(s)]
}
```
