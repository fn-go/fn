package util

func Ptr[T any](s T) *T {
	return &s
}
