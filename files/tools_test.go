package anagrams

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

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

var concatSlicesTestResults = []struct {
    first    []string
    second   []string
    expected []string
} {
    {
        nil,
        nil,
        nil,
    },
    {
        []string{"a"},
        nil,
        nil,
    },
    {
        []string{"a"},
        []string{"b", "c"},
        []string{"ab", "ac"},
    },
    {
        []string{"a", "b"},
        []string{"c", "d"},
        []string{"ac", "ad", "bc", "bd"},
    },
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

func TestConcatSlices(t *testing.T) {
    for _, res := range concatSlicesTestResults {
        actual := concatSlices(res.first, res.second)
        assert.Len(t, actual, len(res.expected))
        
        for _, e := range res.expected {
            assert.Contains(t, actual, e)
        }
    }
}
