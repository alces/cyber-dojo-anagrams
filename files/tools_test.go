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