package anagrams

type chopped struct {
    head string
    tail string
}

type stringSet map[string]struct{}

func (ss stringSet) add(word string) {
    var value struct{}
    ss[word] = value
}

func (ss stringSet) toSlice() []string {
    result := make([]string, 0, len(ss))
    
    for k := range ss {
        result = append(result, k)
    }
    
    return result
}