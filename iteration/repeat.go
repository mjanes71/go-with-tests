package iteration

func Repeat(a string, f int) (repeated string) {
	for i := 0; i < f; i++ {
		repeated += a
	}
	return repeated
}