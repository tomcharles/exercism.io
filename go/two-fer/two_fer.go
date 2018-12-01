// Package twofer implements a function to share things
package twofer

import "fmt"

// ShareWith shares something with (name) and itself.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
