package functions

import "testing"

//
// =====================================================
// Main Test Suite for ProcessText
// =====================================================
//

// TestProcessText runs multiple independent test cases
// to validate the behavior of the full text pipeline.
func TestProcessText(t *testing.T) {

	// Define multiple test scenarios
	tests := []struct {
		name     string // Name of the test case
		input    string // Input text
		expected string // Expected processed output
	}{

		//
		// -------------------------------------------------
		// Test 1: Basic punctuation spacing
		// Ensures punctuation attaches to previous word
		// -------------------------------------------------
		{
			name:     "Punctuation spacing",
			input:    "hello , world !",
			expected: "hello, world!",
		},

		//
		// -------------------------------------------------
		// Test 2: Punctuation groups
		// Ensures sequences like ... or !! remain intact
		// -------------------------------------------------
		{
			name:     "Punctuation groups",
			input:    "wait ... what ?!",
			expected: "wait... what?!",
		},

		//
		// -------------------------------------------------
		// Test 3: Hexadecimal conversion
		// 1e (hex) = 30 in decimal
		// -------------------------------------------------
		{
			name:     "Hex conversion",
			input:    "1e (hex)",
			expected: "30",
		},

		//
		// -------------------------------------------------
		// Test 4: Binary conversion
		// 101 (bin) = 5 in decimal
		// -------------------------------------------------
		{
			name:     "Binary conversion",
			input:    "101 (bin)",
			expected: "5",
		},

		//
		// -------------------------------------------------
		// Test 5: Uppercase transformation
		// Applies (up) to the previous word
		// -------------------------------------------------
		{
			name:     "Uppercase transformation",
			input:    "hello (up)",
			expected: "HELLO",
		},

		//
		// -------------------------------------------------
		// Test 6: Uppercase multiple words
		// (up,2) should transform the last two words
		// -------------------------------------------------
		{
			name:     "Uppercase multiple",
			input:    "hello world (up,2)",
			expected: "HELLO WORLD",
		},

		//
		// -------------------------------------------------
		// Test 7: Lowercase transformation
		// Ensures words convert to lowercase
		// -------------------------------------------------
		{
			name:     "Lowercase transformation",
			input:    "HELLO (low)",
			expected: "hello",
		},

		//
		// -------------------------------------------------
		// Test 8: Capitalization
		// (cap) should uppercase the first letter
		// -------------------------------------------------
		{
			name:     "Capitalization",
			input:    "hello (cap)",
			expected: "Hello",
		},

		//
		// -------------------------------------------------
		// Test 9: Multiple capitalization
		// (cap,3) should capitalize the last 3 words
		// -------------------------------------------------
		{
			name:     "Capitalize multiple",
			input:    "this is amazing (cap,3)",
			expected: "This Is Amazing",
		},

		//
		// -------------------------------------------------
		// Test 10: Quote spacing
		// Ensures spaces inside quotes are removed
		// -------------------------------------------------
		{
			name:     "Quote spacing",
			input:    "' hello world '",
			expected: "'hello world'",
		},

		//
		// -------------------------------------------------
		// Test 11: Article correction
		// "a apple" should become "an apple"
		// -------------------------------------------------
		{
			name:     "Article correction",
			input:    "a apple",
			expected: "an apple",
		},

		//
		// -------------------------------------------------
		// Test 12: Article capitalization
		// "A apple" should become "An apple"
		// -------------------------------------------------
		{
			name:     "Article capitalization",
			input:    "A apple",
			expected: "An apple",
		},

		//
		// -------------------------------------------------
		// Test 13: Full pipeline test
		// Tests multiple transformations together
		// -------------------------------------------------
		{
			name:     "Full pipeline",
			input:    "this is a apple , and 1e (hex) apples ...",
			expected: "this is an apple, and 30 apples...",
		},
	}

	//
	// =====================================================
	// Run each test case
	// =====================================================
	//
	for _, test := range tests {

		// Execute the function being tested
		result := ProcessText(test.input)

		// Compare result with expected output
		if result != test.expected {

			// Report test failure
			t.Errorf(
				"Test %s failed\nInput: %q\nExpected: %q\nGot: %q",
				test.name,
				test.input,
				test.expected,
				result,
			)
		}
	}
}
