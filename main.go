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

func verify(regex string, a, b []string) bool {
	r, err := regexp.Compile(regex)

	if err != nil {
		fmt.Println(err)
	}

	var unmatchedA []string
	var matchedB []string

	for _, i := range a {
		if r.MatchString(i) == false {
			unmatchedA = append(unmatchedA, i)
		}
	}

	for _, i := range b {
		if r.MatchString(i) == true {
			matchedB = append(matchedB, i)
		}
	}

	result := true

	if len(unmatchedA) > 0 {
		result = false
		fmt.Println("Missed ", unmatchedA)
	}

	if len(matchedB) > 0 {
		result = false
		fmt.Println("Should have missed ", matchedB)
	}

	return result
}

func findregex(winners, losers []string) {
	/*
	   pool = candidate_components(winners, losers)
	   cover = []
	   while winners:
	       best = max(pool, key=lambda c: 3*len(matches(c, winners)) - len(c))
	       cover.append(best)
	       pool.remove(best)
	       winners = winners - matches(best, winners)
	   return '|'.join(cover)
	*/
}

func regex_components(winners, losers []string) []string {
	var parts []string
	var wholes []string

	for _, winner := range winners {
		wholes = append(wholes, "^"+winner+"$")
		parts = append(parts, "^"+winner+"$")
	}

	for _, w := range wholes {
		for _, p := range subparts(w) {
			for _, d := range dotify(p) {
				if m := matches(d, losers); len(m) == 0 {
					parts = append(parts, d)
				}
			}
		}
	}

	return parts
}

func subparts(word string) []string {
	if word == "" {
		return []string{""}
	}

	var results []string

	for _, n := range []int{1, 2, 3, 4} {
		for i := 0; i < len(word); i++ {
			if i+n > len(word) {
				continue
			}
			results = append(results, word[i:i+n])
		}
	}

	return results
}

func dotify(part string) []string {
	if part == "" {
		return []string{""}
	}

	var results []string

	for _, c := range replacements(string(part[0])) {
		for _, rest := range dotify(part[1:]) {
			results = append(results, string(c)+rest)
		}
	}

	return results
}

func replacements(char string) string {
	if char == "^" || char == "$" {
		return char
	}

	return char + "."
}

func matches(regex string, strings []string) []string {
	var results []string

	r, err := regexp.Compile(regex)

	if err != nil {
		panic(err)
	}

	for _, i := range strings {
		if r.MatchString(i) == true {
			results = append(results, i)
		}
	}

	return results
}

func main() {
	verify("a|b", []string{"ark", "b", "c"}, []string{"art", "b", "x", "c"})
}
