package anagrams

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestChopLetterEmptyString(t *testing.T) {
    assert.Nil(t, chopLetter(""))
}

func TestChopLetterOneChar(t *testing.T) {
    expected := []chopped{
        {
            head: "a",
            tail: "",
        },
    }
    
    assert.Equal(t, expected, chopLetter("a"))
}

func TestChopLetterTwoChars(t *testing.T) {
        expected := []chopped{
        {
            head: "a",
            tail: "b",
        },
        {
            head: "b",
            tail: "a",
        },
    }
    
    actual := chopLetter("ab")
    assert.Len(t, actual, 2)
    
    for _, e := range expected {
        assert.Contains(t, actual, e)
    }
}

func TestChopLetterThreeChars(t *testing.T) {
    expected := []chopped{
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
    }
    
    actual := chopLetter("abc")
    assert.Len(t, actual, 3)
    
    for _, e := range expected {
        assert.Contains(t, actual, e)
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

