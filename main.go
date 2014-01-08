package main

import (
	"fmt"
	"regexp"
)

func intersect(a, b []string) []string {
	var r []string

	for _, i := range a {
		exists := false

		for _, j := range b {
			if i == j {
				exists = true
			}
		}

		if exists == true {
			r = append(r, i)
		}
	}

	return r
}

func verify(regex string, a, b []string) {
	r, err := regexp.Compile(regex)

	if err != nil {
		fmt.Println(err)
	}

	var matched []string
	var unmatched []string

	for _, i := range a {
		if r.MatchString(i) == true {
			matched = append(matched, i)
		} else {
			unmatched = append(unmatched, i)
		}
	}
}

func main() {
	verify("a", []string{"ark", "b", "c"}, []string{"a", "b", "x", "c"})
}
