---
title: 'LeetCode: Valid Number'
date: 2020-03-12 13:48:17
tags:
---

```golang
func isNumber(s string) bool {
    pattern := `^\s*[+-]?((\d+(\.\d*)?)|(\.\d+))(e[+-]?\d+)?\s*$`
    match, _ := regexp.MatchString(pattern, s)
    return match
}
```
