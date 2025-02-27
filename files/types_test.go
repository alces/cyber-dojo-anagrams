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

func TestStringSetToSlice(t *testing.T) {
    ss := stringSet{}
    ss.add("abc")
    ss.add("bcd")
    
    keys := ss.toSlice()
    assert.Len(t, keys, len(ss))
    assert.Contains(t, keys, "abc")
    assert.Contains(t, keys, "bcd")
}