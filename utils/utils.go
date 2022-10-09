package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func MapKeys[Key comparable, Val any](mapObj map[Key]Val) []Key {
	s := make([]Key, 0, len(mapObj))
	for k := range mapObj {
		s = append(s, k)
	}
	return s
}

func TimeTaken(t string) float64 {
	strSplit := strings.Split(t, ":")
	min, _ := strconv.ParseFloat(strSplit[0], 64)
	sec, _ := strconv.ParseFloat(strSplit[1], 64)
	val, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", min+(sec/60)), 64)
	return val
}

func Date(year int, month int, day int) time.Time {
	return time.Date(year, time.Month(month), day, 1, 1, 1, 1, time.UTC)
}

func GetStartDateOfWeek(week int) int {
	return (week * 7) + 1
}
