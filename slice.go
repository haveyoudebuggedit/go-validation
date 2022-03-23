package validation

// InSlice returns true if the element specified is contained in the slice.
func InSlice[K comparable](element K, slice []K) bool {
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
}

// NotInSlice returns true if the element specified is not contained in the slice.
func NotInSlice[K comparable](element K, slice []K) bool {
	return !InSlice(element, slice)
}
