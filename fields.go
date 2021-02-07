package qualified

import "strings"

// SplitWithQual basically works like the standard strings.Split() func, but will consider a text qualifier if set.
func SplitWithQual(s, sep, qual string) []string {
	if qual == "" {
		return strings.Split(s, sep) // use standard Split() method if no qualifier is considered
	}
	words := make([]string, 0, strings.Count(s, sep))

	for start := 0; start < len(s); {
		count := length(s[start:], sep, qual)
		words = append(words, s[start:start+count])
		start += count + len(sep)
	}
	return words
}

// FieldLengths returns a map that contains a key representing the column position
// and a value which represents the count for the respective column.
// The counts are based on fields delimited by sep.  If s contains a character
// that matches sep in the data itself, provide that character for qual
// so that it can be properly escaped.  Otherwise provide and empty string for qual.
func FieldLengths(s, sep, qual string) map[int]int {
	var columnNum int
	var temp int
	// count per field
	counts := make(map[int]int)

	for start := 0; start < len(s); {
		temp = length(s[start:], sep, qual)
		start += temp + len(sep)
		if temp > counts[columnNum] {
			counts[columnNum] = temp
		}
		columnNum++
		temp = 0
	}
	return counts
}

// length is used to determine the length of a field in delimited data.
// sep is used as the delimiter in the data of s.  If it is possible that the data could
// contain the sep character within quotes, etc., then provide the text identifier, aka qualifier,
// to qual.  If there is not expected qualifier, then use "".
func length(s, sep, qual string) int {
	var endIdx int
	if len(qual) > 0 && strings.HasPrefix(s, qual) {
		endIdx += len(qual)

		endIdx += strings.Index(s[endIdx:], qual) + len(qual)

		return len(s[:endIdx])
	}

	endIdx += strings.Index(s, sep)

	if endIdx == -1 {
		// last field
		return len(s)
	}

	return len(s[:endIdx])
}
