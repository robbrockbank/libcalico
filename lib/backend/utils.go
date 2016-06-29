package backend

func idOrWildcard(id string) string {
	if id != "" {
		return id
	} else {
		return "*"
	}
}