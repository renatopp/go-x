package strx

import "strings"

// STYLES ---------------------------------------------------------------------

// TreeStyle defines the characters used to render the tree branches.
//
// Example with TreeStyleUnicode:
//
//	root
//	├─ child            <- Child connector
//	│  └─ grandchild  <- Pipe prefix + LastChild connector
//	└─ last child       <- LastChild connector
//	    └─ grandchild   <- Indent prefix + LastChild connector
type TreeStyle struct {
	Child     string // connector for non-last children  (e.g. "├─ ")
	LastChild string // connector for the last child     (e.g. "└─ ")
	Pipe      string // prefix under a non-last child    (e.g. "│  ")
	Indent    string // prefix under the last child      (e.g. "   ")
}

var TreeStyleAscii = TreeStyle{
	Child:     "+- ",
	LastChild: "+- ",
	Pipe:      "|  ",
	Indent:    "   ",
}

var TreeStyleUnicode = TreeStyle{
	Child:     "├─ ",
	LastChild: "└─ ",
	Pipe:      "│  ",
	Indent:    "   ",
}

var TreeStyleUnicodeRound = TreeStyle{
	Child:     "├─ ",
	LastChild: "╰─ ",
	Pipe:      "│  ",
	Indent:    "   ",
}

var TreeStyleUnicodeDouble = TreeStyle{
	Child:     "╠═ ",
	LastChild: "╚═ ",
	Pipe:      "║  ",
	Indent:    "   ",
}

var TreeStyleUnicodeCompact = TreeStyle{
	Child:     "├ ",
	LastChild: "└ ",
	Pipe:      "│ ",
	Indent:    "  ",
}
var TreeStyleUnicodeDoubleCompact = TreeStyle{
	Child:     "╠ ",
	LastChild: "╚ ",
	Pipe:      "║ ",
	Indent:    "  ",
}

// TREE NODE ------------------------------------------------------------------

// TreeNode represents a single node in a tree, with a label and zero or more
// children.
type TreeNode struct {
	Label    string
	Children []*TreeNode
}

// NewTreeNode creates a new standalone node with the given label.
func NewTreeNode(label string) *TreeNode {
	return &TreeNode{Label: label}
}

// Add creates a new child node with the given label, appends it to the node's
// children, and returns the new child. This allows building a tree top-down:
//
//	child := parent.Add("child")
//	child.Add("grandchild")
func (n *TreeNode) Add(label string) *TreeNode {
	child := &TreeNode{Label: label}
	n.Children = append(n.Children, child)
	return child
}

// AddNode appends an existing node as a child and returns the parent, allowing
// horizontal chaining when attaching pre-built subtrees:
//
//	subtree := strx.NewTreeNode("sub")
//	subtree.Add("a")
//	root.AddNode(subtree).AddNode(otherSubtree)
func (n *TreeNode) AddNode(child *TreeNode) *TreeNode {
	n.Children = append(n.Children, child)
	return n
}

// TREE -----------------------------------------------------------------------

// Tree renders a hierarchical tree structure as a string. Nodes are built via
// the root TreeNode; the Tree itself only controls rendering style.
//
// Usage:
//
//	t := strx.NewTree("root")
//	a := t.Add("fruits")
//	a.Add("apple")
//	a.Add("banana")
//	t.Add("vegetables").Add("carrot")
//	t.WithStyle(strx.TreeStyleUnicode)
//	fmt.Println(t.Render())
//
// Output:
//
//	root
//	├── fruits
//	│   ├── apple
//	│   └── banana
//	└── vegetables
//	    └── carrot
type Tree struct {
	root  *TreeNode
	style TreeStyle
}

// NewTree creates a new Tree whose root node has the given label. The default
// style is TreeStyleUnicode.
func NewTree(label string) *Tree {
	return &Tree{
		root:  &TreeNode{Label: label},
		style: TreeStyleUnicode,
	}
}

// NewTreeFrom creates a new Tree using an existing node as the root.
func NewTreeFrom(root *TreeNode) *Tree {
	return &Tree{
		root:  root,
		style: TreeStyleUnicode,
	}
}

// Root returns the root TreeNode, allowing direct manipulation.
func (t *Tree) Root() *TreeNode {
	return t.root
}

// Add creates a new child of the root with the given label and returns it.
// This is a convenience shorthand for t.Root().Add(label).
func (t *Tree) Add(label string) *TreeNode {
	return t.root.Add(label)
}

// AddNode appends an existing node as a direct child of the root and returns
// the Tree for chaining.
func (t *Tree) AddNode(child *TreeNode) *Tree {
	t.root.AddNode(child)
	return t
}

// WithStyle sets the branch style used when rendering.
func (t *Tree) WithStyle(style TreeStyle) *Tree {
	t.style = style
	return t
}

// Render generates and returns the tree as a string. The trailing newline is
// stripped so the result composes cleanly with fmt.Println.
func (t *Tree) Render() string {
	b := &strings.Builder{}
	t.renderNode(b, t.root, "", true)
	result := b.String()
	if len(result) > 0 && result[len(result)-1] == '\n' {
		result = result[:len(result)-1]
	}
	return result
}

// renderNode writes one node and then recurses into its children.
// prefix is the string prepended to every line of the current node
// (does not include the connector itself — the caller writes that).
// isRoot distinguishes the root node which has no connector.
func (t *Tree) renderNode(b *strings.Builder, node *TreeNode, prefix string, isRoot bool) {
	labelLines := Lines(node.Label)

	if isRoot {
		// Root label: first line has no connector, continuation lines are
		// flush-left (no extra prefix because the root has no parent pipe).
		b.WriteString(labelLines[0])
		b.WriteString("\n")
		for _, line := range labelLines[1:] {
			b.WriteString(line)
			b.WriteString("\n")
		}
	}
	// Non-root nodes: the connector + first label line are written by the
	// parent's loop below; this function only handles children.

	for i, child := range node.Children {
		isLast := i == len(node.Children)-1

		var connector, childPrefix string
		if isLast {
			connector = t.style.LastChild
			childPrefix = prefix + t.style.Indent
		} else {
			connector = t.style.Child
			childPrefix = prefix + t.style.Pipe
		}

		childLines := Lines(child.Label)

		// First line of the child label, preceded by the branch connector.
		b.WriteString(prefix)
		b.WriteString(connector)
		b.WriteString(childLines[0])
		b.WriteString("\n")

		// Continuation lines of a multi-line label use childPrefix so they
		// visually align under the label text, not under the connector.
		for _, line := range childLines[1:] {
			b.WriteString(childPrefix)
			b.WriteString(line)
			b.WriteString("\n")
		}

		// Recurse — children of this child use childPrefix as their base.
		t.renderNode(b, child, childPrefix, false)
	}
}
