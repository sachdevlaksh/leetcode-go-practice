package main

func reverseString(s []byte) string {
	low := 0
	high := len(s) - 1
	for low < high {
		temp := s[low]
		s[low] = s[high]
		s[high] = temp
		low++
		high--
	}
	return string(s)
}
