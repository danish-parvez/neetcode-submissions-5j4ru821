func longestCommonPrefix(strs []string) string {
    var commonPrefix string
	i := 0
	for i < len(strs[0]) {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || c != strs[j][i] {
				return commonPrefix
			}
		}
		commonPrefix += string(c)
		i++
	}
	return commonPrefix
}