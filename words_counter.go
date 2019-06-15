package wordscounter

import (
	"log"
	"sort"
	"strings"
	"unicode"
)

func removePunct(s string) string {
	return strings.Map(
		func(r rune) rune {
			if unicode.IsPunct(r) {
				return 32
			}
			return r
		},
		s,
	)
}

func stringToMap(s string) map[string]int {
	s = removePunct(s)
	m := map[string]int{}
	arr_s := strings.Fields(s)
	for _, v := range arr_s {
		m[strings.ToLower(v)] += 1
	}
	return m
}

func TopWordsByCount(c int, s string) map[string]int {
	resaultMap := map[string]int{}

	if c == 0 {
		return resaultMap
	}

	m := stringToMap(s)

	if len(m) < c {
		log.Printf("Max words in text = '%v'", len(m))
		return m
	}
	index := 0
	a := make([]int, len(m))
	for _, v := range m {
		a[index] = v
		index += 1
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	sliceTopCount := a[0:c]
	for k, v := range m {
		for i, value := range sliceTopCount {
			if v == value {
				resaultMap[k] = v
				c -= 1
				sliceTopCount[i] = 0
				if c == 0 {
					return resaultMap
				}
				break
			}
		}
	}
	return resaultMap
}
