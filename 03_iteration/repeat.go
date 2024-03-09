package iteration

func Repeat(character string, quantity int) string {
	var repeated string
	for i := 0; i < quantity; i++ {
		repeated += character
	}
	return repeated
}
