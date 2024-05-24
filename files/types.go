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

func (ss stringSet) keys() []string {
    return []string{}
}