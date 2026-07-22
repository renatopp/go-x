package iterx

import (
	"slices"
	"testing"

	"github.com/renatopp/go-x/testx"
)

//-----------------------------------------------------------------------------
// WINDOW
//-----------------------------------------------------------------------------

func TestWindow(t *testing.T) {
	var got [][]int
	for w := range Window([]int{1, 2, 3, 4}, 2) {
		got = append(got, slices.Clone(w))
	}
	testx.Equal(t, 3, len(got))
	testx.True(t, slices.Equal(got[0], []int{1, 2}))
	testx.True(t, slices.Equal(got[1], []int{2, 3}))
	testx.True(t, slices.Equal(got[2], []int{3, 4}))
}

func TestWindowSizeGreaterThanSeq(t *testing.T) {
	var got [][]int
	for w := range Window([]int{1, 2}, 5) {
		got = append(got, w)
	}
	testx.Equal(t, 0, len(got))
}

func TestWindowBreak(t *testing.T) {
	var got [][]int
	for w := range Window([]int{1, 2, 3, 4}, 2) {
		got = append(got, slices.Clone(w))
		break
	}
	testx.Equal(t, 1, len(got))
}

func TestWindowPanicsOnNonPositiveSize(t *testing.T) {
	defer func() {
		testx.NotNil(t, recover())
	}()
	for range Window([]int{1, 2, 3}, 0) {
	}
}

func TestWindowString(t *testing.T) {
	var got []string
	for w := range WindowString("abcd", 2) {
		got = append(got, w)
	}
	testx.True(t, slices.Equal(got, []string{"ab", "bc", "cd"}))
}

func TestWindowStringSizeGreaterThanSeq(t *testing.T) {
	var got []string
	for w := range WindowString("ab", 5) {
		got = append(got, w)
	}
	testx.Equal(t, 0, len(got))
}

func TestWindowStringPanicsOnNonPositiveSize(t *testing.T) {
	defer func() {
		testx.NotNil(t, recover())
	}()
	for range WindowString("abc", 0) {
	}
}

//-----------------------------------------------------------------------------
// SHUFFLE
//-----------------------------------------------------------------------------

func TestShuffle(t *testing.T) {
	original := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var got []int
	for v := range Shuffle(original) {
		got = append(got, v)
	}
	testx.Equal(t, len(original), len(got))
	sortedGot := slices.Clone(got)
	slices.Sort(sortedGot)
	sortedOriginal := slices.Clone(original)
	slices.Sort(sortedOriginal)
	testx.True(t, slices.Equal(sortedGot, sortedOriginal))
}

func TestShuffleBreak(t *testing.T) {
	count := 0
	for range Shuffle([]int{1, 2, 3, 4}) {
		count++
		break
	}
	testx.Equal(t, 1, count)
}

//-----------------------------------------------------------------------------
// RANGE
//-----------------------------------------------------------------------------

func TestRangeInt(t *testing.T) {
	var got []int
	for v := range RangeInt(0, 5) {
		got = append(got, v)
	}
	testx.True(t, slices.Equal(got, []int{0, 1, 2, 3, 4}))
}

func TestRangeIntEmpty(t *testing.T) {
	var got []int
	for v := range RangeInt(5, 5) {
		got = append(got, v)
	}
	testx.Equal(t, 0, len(got))
}

func TestRangeIntBreak(t *testing.T) {
	var got []int
	for v := range RangeInt(0, 100) {
		got = append(got, v)
		if v == 2 {
			break
		}
	}
	testx.True(t, slices.Equal(got, []int{0, 1, 2}))
}

func TestRangeIntStep(t *testing.T) {
	var got []int
	for v := range RangeIntStep(0, 10, 2) {
		got = append(got, v)
	}
	testx.True(t, slices.Equal(got, []int{0, 2, 4, 6, 8}))
}

func TestRangeIntStepNegative(t *testing.T) {
	var got []int
	for v := range RangeIntStep(10, 0, -2) {
		got = append(got, v)
	}
	testx.True(t, slices.Equal(got, []int{10, 8, 6, 4, 2}))
}

func TestRangeIntStepPanicsOnZeroStep(t *testing.T) {
	defer func() {
		testx.NotNil(t, recover())
	}()
	for range RangeIntStep(0, 10, 0) {
	}
}

func TestRangeFloat(t *testing.T) {
	var got []float64
	for v := range RangeFloat(0.0, 3.0) {
		got = append(got, v)
	}
	testx.True(t, slices.Equal(got, []float64{0.0, 1.0, 2.0}))
}

func TestRangeFloatStep(t *testing.T) {
	var got []float64
	for v := range RangeFloatStep(0.0, 1.0, 0.2) {
		got = append(got, v)
	}
	testx.Equal(t, 5, len(got))
	testx.AlmostEqual(t, 0.0, got[0], 1e-9)
	testx.AlmostEqual(t, 0.2, got[1], 1e-9)
	testx.AlmostEqual(t, 0.4, got[2], 1e-9)
	testx.AlmostEqual(t, 0.6, got[3], 1e-9)
	testx.AlmostEqual(t, 0.8, got[4], 1e-9)
}

func TestRangeFloatStepNegative(t *testing.T) {
	// Uses -0.25 (an exact binary fraction) to avoid floating-point drift
	// that occurs with non-power-of-two steps like -0.2.
	var got []float64
	for v := range RangeFloatStep(1.0, 0.0, -0.25) {
		got = append(got, v)
	}
	testx.True(t, slices.Equal(got, []float64{1.0, 0.75, 0.5, 0.25}))
}

func TestRangeFloatStepPanicsOnZeroStep(t *testing.T) {
	defer func() {
		testx.NotNil(t, recover())
	}()
	for range RangeFloatStep(0.0, 1.0, 0.0) {
	}
}

//-----------------------------------------------------------------------------
// REPEAT
//-----------------------------------------------------------------------------

func TestRepeat(t *testing.T) {
	count := 0
	for v := range Repeat("x") {
		testx.Equal(t, "x", v)
		count++
		if count == 5 {
			break
		}
	}
	testx.Equal(t, 5, count)
}

func TestRepeatN(t *testing.T) {
	var got []int
	for v := range RepeatN(7, 3) {
		got = append(got, v)
	}
	testx.True(t, slices.Equal(got, []int{7, 7, 7}))
}

func TestRepeatNZeroOrNegative(t *testing.T) {
	count := 0
	for range RepeatN(7, 0) {
		count++
	}
	testx.Equal(t, 0, count)

	count = 0
	for range RepeatN(7, -1) {
		count++
	}
	testx.Equal(t, 0, count)
}

func TestRepeatWhile(t *testing.T) {
	i := 0
	var got []int
	for v := range RepeatWhile(9, func() bool {
		i++
		return i <= 3
	}) {
		got = append(got, v)
	}
	testx.True(t, slices.Equal(got, []int{9, 9, 9}))
}

func TestRepeatUntil(t *testing.T) {
	i := 0
	var got []int
	for v := range RepeatUntil(9, func() bool {
		i++
		return i > 3
	}) {
		got = append(got, v)
	}
	testx.True(t, slices.Equal(got, []int{9, 9, 9}))
}

//-----------------------------------------------------------------------------
// CYCLE
//-----------------------------------------------------------------------------

func TestCycle(t *testing.T) {
	var got []int
	count := 0
	for v := range Cycle([]int{1, 2, 3}) {
		got = append(got, v)
		count++
		if count == 7 {
			break
		}
	}
	testx.True(t, slices.Equal(got, []int{1, 2, 3, 1, 2, 3, 1}))
}

func TestCycleN(t *testing.T) {
	var got []int
	for v := range CycleN([]int{1, 2, 3}, 2) {
		got = append(got, v)
	}
	testx.True(t, slices.Equal(got, []int{1, 2, 3, 1, 2, 3}))
}

func TestCycleNZeroOrNegative(t *testing.T) {
	count := 0
	for range CycleN([]int{1, 2, 3}, 0) {
		count++
	}
	testx.Equal(t, 0, count)
}

func TestCycleWhile(t *testing.T) {
	i := 0
	var got []int
	for v := range CycleWhile([]int{1, 2}, func() bool {
		i++
		return i <= 2
	}) {
		got = append(got, v)
	}
	testx.True(t, slices.Equal(got, []int{1, 2, 1, 2}))
}

func TestCycleUntil(t *testing.T) {
	i := 0
	var got []int
	for v := range CycleUntil([]int{1, 2}, func() bool {
		i++
		return i > 2
	}) {
		got = append(got, v)
	}
	testx.True(t, slices.Equal(got, []int{1, 2, 1, 2}))
}

//-----------------------------------------------------------------------------
// TAKE
//-----------------------------------------------------------------------------

func TestTake(t *testing.T) {
	var got []int
	for v := range Take([]int{1, 2, 3, 4, 5}, 3) {
		got = append(got, v)
	}
	testx.True(t, slices.Equal(got, []int{1, 2, 3}))
}

func TestTakeZeroOrNegative(t *testing.T) {
	count := 0
	for range Take([]int{1, 2, 3}, 0) {
		count++
	}
	testx.Equal(t, 0, count)
}

func TestTakeMoreThanLength(t *testing.T) {
	var got []int
	for v := range Take([]int{1, 2, 3}, 10) {
		got = append(got, v)
	}
	testx.True(t, slices.Equal(got, []int{1, 2, 3}))
}

func TestTakeWhile(t *testing.T) {
	var got []int
	for v := range TakeWhile([]int{1, 2, 3, 4, 1}, func(v int) bool { return v < 4 }) {
		got = append(got, v)
	}
	testx.True(t, slices.Equal(got, []int{1, 2, 3}))
}

func TestTakeUntil(t *testing.T) {
	var got []int
	for v := range TakeUntil([]int{1, 2, 3, 4, 1}, func(v int) bool { return v == 4 }) {
		got = append(got, v)
	}
	testx.True(t, slices.Equal(got, []int{1, 2, 3}))
}

//-----------------------------------------------------------------------------
// ZIP / ENUMERATE
//-----------------------------------------------------------------------------

func TestZip(t *testing.T) {
	var keys []int
	var vals []string
	for k, v := range Zip([]int{1, 2, 3}, []string{"a", "b", "c"}) {
		keys = append(keys, k)
		vals = append(vals, v)
	}
	testx.True(t, slices.Equal(keys, []int{1, 2, 3}))
	testx.True(t, slices.Equal(vals, []string{"a", "b", "c"}))
}

func TestZipUnevenLength(t *testing.T) {
	count := 0
	for range Zip([]int{1, 2, 3}, []string{"a", "b"}) {
		count++
	}
	testx.Equal(t, 2, count)
}

func TestEnumerate(t *testing.T) {
	var idxs []int
	var vals []string
	for i, v := range Enumerate([]string{"a", "b", "c"}) {
		idxs = append(idxs, i)
		vals = append(vals, v)
	}
	testx.True(t, slices.Equal(idxs, []int{0, 1, 2}))
	testx.True(t, slices.Equal(vals, []string{"a", "b", "c"}))
}

func TestEnumerateBreak(t *testing.T) {
	count := 0
	for range Enumerate([]string{"a", "b", "c"}) {
		count++
		break
	}
	testx.Equal(t, 1, count)
}
