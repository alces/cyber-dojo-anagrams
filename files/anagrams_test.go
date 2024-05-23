package anagrams

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestAnagramsEmptyString(t *testing.T) {
    assert.Equal(t, []string{}, anagrams(""))
}

func TestAnagramsOneChar(t *testing.T) {
    assert.Equal(t, []string{"a"}, anagrams("a"))
}