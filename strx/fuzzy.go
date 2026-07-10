package strx

import (
	"sort"
	"strings"
	"unicode"
)

// LevenshteinDistance returns the minimum number of single-rune
// insertions, deletions or substitutions required to turn a into b.
// It is rune-aware, so multi-byte characters are treated as a single
// unit. For example, LevenshteinDistance("kitten", "sitting")
// returns 3.
func LevenshteinDistance(a, b string) int {
	ra := []rune(a)
	rb := []rune(b)

	if len(ra) == 0 {
		return len(rb)
	}
	if len(rb) == 0 {
		return len(ra)
	}

	prev := make([]int, len(rb)+1)
	curr := make([]int, len(rb)+1)
	for j := range prev {
		prev[j] = j
	}

	for i := 1; i <= len(ra); i++ {
		curr[0] = i
		for j := 1; j <= len(rb); j++ {
			cost := 1
			if ra[i-1] == rb[j-1] {
				cost = 0
			}
			del := prev[j] + 1
			ins := curr[j-1] + 1
			sub := prev[j-1] + cost
			curr[j] = min(del, min(ins, sub))
		}
		prev, curr = curr, prev
	}

	return prev[len(rb)]
}

// FuzzySimilarity returns a value between 0 and 1 that represents how
// similar a and b are, based on the Levenshtein distance normalized
// by the length of the longest string. A result of 1 means the
// strings are equal, and 0 means they share no similarity. If both
// strings are empty, it returns 1.
func FuzzySimilarity(a, b string) float64 {
	la := Length(a)
	lb := Length(b)
	if la == 0 && lb == 0 {
		return 1
	}

	maxLen := max(lb, la)

	dist := LevenshteinDistance(a, b)
	return 1 - float64(dist)/float64(maxLen)
}

// FuzzyMatch returns true if every rune of pattern appears in s in
// the same order, ignoring case. It does not require the runes to be
// contiguous. For example, FuzzyMatch("gcp", "GoCatPro") returns
// true, since 'g', 'c' and 'p' appear in that order.
func FuzzyMatch(pattern, s string) bool {
	return FuzzyScore(pattern, s) >= 0
}

// FuzzyScore returns a score for how well pattern fuzzy-matches s. It
// returns -1 if pattern is not a subsequence of s. Higher scores
// represent better matches: consecutive matches and matches at word
// boundaries (start of string, after a separator, or at a
// lower-to-upper case transition) are weighted more. The comparison
// is case-insensitive.
func FuzzyScore(pattern, s string) int {
	pRunes := []rune(strings.ToLower(pattern))
	if len(pRunes) == 0 {
		return 0
	}

	sRunes := []rune(s)
	sLower := []rune(strings.ToLower(s))

	pi := 0
	score := 0
	consecutive := 0

	for i := 0; i < len(sLower) && pi < len(pRunes); i++ {
		if sLower[i] != pRunes[pi] {
			consecutive = 0
			continue
		}

		score++
		if consecutive > 0 {
			score += 2
		}
		if isWordBoundary(sRunes, i) {
			score += 3
		}
		consecutive++
		pi++
	}

	if pi < len(pRunes) {
		return -1
	}
	return score
}

// FuzzyBestMatch returns the candidate with the highest FuzzyScore against
// pattern. The second return value is false if no candidate matches
// pattern at all, in which case the first return value is the zero
// value.
func FuzzyBestMatch(pattern string, candidates []string) (string, bool) {
	return FuzzyBestMatchFunc(pattern, candidates, func(s string) string {
		return s
	})
}

// FuzzyBestMatchFunc is like BestMatch, but it takes a slice of items of
// any type and a function that extracts the string to match against
// from each item.
func FuzzyBestMatchFunc[T any](pattern string, items []T, fn func(T) string) (T, bool) {
	var best T
	bestScore := -1
	found := false

	for _, item := range items {
		score := FuzzyScore(pattern, fn(item))
		if score < 0 {
			continue
		}
		if !found || score > bestScore {
			best = item
			bestScore = score
			found = true
		}
	}

	return best, found
}

// FuzzySortByMatch returns a new slice with the candidates that match
// pattern sorted by descending FuzzyScore. Candidates that do not
// match pattern are placed at the end, in their original relative
// order. The input slice is not modified.
func FuzzySortByMatch(pattern string, candidates []string) []string {
	type scored struct {
		value string
		score int
	}

	items := make([]scored, len(candidates))
	for i, c := range candidates {
		items[i] = scored{value: c, score: FuzzyScore(pattern, c)}
	}

	sort.SliceStable(items, func(i, j int) bool {
		si, sj := items[i].score, items[j].score
		if si < 0 && sj < 0 {
			return false
		}
		if si < 0 {
			return false
		}
		if sj < 0 {
			return true
		}
		return si > sj
	})

	res := make([]string, len(items))
	for i, it := range items {
		res[i] = it.value
	}
	return res
}

// isWordBoundary returns true if the rune at index i in s starts a
// new word: it is the first rune, follows a separator, or follows a
// lowercase rune while being uppercase itself.
func isWordBoundary(s []rune, i int) bool {
	if i == 0 {
		return true
	}

	switch s[i-1] {
	case ' ', '-', '_', '.', '/', '\\':
		return true
	}

	return unicode.IsUpper(s[i]) && unicode.IsLower(s[i-1])
}
