package authcore

func HasPermission(perms []string, permission string) bool {
	for _, p := range perms {
		if p == permission {
			return true
		}
	}
	return false
}
