package helpers

// HasDirectDiff returns true, if mere one element from "it" is in "in"
func HasDirectDiff(it []string, in []string) bool {
	for _, s := range it {
		for _, s2 := range in {
			if s == s2 {
				goto here
			}
		}
		return true
	here:
	}

	return false
}
