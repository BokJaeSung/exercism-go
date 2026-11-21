package hamming

import "errors"

func Distance(a, b string) (int, error) {
	answer := 0
	if len(a) != len(b) {
		e := errors.New("문자열 길이가 다름")
		return 0, e
	}
	for i, _ := range a {
		if b[i] != a[i] {
			answer += 1
		}
	}
	return answer, nil
}
