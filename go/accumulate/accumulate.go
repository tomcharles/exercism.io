package accumulate

// Accumulate applies the [converter] function to each element in the input slice
func Accumulate(in []string, converter func(string) string) (out []string) {
	for _, str := range in {
		out = append(out, converter(str))
	}
	return
}
