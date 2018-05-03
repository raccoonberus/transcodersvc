package util

func InSlice(needle interface{}, haystack []interface{}) bool {
	for _, b := range haystack {
		if b == needle {
			return true
		}
	}
	return false
}