package anagrams

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

var anagramsTestResults = []struct {
    argument string
    expected []string
} {
    {"", nil},
    {"a", []string{"a"}},
    {"ab", []string{"ab", "ba"}},
    {"abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
    {"abcd", []string{"abcd", "abdc", "acbd", "acdb", "adbc", "adcb",
                      "bacd", "badc", "bcad", "bcda", "bdac", "bdca",
                      "cabd", "cadb", "cbad", "cbda", "cdab", "cdba",
                      "dacb", "dabc", "dbac", "dbca", "dcab", "dcba"},
    },
}

func TestAnagrams(t *testing.T) {
    for _, res := range anagramsTestResults {
        actual := anagrams(res.argument)
        assert.Len(t, actual, len(res.expected))
        
        for _, e := range res.expected {
            assert.Contains(t, actual, e)
        }
    }
}
