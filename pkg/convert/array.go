package convert

import "strconv"

func StringArrayToIntSlice(stringArray []string) []int {
	intArray := make([]int, 0, len(stringArray))

	for _, stringValue := range stringArray {
		intValue, err := strconv.Atoi(stringValue)
		if err != nil {
			continue
		}

		intArray = append(intArray, intValue)
	}
	return intArray
}
