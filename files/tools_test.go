package anagrams

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestChopLetterEmptyString(t *testing.T) {
    assert.Nil(t, chopLetter(""))
}