package main

// ContainsString checks if a string array contains a given string
func ContainsString(arr []string, str string) bool {
	for _, s := range arr {
		if str == s {
			return true
		}
	}
	return false
}
