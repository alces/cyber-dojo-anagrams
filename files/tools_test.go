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
        []string{"a"},
        []string{"b"},
        []string{"ab"},
    },
    {
        []string{"a"},
        []string{"b", "c"},
        []string{"ab", "ac"},
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

func TestConcatSlicesEmpty(t *testing.T) {
    assert.Nil(t, concatSlices([]string{}, []string{})) 
}

func TestConcatSlicesOneItem(t *testing.T) {
    assert.Nil(t, concatSlices([]string{"a"}, []string{})) 
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
