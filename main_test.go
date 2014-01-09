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
		{"the", []string{"t", "h", "e", "th", "he", "the"}},
		{"this", []string{"t", "h", "i", "s", "th", "hi", "is", "thi", "his", "this"}},
	}

	for _, test := range tests {
		actual := subparts(test.input)

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

func TestMappend(t *testing.T) {
	tests := []struct {
		function mapper
		input    []string
		expected []string
	}{
		{dotify, []string{"abc", "123", "WXYZ"}, []string{"abc", "ab.", "a.c", "a..", ".bc", ".b.", "..c", "...", "123", "12.", "1.3", "1..", ".23", ".2.", "..3", "...", "WXYZ", "WXY.", "WX.Z", "WX..", "W.YZ", "W.Y.", "W..Z", "W...", ".XYZ", ".XY.", ".X.Z", ".X..", "..YZ", "..Y.", "...Z", "...."}},
	}

	for _, test := range tests {
		actual := mappend(test.function, test.input)

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

func TestCandidateComponents(t *testing.T) {
	actual := candidate_components([]string{"this"}, []string{"losers", "something", "history"})
	t.Errorf("%v", actual)

	// == {'th.s', '^this$', '..is', 'this', 't.is', 't..s', '.his', '.h.s'}
}
