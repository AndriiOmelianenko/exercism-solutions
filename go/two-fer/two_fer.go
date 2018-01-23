// Package twofer returns sentence ""One for X, one for me.", where X is name or "you"
package twofer

func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
