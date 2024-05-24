package anagrams

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestStringSetAdd(t *testing.T) {
    ss := stringSet{}
    ss.add("abc")
    
    assert.Len(t, ss, 1)
    assert.Contains(t, ss, "abc")
}

func TestStringSetKeys(t *testing.T) {
    ss := stringSet{}
    ss.add("abc")
    ss.add("bcd")
    
    keys := ss.keys()
    assert.Contains(t, keys, "abc")
}