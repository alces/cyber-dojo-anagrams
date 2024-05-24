package anagrams

func anagrams(word string) (result []string) {
    if len(word) == 0 {
        return
    }
    
    if len(word) == 1 {
        return []string{word}
    }
    
    for _, c := range chopLetter(word) {
        for _, t := range(anagrams(c.tail)) {
            result = append(result, c.head + t)
        }
    }    
    
    return
}
