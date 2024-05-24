package anagrams

type chopped struct {
    head string
    tail string
}

func anagrams(word string) (result []string) {
    for _, c := range chopLetter(word) {
        for _, t := range(anagrams(c.tail)) {
            result = append(result, c.head + t)
        }
    }    
    
    return
}

func chopLetter(word string) (result []chopped) { 
    for i := 0; i < len(word); i++ {
        result = append(result, chopped{string(word[i]), word[:i] + word[i+1:]})
    }    

    return
}