func groupAnagrams(strs []string) [][]string {
    res := make(map[string][]string)
    for _, str := range strs {
        sortedStr := sortString(str)
        res[sortedStr] = append(res[sortedStr], str)
    }
    var result [][]string
    for _, group := range res {
        result = append(result, group)
    }
    return result
}

func sortString(str string) string {
    characters := []rune(str)
    sort.Slice(characters, func(i, j int) bool {
        return characters[i] < characters[j]
    })
    return string(characters)
}
