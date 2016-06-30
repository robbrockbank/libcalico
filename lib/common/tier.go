package common

const (
	DefaultTierName = ".default"
	blank           = ""
)

// Return the tier name, or the default if blank.
func tierOrDefault(tier string) string {
	if tier == blank {
		return DefaultTierName
	} else {
		return tier
	}
}

// Return the tier name, or blank if the default.
func tierOrBlank(tier string) string {
	if tier == DefaultTierName {
		return blank
	} else {
		return tier
	}
}
