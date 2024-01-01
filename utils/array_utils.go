package utils

func IsInAny[T comparable](item T, sources ...[]T) bool {
	for _, source := range sources {
		for _, v := range source {
			if v == item {
				return true
			}
		}
	}

	return false
}
