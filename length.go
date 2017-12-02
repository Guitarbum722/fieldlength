package fieldlength

import "strings"

// FieldLen is used to determine the length of a field in delimited data.
// sep is used as the delimiter in the data of s.  If it is possible that the data could
// contain the sep character within quotes, etc., then provide the text identifier, aka qualifier,
// to qual.  If there is not expected qualifier, then use "".
func FieldLen(s, sep, qual string) int {
	i := 0
	if qual == "" || !strings.HasPrefix(s, qual) {
		i = strings.Index(s, sep)
	} else {
		i = strings.Index(s, qual+sep)

		if i == -1 {
			return len(s)
		}
		return len(s[:i+len(qual)])
	}

	if i == -1 {
		return len(s)
	}

	return len(s[:i])
}
