package utils

func MergeMaps[Map ~map[K]V, K comparable, V any](maps ...Map) Map {
	merged := make(Map)

	for _, currentMap := range maps {
		for key, value := range currentMap {
			merged[key] = value
		}
	}

	return merged
}
