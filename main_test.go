package main

import (
	"testing"
)

func TestMatches(t *testing.T) {
	tests := []struct {
		regex    string
		input    []string
		expected []string
	}{
		{"a|b|c", []string{"a", "b", "c", "d", "e"}, []string{"a", "b", "c"}},
		{"a|b|c", []string{"any", "bee", "succeed", "dee", "eee!"}, []string{"any", "bee", "succeed"}},
	}

	for _, test := range tests {
		actual := matches(test.regex, test.input)

		if len(actual) != len(test.expected) {
			t.Errorf("expected: %v actual: %v", test.expected, actual)
			break
		}

		for i := range test.expected {
			if actual[i] != test.expected[i] {
				t.Errorf("expected: %v actual: %v", test.expected, actual)
				break
			}
		}
	}
}

func TestReplacements(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"x", "x."},
		{"^", "^"},
		{"$", "$"},
	}

	for _, test := range tests {
		actual := replacements(test.input)

		if actual != test.expected {
			t.Errorf("expected: %v actual: %v", test.expected, actual)
		}
	}
}

func TestDotify(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"", []string{""}},
		{"it", []string{"it", "i.", ".t", ".."}},
		{"^it$", []string{"^it$", "^i.$", "^.t$", "^..$"}},
		{"this", []string{"this", "thi.", "th.s", "th..", "t.is", "t.i.", "t..s", "t...", ".his", ".hi.", ".h.s", ".h..", "..is", "..i.", "...s", "...."}},
	}

	for _, test := range tests {
		actual := dotify(test.input)

		if len(actual) != len(test.expected) {
			t.Errorf("expected: %v actual: %v", test.expected, actual)
			break
		}

		for i := range test.expected {
			if actual[i] != test.expected[i] {
				t.Errorf("expected: %v actual: %v", test.expected, actual)
				break
			}
		}
	}
}

func TestSubParts(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"^it$", []string{"^", "i", "t", "$", "^i", "it", "t$", "^it", "it$", "^it$"}},
		{"the", []string{"t", "h", "e", "th", "he", "the"}},
		{"this", []string{"t", "h", "i", "s", "th", "hi", "is", "thi", "his", "this"}},
		{"^win$", []string{"^", "w", "i", "n", "$", "^w", "wi", "in", "n$", "^wi", "win", "in$", "^win", "win$"}},
	}

	for _, test := range tests {
		actual := subparts(test.input)

		if len(actual) != len(test.expected) {
			t.Errorf("expected: %v actual: %v", test.expected, actual)
			continue
		}

		for i := range test.expected {
			if actual[i] != test.expected[i] {
				t.Errorf("expected: %v actual: %v", test.expected, actual)
				continue
			}
		}
	}
}

func TestRegexComponents(t *testing.T) {
	tests := []struct {
		winners  []string
		losers   []string
		expected []string
	}{
		{[]string{"win"}, []string{"losers", "bin", "won"}, []string{"^win$", "wi", "^wi", "win", "wi.", "^win", "^wi.", "win$", "wi.$"}},
	}

	for _, test := range tests {
		actual := regex_components(test.winners, test.losers)

		if len(actual) != len(test.expected) {
			t.Errorf("expected: %v actual: %v", test.expected, actual)
			break
		}

		for i := range test.expected {
			if actual[i] != test.expected[i] {
				t.Errorf("expected: %v actual: %v", test.expected, actual)
				break
			}
		}
	}
}

func TestFindRegex(t *testing.T) {
	boys := []string{"jacob", "mason", "ethan", "noah", "william", "liam", "jayden", "michael", "alexander", "aiden"}
	girls := []string{"sophia", "emma", "isabella", "olivia", "ava", "emily", "abigail", "mia", "madison", "elizabeth"}

	nfl_in := []string{"colts", "saints", "chargers", "49ers", "seahawks", "patriots", "panthers", "broncos", "chiefs", "eagles", "bengals", "packers"}
	nfl_out := []string{"jets", "dolphins", "bills", "steelers", "ravens", "browns", "titans", "jaguars", "texans", "raiders", "cowboys", "giants", "redskins", "bears", "lions", "vikings", "falcons", "buccaneers", "cardinals", "rams"}

	tests := []struct {
		winners  []string
		losers   []string
		expected string
	}{
		{[]string{"ahahah", "ciao"}, []string{"ahaha", "bye"}, "a.$"},
		{boys, girls, "a.$|e.$|a.o"},
		{girls, boys, "a$|^..i|is"},
		{nfl_in, nfl_out, "^p|g..s|4|lt|sa|ch|ha|nc"},
		{nfl_out, nfl_in, "ns|^.i|j|d|ee|y|m|^bears$"},
	}

	for _, test := range tests {
		actual := findregex(test.winners, test.losers)

		if actual != test.expected {
			t.Errorf("expected: %v actual: %v", test.expected, actual)
		}
	}
}
