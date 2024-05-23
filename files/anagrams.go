package anagrams

func anagrams(word string) (result []string) {
    if len(word) == 0 {
        return
    }
    
    if len(word) == 1 {
        return []string{word}
    }
    
    return
}
