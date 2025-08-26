package core

import (
	"cmp"
	"sort"
	"strings"
)

func Pointer[T any](input T) *T {
	return &input
}

func GetAsPointer[T any](o any) *T {
	if t, ok := o.(T); ok {
		return &t
	}

	if t, ok := o.(*T); ok {
		return t
	}

	return nil
}

// GetImageTag returns the tag of the image. If the image does not have a tag, empty string is returned.
func GetImageTag(image string) string {
	index := strings.LastIndex(image, ":")
	if index == -1 {
		return ""
	}

	return image[index+1:]
}

func SortedKeys[TKey cmp.Ordered, TValue any](m map[TKey]TValue) []TKey {
	keys := make([]TKey, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	return keys
}
