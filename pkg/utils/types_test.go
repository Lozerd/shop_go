package utils

import "testing"

// TestStringOrDefault tests the StringOrDefault function.
func TestStringOrDefault(t *testing.T) {
    // Define test cases
    testCases := []struct {
        name           string
        inputValue     string
        defaultValue   string
        expectedOutput string
    }{
        {
            name:           "Non-empty string",
            inputValue:     "hello",
            defaultValue:   "default",
            expectedOutput: "hello",
        },
        {
            name:           "Empty string with default",
            inputValue:     "",
            defaultValue:   "default",
            expectedOutput: "default",
        },
        {
            name:           "String with spaces only",
            inputValue:     "    ",
            defaultValue:   "default",
            expectedOutput: "default",
        },
        {
            name:           "String with surrounding spaces",
            inputValue:     "   world   ",
            defaultValue:   "default",
            expectedOutput: "world",
        },
    }

    // Iterate through the test cases and run the tests
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result := StringOrDefault(tc.inputValue, tc.defaultValue)
            if result != tc.expectedOutput {
                t.Errorf("StringOrDefault(%q, %q) = %q; want %q", tc.inputValue, tc.defaultValue, result, tc.expectedOutput)
            }
        })
    }
}

