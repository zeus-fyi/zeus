package strings_filter

import "strings"

// FilterStringWithOpts returns true if the word passes the filter
func FilterStringWithOpts(word string, filter *FilterOpts) bool {
	// empty doesNotContain in means ignore, also must have at least one letter
	if filter == nil {
		return true
	}
	for _, check := range filter.DoesNotInclude {
		if len(check) > 0 && strings.Contains(word, check) {
			return false
		}
	}

	if len(filter.StartsWith) > 0 {
		if !strings.HasPrefix(word, filter.StartsWith) {
			return false
		}
	}

	if len(filter.Contains) > 0 {
		if !strings.Contains(word, filter.Contains) {
			return false
		}
	}

	if len(filter.DoesNotStartWithThese) > 0 {
		for _, wordFilter := range filter.DoesNotStartWithThese {
			if strings.HasPrefix(word, wordFilter) {
				return false
			}
		}
	}
	//if len(filter.MustHaveSuffixWithAnyOfThese) > 0 {
	//	if !CheckForSuffixMatch(word, filter.StartsWithAnyOfThese) {
	//		return false
	//	}
	//}
	if len(filter.StartsWithAnyOfThese) > 0 {
		matchFound := false
		for _, wordFilter := range filter.StartsWithAnyOfThese {
			if strings.HasPrefix(word, wordFilter) {
				matchFound = true
			}
		}
		return matchFound
	}
	return true
}

func CheckForSuffixMatch(word string, wordFilterSlice []string) bool {
	for _, wordFilter := range wordFilterSlice {
		if strings.HasSuffix(word, wordFilter) {
			return true
		}
	}
	return false
}

type FilterOpts struct {
	DoesNotStartWithThese []string
	StartsWithAnyOfThese  []string
	StartsWith            string
	Contains              string
	DoesNotInclude        []string
}
