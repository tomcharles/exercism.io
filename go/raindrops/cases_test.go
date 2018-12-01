package raindrops

// Source: exercism/problem-specifications
// Commit: 99de15d raindrops: apply "input" policy
// Problem Specifications Version: 1.1.0

var tests = []struct {
	input    int
	expected string
}{
	{15, "PlingPlang"},
	{21, "PlingPlong"},
	{25, "Plang"},
	{27, "Pling"},
	{35, "PlangPlong"},
	{49, "Plong"},
	{52, "52"},
	{105, "PlingPlangPlong"},
	{3125, "Plang"},
}
