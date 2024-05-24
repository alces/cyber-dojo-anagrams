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
