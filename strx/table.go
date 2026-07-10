package strx

import (
	"strings"
)

type TableAlignment int8

// CONSTANTS AND VARIABLES ----------------------------------------------------

const (
	TableAuto TableAlignment = iota
	TableLeft
	TableCenter
	TableRight
)

var TableStyleAscii = TableStyle{
	"+", "+", "+", "+", "+", "+", "+", "+", "+", "-", "-", "|", "|",
	"+", "+", "+", "", "", "", "+", "+", "+", "-", "", "|", "|",
}
var TableStyleAsciiSeparated = TableStyle{
	"+", "+", "+", "+", "+", "+", "+", "+", "+", "=", "=", "|", "|",
	"+", "+", "+", "+", "+", "+", "+", "+", "+", "-", "-", "|", "|",
}
var TableStyleAsciiCompact = TableStyle{
	"", "", "", "", " ", "", "", "", "", "", "-", "", " ",
	"", "", "", "", "", "", "", "", "", "", "", "", " ",
}
var TableStyleAsciiDots = TableStyle{
	".", ".", ".", ":", ":", ":", ":", ":", ":", ".", ".", ":", ":",
	".", ".", ".", "", "", "", ":", ":", ":", ".", "", ":", ":",
}
var TableStyleGithub = TableStyle{
	"", "", "", "|", "|", "|", "", "", "", "", "-", "|", "|",
	"", "", "", "", "", "", "", "", "", "", "", "|", "|",
}
var TableStyleReddit = TableStyle{
	"", "", "", "", "|", "", "", "", "", "", "-", "", "|",
	"", "", "", "", "", "", "", "", "", "", "", "", "|",
}
var TableStyleRestructuredGrid = TableStyle{
	"+", "+", "+", "+", "+", "+", "+", "+", "+", "-", "=", "|", "|",
	"+", "+", "+", "", "", "", "+", "+", "+", "-", "", "|", "|",
}
var TableStyleRestructuredSimple = TableStyle{
	"", " ", "", "", " ", "", "", " ", "", "=", "=", "", " ",
	"", " ", "", "", "", "", "", " ", "", "=", "", "", " ",
}
var TableStyleUnicode = TableStyle{
	"┌", "┬", "┐", "├", "┼", "┤", "└", "┴", "┘", "─", "─", "│", "│",
	"┌", "┬", "┐", "", "", "", "└", "┴", "┘", "─", "", "│", "│",
}
var TableStyleUnicodeGrid = TableStyle{
	"┌", "┬", "┐", "├", "┼", "┤", "└", "┴", "┘", "─", "─", "│", "│",
	"┌", "┬", "┐", "├", "┼", "┤", "└", "┴", "┘", "─", "─", "│", "│",
}
var TableStyleUnicodeDouble = TableStyle{
	"╔", "╦", "╗", "╠", "╬", "╣", "╚", "╩", "╝", "═", "═", "║", "║",
	"╔", "╦", "╗", "", "", "", "╚", "╩", "╝", "═", "", "║", "║",
}
var TableStyleUnicodeDoubleGrid = TableStyle{
	"╔", "╦", "╗", "╠", "╬", "╣", "╚", "╩", "╝", "═", "═", "║", "║",
	"╔", "╦", "╗", "╠", "╬", "╣", "╚", "╩", "╝", "═", "═", "║", "║",
}

// TableStyle defines the characters used to render the table. It has two sets
// of characters: one for meta rows and another for data rows.
//
// a===b===c
// $   |   $
// d---e---f
// $   |   $
// g===h===i
//
// a: TL, b: TM, c: TR (top left, middle and right)
// d: ML, e: MM, f: MR (middle left, middle and right)
// g: BL, h: BM, i: BR (bottom left, middle and right)
// $: VE, |: VI (vertical external and internal)
// =: HE, -: HI (horizontal external and internal)
type TableStyle struct {
	MTL, MTM, MTR, MML, MMM, MMR, MBL, MBM, MBR, MHE, MHI, MVE, MVI string
	DTL, DTM, DTR, DML, DMM, DMR, DBL, DBM, DBR, DHE, DHI, DVE, DVI string
}

// TABLE ROW ------------------------------------------------------------------

// TableRow represents a single row in the table.
type TableRow struct {
	lines      int
	values     []string
	alignments []TableAlignment
	meta       bool
	section    bool
}

// ToLeft sets the alignment of all cells in the row to left. After using it,
// the row loses the auto-alignment information.
func (t *TableRow) ToLeft() *TableRow {
	for i := range t.values {
		t.alignments[i] = TableLeft
	}
	return t
}

// ToCenter sets the alignment of all cells in the row to center. After using
// it, the row loses the auto-alignment information.
func (t *TableRow) ToCenter() *TableRow {
	for i := range t.values {
		t.alignments[i] = TableCenter
	}
	return t
}

// ToRight sets the alignment of all cells in the row to right. After using it,
// the row loses the auto-alignment information.
func (t *TableRow) ToRight() *TableRow {
	for i := range t.values {
		t.alignments[i] = TableRight
	}
	return t
}

// WithAlignments sets the alignment of each cell in the row according to the
// provided alignments. After using it, the row loses the auto-alignment
// information.
func (t *TableRow) WithAlignments(alignments ...TableAlignment) *TableRow {
	t.alignments = alignments[:len(t.values)]
	return t
}

// TABLE ----------------------------------------------------------------------

// Table is a data structure to render a table in a string format. It supports
// meta rows (such as headers and footers), data rows (for the main content),
// arbitrary number of columns, section rows (a single cell that spans for all
// columns, for both meta and data) and different styles (ASCII, Unicode, etc).
//
// Usage:

// t := strx.NewTable()
// t.MetaSection("This is a title")
// t.Meta("Id", "Name", "Value")
// t.Data("1", "Foo", "10")
// t.Data("2", "Bar", "20")
// t.MetaSection("This is a footer")
// t.WithStyle(strx.TableStyleUnicode)
// println(t.Render())
type Table struct {
	cols    int
	rows    []*TableRow
	lengths []int
	style   TableStyle
	align   TableAlignment
}

// NewTable creates a new table instance.
func NewTable() *Table {
	return &Table{
		cols:    0,
		rows:    make([]*TableRow, 0),
		lengths: make([]int, 0),
		style:   TableStyleAscii,
		align:   TableAuto,
	}
}

// Meta adds a meta row to the table with the provided values. Meta rows are
// typically used for headers and footers. The values are automatically aligned
// to the center.
func (t *Table) Meta(values ...any) *TableRow {
	return t.addRow(true, false, values...)
}

// MetaSection adds a meta section row to the table with the provided value. A
// section row is a single cell that spans for all columns. Meta rows are
// typically used for headers and footers. The value is automatically aligned to
// the center.
func (t *Table) MetaSection(values any) *TableRow {
	return t.addRow(true, true, values)
}

// Data adds a data row to the table with the provided values. Data rows are
// typically used for the main content. The values are automatically aligned to
// the left for strings and to the right for numbers.
func (t *Table) Data(values ...any) *TableRow {
	return t.addRow(false, false, values...)
}

// DataSection adds a data section row to the table with the provided value. A
// section row is a single cell that spans for all columns. Data rows are typically
// used for the main content. The value is automatically aligned to the left for
// strings and to the right for numbers.
func (t *Table) DataSection(values any) *TableRow {
	return t.addRow(false, true, values)
}

// WithStyle sets the style of the table.
func (t *Table) WithStyle(style TableStyle) *Table {
	t.style = style
	return t
}

// WithAlignAuto sets the alignment of the table to auto. In this mode, the
// alignment of each cell is determined by the type of its value or by setting
// the alignment of the row.
func (t *Table) WithAlignAuto() *Table {
	t.align = TableAuto
	return t
}

// WithAlignLeft forces the alignment of all cells in the table to left. For
// both meta and data rows.
func (t *Table) WithAlignLeft() *Table {
	t.align = TableLeft
	return t
}

// WithAlignCenter forces the alignment of all cells in the table to center. For
// both meta and data rows.
func (t *Table) WithAlignCenter() *Table {
	t.align = TableCenter
	return t
}

// WithAlignRight forces the alignment of all cells in the table to right. For
// both meta and data rows.
func (t *Table) WithAlignRight() *Table {
	t.align = TableRight
	return t
}

// WithLength sets the length of each column in the table. If a length is not
// provided for a column, it is automatically determined by the longest cell in
// that column. If you add new data after setting the length, the cell will be
// adjusted to the new length.
func (t *Table) WithLength(length ...int) *Table {
	for i := range length {
		if i < len(t.lengths) {
			t.lengths[i] = length[i]
		} else {
			t.lengths = append(t.lengths, length[i])
		}
	}

	return t
}

// Render generates the table string.
func (t *Table) Render() string {
	if len(t.rows) == 0 {
		return ""
	}

	b := &strings.Builder{}
	t.renderBorder(b, 0)
	for i := range t.rows {
		t.renderRow(b, i)
		t.renderBorder(b, i+1)
	}
	result := b.String()
	return result[:len(result)-1] // remove last newline
}

func (t *Table) addRow(meta bool, section bool, values ...any) *TableRow {
	row := &TableRow{
		values:     make([]string, len(values)),
		alignments: make([]TableAlignment, len(values)),
		meta:       meta,
		section:    section,
	}
	for i := range row.values {
		v, a := t.toString(values[i])
		row.values[i] = v
		row.alignments[i] = a
		row.lines = max(row.lines, Count(v, "\n")+1)
		if meta || section {
			row.alignments[i] = TableCenter
		}
		if !section {
			if i >= len(t.lengths) {
				t.lengths = append(t.lengths, 0)
			}
			for _, line := range IterLines(v) {
				t.lengths[i] = max(t.lengths[i], Length(line))
			}
		}
	}
	t.cols = max(len(values), t.cols)
	t.rows = append(t.rows, row)
	return row
}

func (t *Table) toString(value any) (string, TableAlignment) {
	switch v := value.(type) {
	case string:
		return v, TableLeft
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		return Format("%d", v), TableRight
	case float32, float64:
		return Format("%.2f", v), TableRight
	default:
		return Format("%v", v), TableLeft
	}
}

func (t *Table) renderBorder(b *strings.Builder, i int) {
	left, mid, right, h := t.getBorderParts(i)
	if h == "" {
		return
	}

	rowAbove, rowBelow := t.getBorderRows(i)
	basedRow := rowAbove
	if (rowAbove == nil || rowAbove.section) && rowBelow != nil {
		basedRow = rowBelow
	}

	if basedRow.section {
		length := t.getTotalLength()
		dividers := Length(mid) * max(1, t.cols)
		paddings := 2 * max(1, t.cols)
		b.WriteString(Format("%s%s%s", left, strings.Repeat(h, length+dividers+paddings-1), right))
	} else {
		b.WriteString(left)
		for i, length := range t.lengths {
			b.WriteString(strings.Repeat(h, length+2))
			if i < len(t.lengths)-1 {
				b.WriteString(mid)
			}
		}
		b.WriteString(right)
	}
	b.WriteString("\n")
}

func (t *Table) renderRow(b *strings.Builder, i int) {
	row := t.rows[i]
	ve, vi := t.getVerticalSeparator(i)

	cells := make([][]string, len(row.values))
	lines := row.lines

	if row.section {
		length := t.getTotalLength()
		dividers := Length(vi) * t.cols
		paddings := 2 * t.cols
		size := length + dividers + paddings - 3
		value := WrapWord(row.values[0], size)
		lines = Count(value, "\n") + 1
		align := row.alignments[0]
		if t.align != TableAuto {
			align = t.align
		}
		cells[0] = t.getCellLines(value, align, size, lines)

	} else {
		for c, value := range row.values {
			length := t.lengths[c]
			align := row.alignments[c]
			if t.align != TableAuto {
				align = t.align
			}
			cells[c] = t.getCellLines(value, align, length, row.lines)
		}
	}

	for l := 0; l < lines; l++ {
		b.WriteString(ve)
		for c, cell := range cells {
			b.WriteString(Format(" %s ", cell[l]))
			if c < len(row.values)-1 {
				b.WriteString(vi)
			}
		}
		b.WriteString(ve)
		b.WriteString("\n")
	}
}

func (t *Table) getCellLines(value string, align TableAlignment, length int, lines int) []string {
	result := make([]string, lines)
	for i := range result {
		result[i] = Repeat(" ", length)
	}
	for i, line := range IterLines(value) {
		result[i] = t.alignValue(line, align, length)
	}
	return result
}

func (t *Table) getBorderRows(i int) (above, below *TableRow) {
	var rowAbove, rowBelow *TableRow
	if i > 0 {
		rowAbove = t.rows[i-1]
	}
	if i < len(t.rows) {
		rowBelow = t.rows[i]
	}
	return rowAbove, rowBelow
}

func (t *Table) getBorderParts(i int) (left, mid, right, h string) {
	rowAbove, rowBelow := t.getBorderRows(i)

	isMeta := rowAbove != nil && rowAbove.meta || rowBelow != nil && rowBelow.meta

	switch {
	case rowAbove == nil && isMeta:
		return t.style.MTL, t.style.MTM, t.style.MTR, t.style.MHE
	case rowAbove == nil && !isMeta:
		return t.style.DTL, t.style.DTM, t.style.DTR, t.style.DHE
	case rowBelow == nil && isMeta:
		return t.style.MBL, t.style.MBM, t.style.MBR, t.style.MHE
	case rowBelow == nil && !isMeta:
		return t.style.DBL, t.style.DBM, t.style.DBR, t.style.DHE
	case isMeta:
		switch {
		case rowAbove.section:
			return t.style.MML, t.style.MTM, t.style.MMR, t.style.MHI
		case rowBelow.section:
			return t.style.MML, t.style.MBM, t.style.MMR, t.style.MHI
		default:
			return t.style.MML, t.style.MMM, t.style.MMR, t.style.MHI
		}
	default:
		switch {
		case rowAbove.section:
			return t.style.DML, t.style.DTM, t.style.DMR, t.style.DHI
		case rowBelow.section:
			return t.style.DML, t.style.DBM, t.style.DMR, t.style.DHI
		default:
			return t.style.DML, t.style.DMM, t.style.DMR, t.style.DHI
		}
	}
}

func (t *Table) getVerticalSeparator(i int) (ve, vi string) {
	if t.rows[i].meta {
		return t.style.MVE, t.style.MVI
	} else {
		return t.style.DVE, t.style.DVI
	}
}

func (t *Table) getTotalLength() int {
	// In case of only section rows
	if len(t.lengths) == 0 {
		length := 0
		for _, row := range t.rows {
			for _, value := range row.values {
				for _, line := range IterLines(value) {
					length = max(Length(line), length)
				}
			}
		}
		return length
	}

	total := 0
	for _, length := range t.lengths {
		total += length
	}
	return total
}

func (t *Table) alignValue(v string, align TableAlignment, length int) string {
	b := &strings.Builder{}
	for _, line := range IterLines(v) {
		switch align {
		case TableLeft:
			b.WriteString(PadRight(line, length))
		case TableCenter:
			b.WriteString(PadCenter(line, length))
		case TableRight:
			b.WriteString(PadLeft(line, length))
		default:
			b.WriteString(line)
		}
	}
	return b.String()
}
