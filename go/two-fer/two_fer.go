// Package twofer provides a tool to share cookies with someone else.
package twofer

// ShareWith returns a string with a message after sharing cookies with someone else.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return "One for " + name + ", one for me."
}
