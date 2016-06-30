package backend

const (
	defaultTierName = ".default"
	wildcard        = "*"
	blank           = ""
)

// Return the ID, or a wildcard if the ID is blank.
func idOrWildcard(id string) string {
	if id != "" {
		return id
	} else {
		return wildcard
	}
}

// Return the tier name, or the default if blank.
func tierOrDefault(tier string) string {
	if tier == blank {
		return defaultTierName
	} else {
		return tier
	}
}

// Return the tier name, or blank if the default.
func tierOrBlank(tier string) string {
	if tier == defaultTierName {
		return blank
	} else {
		return tier
	}
}
