package hamming

func Distance(a, b string) (int, error) {
	answer := 0
	for i, _ := range a {
		if b[i] != a[i] {
			answer += 1
		}
	}
	return answer, nil
}
