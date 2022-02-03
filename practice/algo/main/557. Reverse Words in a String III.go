package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	s1 := strings.Split(s, " ")
	//fmt.Println(s1)
	for s2I, s2 := range s1 {
		s1[s2I] = reverseString([]byte(s2))
	}
	fmt.Println(strings.Join(s1, " "))
	return strings.Join(s1, " ")
}
