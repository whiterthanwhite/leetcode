package groupanagrams

import "sort"

func groupAnagrams(strs []string) [][]string {
	n := len(strs)
	mapStrs := make(map[string][]string)
	for i := 0; i < n; i++ {
		runeKey := make([]rune, 0)
		for _, r := range strs[i] {
			runeKey = append(runeKey, r)
		}
		sort.Slice(runeKey, func(i, j int) bool {
			return runeKey[i] < runeKey[j]
		})
		key := string(runeKey)
		if _, ok := mapStrs[key]; !ok {
			mapStrs[key] = []string{}
		}
		mapStrs[key] = append(mapStrs[key], strs[i])
	}

	result := make([][]string, len(mapStrs))
	i := 0
	for _, val := range mapStrs {
		result[i] = val
		i++
	}

	return result
}
