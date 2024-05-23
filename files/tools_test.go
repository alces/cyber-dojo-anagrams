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
    assert.Len(t, 2, actual)
    
    for _, e := range expected {
        assert.Contains(t, e, actual)
    }
}