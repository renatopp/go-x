package strx_test

import (
	"strings"
	"testing"

	"github.com/renatopp/go-x/strx"
	"github.com/renatopp/go-x/testx"
)

func makeTable(s ...string) string {
	return strings.Join(s, "\n")
}

func TestTableEmpty(t *testing.T) {
	b := strx.NewTable()
	testx.Equal(t, "", b.Render())
}

func TestTableDataOnly(t *testing.T) {
	expected := makeTable(
		"+---+---+---+",
		"| a | b | c |",
		"| d | e | f |",
		"+---+---+---+",
	)
	b := strx.NewTable()
	b.Data("a", "b", "c")
	b.Data("d", "e", "f")
	testx.Equal(t, expected, b.Render())
}

func TestTableMetaOnly(t *testing.T) {
	expected := makeTable(
		"+===+===+===+",
		"| a | b | c |",
		"+===+===+===+",
	)
	b := strx.NewTable()
	b.Meta("a", "b", "c")
	b.WithStyle(strx.TableStyleAsciiSeparated)
	testx.Equal(t, expected, b.Render())
}

func TestSectionOnly(t *testing.T) {
	expected := makeTable(
		"+=========+",
		"| SECTION |",
		"+=========+",
	)
	b := strx.NewTable()
	b.MetaSection("SECTION")
	b.WithStyle(strx.TableStyleAsciiSeparated)
	testx.Equal(t, expected, b.Render())
}

func TestMixed(t *testing.T) {
	expected := makeTable(
		"╔═══════════════════════════════════════════╗",
		"║ THIS IS A VERY LARGE TITLE THAT SHOULD BE ║",
		"║         BROKEN INTO MULTIPLE LINES        ║",
		"╠════════════╦══════════════════════╦═══════╣",
		"║     ID     ║         NAME         ║  AGE  ║",
		"╠════════════╬══════════════════════╬═══════╣",
		"║      0     ║        renato        ║   39  ║",
		"║            ║        r2p.dev       ║       ║",
		"╠════════════╬══════════════════════╬═══════╣",
		"║          1 ║                maria ║    28 ║",
		"╠════════════╩══════════════════════╩═══════╣",
		"║  This is another very large section that  ║",
		"║    should be broken into multiple lines   ║",
		"╠════════════╦══════════════════════╦═══════╣",
		"║ 2          ║                 joão ║   3   ║",
		"╠════════════╩══════════════════════╩═══════╣",
		"║                   FOOTER                  ║",
		"╚═══════════════════════════════════════════╝",
	)

	b := strx.NewTable()
	b.MetaSection("THIS IS A VERY LARGE TITLE THAT SHOULD BE BROKEN INTO MULTIPLE LINES")
	b.Meta("ID", "NAME", "AGE")
	b.Data("0", "renato\nr2p.dev", 39).ToCenter()
	b.Data("1", "maria", 28).ToRight()
	b.DataSection("This is another very large section that should be broken into multiple lines")
	b.Data("2", "joão", 3).WithAlignments(strx.TableLeft, strx.TableRight, strx.TableCenter)
	b.MetaSection("FOOTER")
	b.WithLength(10, 20, 5)
	b.WithStyle(strx.TableStyleUnicodeDoubleGrid)

	testx.Equal(t, expected, b.Render())
}

func TestNested(t *testing.T) {
	expected := makeTable(
		"+----------------+",
		"|      OUTER     |",
		"+----------------+",
		"| +------------+ |",
		"| |    INNER   | |",
		"| +------------+ |",
		"+----------------+",
	)

	inner := strx.NewTable()
	inner.MetaSection("INNER")
	inner.WithLength(10)

	outer := strx.NewTable()
	outer.MetaSection("OUTER")
	outer.Data(inner.Render())
	testx.Equal(t, expected, outer.Render())
}
