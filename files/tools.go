package anagrams

type chopped struct {
    head string
    tail string
}

func chopLetter(word string) (result []chopped) { 
    for i := 0; i < len(word); i++ {
        result = append(result, chopped{string(word[i]), word[:i] + word[i+1:]})
    }    

    return
}

func concatSlices(first, second []string) (result []string) {
    for _, f := range first {
        for _, s := range second {
            result = append(result, f+s)
        }
    }
   
    return
}