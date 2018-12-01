package reverse

func String(s string) string {
	reversed := []rune{}
	for _, char := range s {
		reversed = append([]rune{char}, reversed...)
	}
	return string(reversed)
}
