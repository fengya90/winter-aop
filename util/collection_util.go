package util

func RemoveDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, ok := allKeys[item]; !ok {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
