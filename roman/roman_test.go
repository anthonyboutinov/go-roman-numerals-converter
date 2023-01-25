package roman_test

import (
	"testing"

	"github.com/anthonyboutinov/go-roman-numerals-converter/roman"
)

type testCase struct {
	name    string
	roman   string
	arabic  int
	wantErr bool
}

// Run tests for a list of test cases with a closure function
func runTests(t *testing.T, testCases []testCase, testFunc func(testCase) (interface{}, error)) {
	for _, test := range testCases {
		got, error := testFunc(test)
		if test.wantErr {
			if error == nil {
				t.Errorf(`Test case '%v': expected error, got nil`, test.name)
				continue
			}
		} else {
			if error != nil {
				t.Errorf(`Test case '%v': unexpected error: %v`, test.name, error)
				continue
			}
			if got != test.arabic {
				t.Errorf(`Test case '%v': expected %v, got %v`, test.name, test.arabic, got)
			}
		}
	}
}

func TestRomanToInt(t *testing.T) {
	testCases := []testCase{
		{
			name:    "Valid roman numeral CXXXIV",
			roman:   "CXXXIV",
			arabic:  134,
			wantErr: false,
		},
		{
			name:    "Valid lowercase roman numeral cxxXiv",
			roman:   "cxxXiv",
			arabic:  134,
			wantErr: false,
		},
		{
			name:    "Valid roman numeral MDCLXVIII",
			roman:   "MDCLXVIII",
			arabic:  1668,
			wantErr: false,
		},
		{
			name:    "Jibberish of roman symbols ICXLXXIVMD",
			roman:   "ICXLXXIVMD",
			arabic:  0,
			wantErr: true,
		},
		{
			name:    "Invalid characters: XI6V",
			roman:   "XI6V",
			arabic:  0,
			wantErr: true,
		},
		{
			name:    "Empty string",
			roman:   "",
			arabic:  0,
			wantErr: true,
		},
	}

	runTests(t, testCases, func(elem testCase) (interface{}, error) {
		return roman.RomanToInteger(elem.roman)
	})
}

func TestIntToRoman(t *testing.T) {
	testCases := []testCase{
		{
			name:    "Valid number 628",
			arabic:  628,
			roman:   "DCXXVIII",
			wantErr: false,
		},
		{
			name:    "Invalid (negative) number -34",
			arabic:  -34,
			roman:   "",
			wantErr: true,
		},
		{
			name:    "Invalid number 0",
			arabic:  0,
			roman:   "",
			wantErr: true,
		},
		{
			name:    "Invalid number 3999 + 1. The largest numeral is MMMCMXCIX (3999)",
			arabic:  3999 + 1,
			roman:   "",
			wantErr: true,
		},
	}

	runTests(t, testCases, func(elem testCase) (interface{}, error) {
		return roman.IntegerToRoman(elem.arabic)
	})

}