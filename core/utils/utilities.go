package utils

func SafeStr(s string) string {
	if s == "" {
		return "(not specified)"
	}
	return s
}

func SafeStrPtr[T any](ptr *T, getter func(*T) string) string {
	if ptr == nil {
		return "(not specified)"
	}
	return getter(ptr)
}

func SafeSlice[T any](s []T) []T {
	if s == nil {
		return []T{}
	}
	return s
}
