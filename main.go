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

func candidate_components(winners, losers []string) {
	/*
	   parts = set(mappend(dotify, mappend(subparts, winners)))
	   wholes = {'^'+winner+'$' for winner in winners}
	   return wholes | {p for p in parts if not matches(p, losers)}
	*/
}

//func mappend(function, *sequences []string) {
// results = map(function, *sequences)
// return [item for result in results for item in result]
//}

func subparts(word string) {
	//  return set(word[i:i+n] for i in range(len(word)) for n in (1, 2, 3, 4))
}

func dotify(part string) []string {
	if part == "" {
		return []string{""}
	}

	var results []string

	for _, c := range []string{".", string(part[0])} {
		for i, rest := range dotify(part[1:]) {
			results = append(results, "")
			copy(results[i+1:], results[i:])
			results[i] = c+rest
		}
	}

	return results
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

func test() {
	/*
	   assert subparts('the')                                                    == {'t', 'h', 'e', 'th', 'he', 'the'}
	   assert subparts('this')                                                   == {'t', 'h', 'i', 's', 'th', 'hi', 'is', 'thi', 'his', 'this'}
	   assert candidate_components({'this'}, {'losers', 'something', 'history'}) == {'th.s', '^this$', '..is', 'this', 't.is', 't..s', '.his', '.h.s'}
	   assert mappend(reversed, ['abc', '123', 'WXYZ'])                          == ['c', 'b', 'a', '3', '2', '1', 'Z', 'Y', 'X', 'W']
	*/
}

func main() {
	verify("a|b", []string{"ark", "b", "c"}, []string{"art", "b", "x", "c"})
}
