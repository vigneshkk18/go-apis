package utils

func GroupBy(year int, month int, week int, day int) string {
	if year != -1 && month != -1 && week != -1 && day != -1 {
		return "GROUP_BY_DAY"
	} else if year != -1 && month != -1 && week != -1 {
		return "GROUP_BY_WEEK"
	} else if year != -1 && month != -1 {
		return "GROUP_BY_MONTH"
	} else if year != -1 {
		return "GROUP_BY_YEAR"
	}
	return "ALL"
}
