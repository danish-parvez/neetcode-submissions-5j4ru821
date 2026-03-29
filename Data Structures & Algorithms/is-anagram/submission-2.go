func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    countS, countT := make(map[rune]int), make(map[rune]int)
    for _, c := range s {
            countS[c] += 1
    }
    for _, c := range t {
        countT[c] += 1
    }
    for k := range countS {
        if countS[k] != countT[k] {
            return false
        }
    }
    return true
}