func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    s_hashmap := make(map[rune]int)
    t_hashmap := make(map[rune]int)
    for _, c := range s {
        if _, ok := s_hashmap[c]; !ok {
            s_hashmap[c] = 1
        } else {
            s_hashmap[c] += 1
        }
    }
    for _, c := range t {
        if _, ok := t_hashmap[c]; !ok {
            t_hashmap[c] = 1
        } else {
            t_hashmap[c] += 1
        }
    }
    for k := range s_hashmap {
        if s_hashmap[k] != t_hashmap[k] {
            return false
        }
    }
    return true
}
