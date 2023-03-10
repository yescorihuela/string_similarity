package string_similarity

func Similarity(s1, s2 string) float32 {
	p1 := pairs(s1)
	p2 := pairs(s2)

	numerator := float32(2 * len(intersect(p1, p2)))
	denominator := float32(len(p1) + len(p2))
	if denominator == 0 {
		denominator = 1
	}

	return (numerator / denominator) * 100

}

func pairs(w string) []string {
	pair := make([]string, 0)
	word := w
	for {
		if len(word) > 0 {
			if len(word) == 1 {
				pair = append(pair, string(word[0]))
			} else {
				w := ""
				for _, v := range word[:2] {
					w += string(v)
				}
				pair = append(pair, w)
			}
			word = word[1:]
		} else {
			break
		}
	}
	return pair
}

func intersect(a, b []string) []string {
	m := make(map[string]struct{})
	result := make([]string, 0)
	for _, v := range a {
		m[v] = struct{}{}
	}
	for _, v := range b {
		if _, ok := m[v]; ok {
			result = append(result, v)
		}
	}
	return result
}
