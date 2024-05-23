package anagrams

type chopped struct {
    head string
    tail string
}

func chopLetter(word string) (result []chopped) {
    if len(word) == 0 {
        return
    }
    
    if len(word) == 1 {
        return []chopped{{head: word}}
    }
    
    result = append(result, chopped{string(word[0]), word[1:]})
    
    for i := 1; i < len(word); i++ {
        result = append(result, chopped{string(word[i]), word[:i] + word[i+1:]})
    }    

    return
}