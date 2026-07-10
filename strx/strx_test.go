package strx_test

import (
	"testing"

	"github.com/renatopp/go-x/strx"
	"github.com/renatopp/go-x/testx"
)

func TestIndent(t *testing.T) {
	testx.Equal(t, "  Hello\n  World", strx.Indent("Hello\nWorld", 2))
}

func TestIndentWith(t *testing.T) {
	testx.Equal(t, "--Hello\n--World", strx.IndentWith("Hello\nWorld", 2, "-"))
}

func TestEscape(t *testing.T) {
	testx.Equal(t, "Hello\\r\\nWorld\\t!", strx.Escape("Hello\r\nWorld\t!"))
}

func TestJoinFunc(t *testing.T) {
	items := []int{1, 2, 3}
	testx.Equal(t, "1-2-3", strx.JoinFunc(items, func(i int) string {
		return string(rune('0' + i))
	}, "-"))
}

func TestHumanList(t *testing.T) {
	testx.Equal(t, "", strx.HumanList([]string{}, "and"))
	testx.Equal(t, "Alice", strx.HumanList([]string{"Alice"}, "and"))
	testx.Equal(t, "Alice and Bob", strx.HumanList([]string{"Alice", "Bob"}, "and"))
	testx.Equal(t, "Alice, Bob and Charlie", strx.HumanList([]string{"Alice", "Bob", "Charlie"}, "and"))
	testx.Equal(t, "Alice, Bob or Charlie", strx.HumanList([]string{"Alice", "Bob", "Charlie"}, "or"))
}

func TestHumanListFunc(t *testing.T) {
	testx.Equal(t, "", strx.HumanListFunc([]int{}, func(i int) string {
		return string(rune('0' + i))
	}, "and"))
	testx.Equal(t, "1", strx.HumanListFunc([]int{1}, func(i int) string {
		return string(rune('0' + i))
	}, "and"))
	testx.Equal(t, "1 and 2", strx.HumanListFunc([]int{1, 2}, func(i int) string {
		return string(rune('0' + i))
	}, "and"))
	testx.Equal(t, "1, 2 and 3", strx.HumanListFunc([]int{1, 2, 3}, func(i int) string {
		return string(rune('0' + i))
	}, "and"))
	testx.Equal(t, "1, 2 or 3", strx.HumanListFunc([]int{1, 2, 3}, func(i int) string {
		return string(rune('0' + i))
	}, "or"))
}

func TestPadLeft(t *testing.T) {
	testx.Equal(t, "renato", strx.PadLeft("renato", -1))
	testx.Equal(t, "renato", strx.PadLeft("renato", 0))
	testx.Equal(t, "renato", strx.PadLeft("renato", 4))
	testx.Equal(t, "   renato", strx.PadLeft("renato", 9))
}

func TestPadLeftWith(t *testing.T) {
	testx.Equal(t, "renato", strx.PadLeftWith("renato", -1, "0"))
	testx.Equal(t, "renato", strx.PadLeftWith("renato", 0, "0"))
	testx.Equal(t, "renato", strx.PadLeftWith("renato", 4, "0"))
	testx.Equal(t, "000renato", strx.PadLeftWith("renato", 9, "0"))
}

func TestPadRight(t *testing.T) {
	testx.Equal(t, "renato", strx.PadRight("renato", -1))
	testx.Equal(t, "renato", strx.PadRight("renato", 0))
	testx.Equal(t, "renato", strx.PadRight("renato", 4))
	testx.Equal(t, "renato   ", strx.PadRight("renato", 9))
}

func TestPadRightWith(t *testing.T) {
	testx.Equal(t, "renato", strx.PadRightWith("renato", -1, "0"))
	testx.Equal(t, "renato", strx.PadRightWith("renato", 0, "0"))
	testx.Equal(t, "renato", strx.PadRightWith("renato", 4, "0"))
	testx.Equal(t, "renato000", strx.PadRightWith("renato", 9, "0"))
}

func TestPadCenter(t *testing.T) {
	testx.Equal(t, "renato", strx.PadCenter("renato", -1))
	testx.Equal(t, "renato", strx.PadCenter("renato", 0))
	testx.Equal(t, "renato", strx.PadCenter("renato", 4))
	testx.Equal(t, "   renato  ", strx.PadCenter("renato", 11))
	testx.Equal(t, "   renato   ", strx.PadCenter("renato", 12))
}

func TestPadCenterWith(t *testing.T) {
	testx.Equal(t, "renato", strx.PadCenterWith("renato", -1, "0"))
	testx.Equal(t, "renato", strx.PadCenterWith("renato", 0, "0"))
	testx.Equal(t, "renato", strx.PadCenterWith("renato", 4, "0"))
	testx.Equal(t, "000renato00", strx.PadCenterWith("renato", 11, "0"))
	testx.Equal(t, "000renato000", strx.PadCenterWith("renato", 12, "0"))
}

func TestFirstUp(t *testing.T) {
	testx.Equal(t, "Renato", strx.FirstUp("renato"))
	testx.Equal(t, "Renato", strx.FirstUp("Renato"))
	testx.Equal(t, "🌱renato", strx.FirstUp("🌱renato"))
}

func TestFirstLow(t *testing.T) {
	testx.Equal(t, "renato", strx.FirstLow("Renato"))
	testx.Equal(t, "renato", strx.FirstLow("renato"))
	testx.Equal(t, "🌱renato", strx.FirstLow("🌱renato"))
}

func TestTrimSpaces(t *testing.T) {
	testx.Equal(t, "Hello   World", strx.TrimSpaces("  Hello   World  "))
	testx.Equal(t, "Hello World", strx.TrimSpaces("Hello World"))
	testx.Equal(t, "", strx.TrimSpaces("    "))
}

func TestIsBlank(t *testing.T) {
	testx.True(t, strx.IsBlank(""))
	testx.True(t, strx.IsBlank("  \n\r\t  "))
	testx.False(t, strx.IsBlank("Hello"))
}

func TestTruncate(t *testing.T) {
	testx.Equal(t, "Hello", strx.Truncate("Hello", 10))
	testx.Equal(t, "Hello", strx.Truncate("Hello", 5))
	testx.Equal(t, "Hell", strx.Truncate("Hello", 4))
	testx.Equal(t, "Hel", strx.Truncate("Hello", 3))
}

func TestEllipsis(t *testing.T) {
	testx.Equal(t, "🌱⚡🧹🥸👍", strx.Ellipsis("🌱⚡🧹🥸👍", 10))
	testx.Equal(t, "🌱⚡🧹🥸👍", strx.Ellipsis("🌱⚡🧹🥸👍", 5))
	testx.Equal(t, "🌱...", strx.Ellipsis("🌱⚡🧹🥸👍", 4))
	testx.Equal(t, "...", strx.Ellipsis("🌱⚡🧹🥸👍", 3))
	testx.Equal(t, "...", strx.Ellipsis("🌱⚡🧹🥸👍", 2))
	testx.Equal(t, "...", strx.Ellipsis("🌱⚡🧹🥸👍", 1))
	testx.Equal(t, "...", strx.Ellipsis("🌱⚡🧹🥸👍", 0))
}

func TestIterString(t *testing.T) {
	values := []string{}
	for i, s := range strx.IterString("Hello") {
		values = append(values, s)
		testx.Equal(t, i, i)
	}
	testx.Equal(t, "H", values[0])
	testx.Equal(t, "e", values[1])
	testx.Equal(t, "l", values[2])
	testx.Equal(t, "l", values[3])
	testx.Equal(t, "o", values[4])
}

func TestIterRunes(t *testing.T) {
	values := []rune{}
	for i, r := range strx.IterRunes("Hello") {
		values = append(values, r)
		testx.Equal(t, i, i)
	}
	testx.Equal(t, 'H', values[0])
	testx.Equal(t, 'e', values[1])
	testx.Equal(t, 'l', values[2])
	testx.Equal(t, 'l', values[3])
	testx.Equal(t, 'o', values[4])
}

func TestFormat(t *testing.T) {
	testx.Equal(t, "formated 010", strx.Format("formated %03d", 10))
}

func TestReverse(t *testing.T) {
	testx.Equal(t, "olleH", strx.Reverse("Hello"))
	testx.Equal(t, "🌱⚡🧹🥸👍", strx.Reverse("👍🥸🧹⚡🌱"))
}

func TestToPrintableAscii(t *testing.T) {
	testx.Equal(t, "HelloWorld", strx.ToPrintableAscii("Hello\nWorld"))
	testx.Equal(t, "", strx.ToPrintableAscii("你好"))
}
