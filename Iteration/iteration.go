package iteration

func Repeat(s string, a int) string {
	var repeated string
	for i := 0; i < a; i++ {
		repeated += s
	}

	return repeated

}
