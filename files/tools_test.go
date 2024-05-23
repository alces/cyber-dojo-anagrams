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
