package strx_test

import (
	"testing"

	"github.com/renatopp/go-x/strx"
	"github.com/renatopp/go-x/testx"
)

func TestLevenshteinDistance(t *testing.T) {
	testx.Equal(t, 0, strx.LevenshteinDistance("", ""))
	testx.Equal(t, 5, strx.LevenshteinDistance("", "hello"))
	testx.Equal(t, 5, strx.LevenshteinDistance("hello", ""))
	testx.Equal(t, 0, strx.LevenshteinDistance("hello", "hello"))
	testx.Equal(t, 3, strx.LevenshteinDistance("kitten", "sitting"))
	testx.Equal(t, 2, strx.LevenshteinDistance("你好世界", "你好"))
}

func TestSimilarity(t *testing.T) {
	testx.Equal(t, float64(1), strx.FuzzySimilarity("", ""))
	testx.Equal(t, float64(1), strx.FuzzySimilarity("hello", "hello"))
	testx.Equal(t, float64(0), strx.FuzzySimilarity("", "hello"))

	sim := strx.FuzzySimilarity("kitten", "sitting")
	if sim <= 0 || sim >= 1 {
		t.Fatalf("expected similarity between 0 and 1, got %v", sim)
	}
}

func TestFuzzyMatch(t *testing.T) {
	testx.Equal(t, true, strx.FuzzyMatch("", "anything"))
	testx.Equal(t, true, strx.FuzzyMatch("gcp", "GoCatPro"))
	testx.Equal(t, true, strx.FuzzyMatch("hello", "Hello"))
	testx.Equal(t, false, strx.FuzzyMatch("xyz", "hello"))
	testx.Equal(t, false, strx.FuzzyMatch("ba", "abc"))
}

func TestFuzzyScore(t *testing.T) {
	testx.Equal(t, 0, strx.FuzzyScore("", "anything"))
	testx.Equal(t, -1, strx.FuzzyScore("xyz", "hello"))

	consecutive := strx.FuzzyScore("hel", "hello")
	scattered := strx.FuzzyScore("hlo", "hello")
	if consecutive <= scattered {
		t.Fatalf("expected consecutive match to score higher: %d vs %d",
			consecutive, scattered)
	}

	boundary := strx.FuzzyScore("gp", "GoPro")
	noBoundary := strx.FuzzyScore("gp", "biggp")
	if boundary <= noBoundary {
		t.Fatalf("expected boundary match to score higher: %d vs %d",
			boundary, noBoundary)
	}
}

func TestBestMatch(t *testing.T) {
	candidates := []string{"apple", "application", "banana", "grape"}

	best, ok := strx.FuzzyBestMatch("app", candidates)
	testx.Equal(t, true, ok)
	testx.Equal(t, "apple", best)

	_, ok = strx.FuzzyBestMatch("zzz", candidates)
	testx.Equal(t, false, ok)

	_, ok = strx.FuzzyBestMatch("app", []string{})
	testx.Equal(t, false, ok)
}

func TestBestMatchFunc(t *testing.T) {
	type item struct{ Name string }
	items := []item{{"apple"}, {"application"}, {"banana"}}

	best, ok := strx.FuzzyBestMatchFunc("app", items, func(i item) string {
		return i.Name
	})
	testx.Equal(t, true, ok)
	testx.Equal(t, "apple", best.Name)

	_, ok = strx.FuzzyBestMatchFunc("zzz", items, func(i item) string {
		return i.Name
	})
	testx.Equal(t, false, ok)
}

func TestSortByMatch(t *testing.T) {
	candidates := []string{"banana", "app", "apple", "application", "grape"}
	sorted := strx.FuzzySortByMatch("app", candidates)

	testx.Equal(t, "app", sorted[0])
	testx.Equal(t, "apple", sorted[1])
	testx.Equal(t, "application", sorted[2])
	testx.Equal(t, "banana", sorted[3])
	testx.Equal(t, "grape", sorted[4])
}
