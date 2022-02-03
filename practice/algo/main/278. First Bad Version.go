package main

import "fmt"

func firstBadVersion(n int) int {
	ini := 1
	mid := 0

	for ini <= n {
		mid = (ini + n) / 2
		if !isBadVersion(mid) {
			ini = mid + 1
		} else {
			for i := ini; i <= mid; i++ {
				if isBadVersion(i) {
					fmt.Println(i)
					return i
				}
			}

		}
	}
	return mid
}

func isBadVersion(n int) bool {
	key := 1
	if n >= key {
		return true
	}
	return false
}
