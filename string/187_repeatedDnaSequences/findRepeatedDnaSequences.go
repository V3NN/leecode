package __repeatedDnaSequences

func findRepeatedDnaSequences(s string) []string {

	ans := []string{}
	hash := map[string]int{}

	for i := 0; i <= len(s)-10; i++ {
		sub := s[i : i+10]
		hash[sub]++
		if hash[sub] == 2 {
			ans = append(ans, sub)
		}
	}

	return ans
}
