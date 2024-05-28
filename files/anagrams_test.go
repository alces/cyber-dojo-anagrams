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
    {"abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
    {"abcd", []string{"abcd", "abdc", "acbd", "acdb", "adbc", "adcb",
                      "bacd", "badc", "bcad", "bcda", "bdac", "bdca",
                      "cabd", "cadb", "cbad", "cbda", "cdab", "cdba",
                      "dacb", "dabc", "dbac", "dbca", "dcab", "dcba"},
    },
    {"abba", []string{"aabb", "abba", "abab", "baba", "baab", "bbaa"},},
}

var chopLetterTestResults = []struct {
    argument string
    expected []chopped
} {
    {
        argument: "",
        expected: nil,
    },
    {
        argument: "a",
        expected: []chopped{
            {"a", ""},
        },
    },
    {
        argument: "ab",
        expected: []chopped{
            {"a", "b"},
            {"b", "a"},
        },
    },
    {
        argument: "abc",
        expected: []chopped{
            {"a", "bc"},
            {"b", "ac"},
            {"c", "ab"},
        },
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

func TestChopLetter(t *testing.T) {
    for _, res := range chopLetterTestResults {
        actual := chopLetter(res.argument)
        assert.Len(t, actual, len(res.expected))
    
        for _, e := range res.expected {
            assert.Contains(t, actual, e)
        }
    }
}
