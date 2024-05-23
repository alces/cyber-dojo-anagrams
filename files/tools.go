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
        return []chopped{
            {
                head: word,
            },
        }
    }

    return
}