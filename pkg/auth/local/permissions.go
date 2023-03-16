package local

func parseLocalPermissions(perms []string) []string {
	res := make([]string, 0, len(perms))
	for i := range perms {
		res = append(res, perms[i])
	}

	return res
}
