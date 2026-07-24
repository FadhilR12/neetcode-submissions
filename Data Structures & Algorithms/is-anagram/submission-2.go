func isAnagram(s string, t string) bool {
	mapS := map[byte]int{}
	mapT := map[byte]int{}

	for i := 0; i < len(s); i++ {
		mapS[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		mapT[t[i]]++
	}
	
	if len(mapS) != len(mapT){
		return false
	}

	for k, valS := range mapS {
		valT, ok := mapT[k]

		if !ok || valS != valT{
			return false
		}
	}
	return true
}
