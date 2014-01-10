package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"
)

func findregex(includes, excludes []string) string {
	pool := regex_components(includes, excludes)
	var cover []string

	for len(includes) > 0 {
		max := 0
		max_loc := 0

		for i, c := range pool {
			v := 3*len(matches(c, includes)) - len(c)

			if v > max {
				max = v
				max_loc = i
			}
		}

		best := pool[max_loc]
		cover = append(cover, best)

		var leftover []string
		best_matches := matches(best, includes)

		for _, include := range includes {
			found := false

			for _, matched_include := range best_matches {
				if include == matched_include {
					found = true
				}
			}

			if found == false {
				leftover = append(leftover, include)
			}
		}

		includes = leftover

		var leftoverpool []string

		for _, c := range pool {
			if len(matches(c, includes)) > 0 {
				leftoverpool = append(leftoverpool, c)
			}
		}

		pool = leftoverpool
	}

	return strings.Join(cover, "|")
}

func regex_components(includes, excludes []string) []string {
	var parts []string
	var wholes []string

	for _, include := range includes {
		wholes = append(wholes, "^"+include+"$")
	}

	for _, w := range wholes {
		for _, p := range subparts(w) {
			for _, d := range dotify(p) {
				if m := matches(d, excludes); len(m) == 0 {
					parts = append(parts, d)
				}
			}
		}
	}

	for _, p := range parts {
		wholes = append(wholes, p)
	}

	return wholes
}

func subparts(word string) []string {
	var results []string

	for n := 1; n <= 4; n++ {
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
	var include = flag.String("include", "a,b,c", "what the regexp should match")
	var exclude = flag.String("exclude", "x,y,z", "what the regexp should not match")

	flag.Parse()

	includeRegexp := findregex(strings.Split(*include, ","), strings.Split(*exclude, ","))
	excludeRegexp := findregex(strings.Split(*exclude, ","), strings.Split(*include, ","))

	fmt.Println("regex for includes", includeRegexp)
	fmt.Println("regex for excludes", excludeRegexp)
}
