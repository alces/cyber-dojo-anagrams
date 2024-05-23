package anagrams

type chopped struct {
    head string
    tail string
}

func chopLetter(word string) (result []chopped) {
    if len(word) == 0 {
        return
    }

    for i := 0; i < len(word); i++ {
        result = append(result, chopped{string(word[i]), word[:i] + word[i+1:]})
    }    

    return
}