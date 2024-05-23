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
        argument: "a",
        expected: {
            {
                head: "a",
                tail: "",
            },
        },
    },
    {
        argument: "ab",
        expected: {
            {
                head: "a",
                tail: "b",
            },
            {
                head: "b",
                tail: "a",
            },
        },
    },
    {
        argument: "abc",
        expected: {
            {
                head: "a",
                tail: "bc",
            },
            {
                head: "b",
                tail: "ac",
            },
            {
                head: "c",
                tail: "ab",
            },
        },
    },
}
func TestChopLetterEmptyString(t *testing.T) {
    assert.Nil(t, chopLetter(""))
}

func TestChopLetterThreeChars(t *testing.T) {
    for _, res := range chopLetterTestResults {
        actual := chopLetter(res.argument)
        assert.Len(t, actual, len(res.expected))
    
        for _, e := range res.expected {
            assert.Contains(t, res.actual, e)
        }
    }
}

func TestConcatSlicesEmpty(t *testing.T) {
    assert.Nil(t, concatSlices([]string{}, []string{})) 
}

func TestConcatSlicesOneItem(t *testing.T) {
    assert.Nil(t, concatSlices([]string{"a"}, []string{})) 
}

func TestConcatSlicesTwoItems(t *testing.T) {
    expected := []string{"ab"}
    actual := concatSlices([]string{"a"}, []string{"b"})
    
    assert.Len(t, actual, len(expected))
    for _, e := range expected {
        assert.Contains(t, actual, e)
    }
}

func TestConcatSlicesThreeItems(t *testing.T) {
    expected := []string{"ab", "ac"}
    actual := concatSlices([]string{"a"}, []string{"b", "c"})
    
    assert.Len(t, actual, len(expected))
    for _, e := range expected {
        assert.Contains(t, actual, e)
    }
}

