package utils

func MapKeys[Key comparable, Val any](mapObj map[Key]Val) []Key {
	s := make([]Key, 0, len(mapObj))
	for k := range mapObj {
		s = append(s, k)
	}
	return s
}
