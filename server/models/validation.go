package models

func IsValidJobTitle(title string) bool {
	for _, t := range AllowedJobTitles {
		if t == title {
			return true
		}
	}
	return false
}

func IsValidLocation(location string) bool {
	for _, l := range AllowedLocations {
		if l == location {
			return true
		}
	}
	return false
}
