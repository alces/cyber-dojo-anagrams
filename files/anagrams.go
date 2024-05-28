package anagrams

func anagrams(word string) []string {
    if len(word) == 1 {
        return []string{word}
    }
    
    result := stringSet{}
    
    for _, c := range chopLetter(word) {
        for _, t := range(anagrams(c.tail)) {
            result.add(c.head + t)
        }
    }    
    
    return result.toSlice()
}

func chopLetter(word string) (result []chopped) {
    for i := 0; i < len(word); i++ {
        result = append(result, chopped{string(word[i]), word[:i] + word[i+1:]})
    }    

    return
}