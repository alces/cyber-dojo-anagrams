package anagrams

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestStringSetAdd(t *testing.T) {
    ss := stringSet{}
    ss.add("abc")
    
    assert.Len(t, ss, 1)
}