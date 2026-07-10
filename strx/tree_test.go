package strx_test

import (
	"testing"

	"github.com/renatopp/go-x/strx"
	"github.com/renatopp/go-x/testx"
)

func makeTree(lines ...string) string {
	result := ""
	for i, l := range lines {
		if i > 0 {
			result += "\n"
		}
		result += l
	}
	return result
}

func TestTreeEmpty(t *testing.T) {
	tree := strx.NewTree("root")
	testx.Equal(t, "root", tree.Render())
}

func TestTreeSingleChild(t *testing.T) {
	expected := makeTree(
		"root",
		"└─ child",
	)
	tree := strx.NewTree("root")
	tree.Add("child")
	testx.Equal(t, expected, tree.Render())
}

func TestTreeMultipleChildren(t *testing.T) {
	expected := makeTree(
		"root",
		"├─ a",
		"├─ b",
		"└─ c",
	)
	tree := strx.NewTree("root")
	tree.Add("a")
	tree.Add("b")
	tree.Add("c")
	testx.Equal(t, expected, tree.Render())
}

func TestTreeNested(t *testing.T) {
	expected := makeTree(
		"root",
		"├─ fruits",
		"│  ├─ apple",
		"│  └─ banana",
		"└─ vegetables",
		"   └─ carrot",
	)
	tree := strx.NewTree("root")
	fruits := tree.Add("fruits")
	fruits.Add("apple")
	fruits.Add("banana")
	tree.Add("vegetables").Add("carrot")
	testx.Equal(t, expected, tree.Render())
}

func TestTreeDeepNesting(t *testing.T) {
	expected := makeTree(
		"root",
		"└─ a",
		"   └─ b",
		"      └─ c",
		"         └─ d",
	)
	tree := strx.NewTree("root")
	tree.Add("a").Add("b").Add("c").Add("d")
	testx.Equal(t, expected, tree.Render())
}

func TestTreeAsciiStyle(t *testing.T) {
	expected := makeTree(
		"root",
		"+- a",
		"|  +- a1",
		"|  +- a2",
		"+- b",
	)
	tree := strx.NewTree("root")
	tree.WithStyle(strx.TreeStyleAscii)
	a := tree.Add("a")
	a.Add("a1")
	a.Add("a2")
	tree.Add("b")
	testx.Equal(t, expected, tree.Render())
}

func TestTreeUnicodeRoundStyle(t *testing.T) {
	expected := makeTree(
		"root",
		"├─ a",
		"╰─ b",
	)
	tree := strx.NewTree("root")
	tree.WithStyle(strx.TreeStyleUnicodeRound)
	tree.Add("a")
	tree.Add("b")
	testx.Equal(t, expected, tree.Render())
}

func TestTreeMultiLineLabel(t *testing.T) {
	expected := makeTree(
		"root",
		"├─ line1",
		"│  line2",
		"│  └─ child",
		"└─ single",
	)
	tree := strx.NewTree("root")
	n := tree.Add("line1\nline2")
	n.Add("child")
	tree.Add("single")
	testx.Equal(t, expected, tree.Render())
}

func TestTreeMultiLineRoot(t *testing.T) {
	expected := makeTree(
		"first",
		"second",
		"└─ child",
	)
	tree := strx.NewTree("first\nsecond")
	tree.Add("child")
	testx.Equal(t, expected, tree.Render())
}

func TestTreeAddNode(t *testing.T) {
	expected := makeTree(
		"root",
		"├─ sub",
		"│  ├─ x",
		"│  └─ y",
		"└─ leaf",
	)
	sub := strx.NewTreeNode("sub")
	sub.Add("x")
	sub.Add("y")

	tree := strx.NewTree("root")
	tree.AddNode(sub)
	tree.Add("leaf")
	testx.Equal(t, expected, tree.Render())
}

func TestTreeNewTreeFrom(t *testing.T) {
	expected := makeTree(
		"root",
		"└─ child",
	)
	root := strx.NewTreeNode("root")
	root.Add("child")
	tree := strx.NewTreeFrom(root)
	testx.Equal(t, expected, tree.Render())
}
