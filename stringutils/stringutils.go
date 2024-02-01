package stringutils

func Pluralize(l int, singular string, plural string) string {
	if l == 1 {
		return singular
	}
	return plural
}
