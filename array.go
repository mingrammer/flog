package main

func ContainsString(arr []string, str string) bool {
	for _, s := range arr {
		if str == s {
			return true
		}
	}
	return false
}
