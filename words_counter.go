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

func makeArr(i int) []int {
	x := make([]int, i)
	return x
}

func TopWordsByCount(c int, s string) map[string]int {
	resaultMap := map[string]int{}

	if c == 0 {
		return resaultMap

	}
	m := stringToMap(s)

	index := 0
	a := makeArr(len(m))
	if len(m) < c {
		log.Printf("Max words in text = '%v'", len(m))
		return m
	}
	for _, v := range m {
		a[index] = v
		index += 1
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	for k, v := range m {
		for i := range a {
			if v == i {
				resaultMap[k] = v
				c -= 1
				if c == 0 {
					return resaultMap
				}
			}
		}
	}
	return resaultMap

}