package iteration

func Repeat(char string, freq int) string {
	var buff string
	for range freq {
		buff += char
	}
	return buff
}
