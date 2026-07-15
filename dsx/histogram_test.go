package dsx_test

import (
	"math"
	"testing"

	"github.com/renatopp/go-x/dsx"
	"github.com/renatopp/go-x/testx"
)

func newTestHistogram() *dsx.Histogram {
	h := dsx.NewHistogram(4, 0, 10)
	h.Add(1, 3, 6, 9)
	return h
}

func TestHistogramAdd(t *testing.T) {
	h := newTestHistogram()
	testx.AlmostEqual(t, 4, h.TotalCount(), 0.0001)
}

func TestHistogramAddExpandsRange(t *testing.T) {
	h := dsx.NewHistogram(4, 0, 10)
	h.Add(-5, 15)
	testx.AlmostEqual(t, -5, h.Min(), 0.0001)
	testx.AlmostEqual(t, 15, h.Max(), 0.0001)
}

func TestHistogramMinMax(t *testing.T) {
	h := newTestHistogram()
	testx.AlmostEqual(t, 0, h.Min(), 0.0001)
	testx.AlmostEqual(t, 10, h.Max(), 0.0001)
}

func TestHistogramBinWidth(t *testing.T) {
	h := newTestHistogram()
	testx.AlmostEqual(t, 2.5, h.BinWidth(), 0.0001)
}

func TestHistogramBinCenter(t *testing.T) {
	h := newTestHistogram()
	testx.AlmostEqual(t, 1.25, h.BinCenter(0), 0.0001)
	testx.AlmostEqual(t, 8.75, h.BinCenter(3), 0.0001)
	testx.AlmostEqual(t, 8.75, h.BinCenter(-1), 0.0001)
	testx.AlmostEqual(t, 0, h.BinCenter(10), 0.0001)
}

func TestHistogramBinCount(t *testing.T) {
	h := newTestHistogram()
	testx.AlmostEqual(t, 1, h.BinCount(0), 0.0001)
	testx.AlmostEqual(t, 1, h.BinCount(-1), 0.0001)
	testx.AlmostEqual(t, 0, h.BinCount(10), 0.0001)
	testx.AlmostEqual(t, 0, h.BinCount(-10), 0.0001)
}

func TestHistogramBinRange(t *testing.T) {
	h := newTestHistogram()
	start, end := h.BinRange(0)
	testx.AlmostEqual(t, 0, start, 0.0001)
	testx.AlmostEqual(t, 2.5, end, 0.0001)

	start, end = h.BinRange(-1)
	testx.AlmostEqual(t, 7.5, start, 0.0001)
	testx.AlmostEqual(t, 10, end, 0.0001)

	start, end = h.BinRange(10)
	testx.AlmostEqual(t, 0, start, 0.0001)
	testx.AlmostEqual(t, 0, end, 0.0001)
}

func TestHistogramPDF(t *testing.T) {
	h := newTestHistogram()
	testx.AlmostEqual(t, 0.1, h.PDF(1), 0.0001)
	testx.AlmostEqual(t, 0, h.PDF(-5), 0.0001)
	testx.AlmostEqual(t, 0, h.PDF(50), 0.0001)
}

func TestHistogramCDF(t *testing.T) {
	h := newTestHistogram()
	testx.AlmostEqual(t, 0, h.CDF(0), 0.0001)
	testx.AlmostEqual(t, 0.25, h.CDF(2.5), 0.0001)
	testx.AlmostEqual(t, 1, h.CDF(10), 0.0001)
	testx.AlmostEqual(t, 0, h.CDF(-5), 0.0001)
	testx.AlmostEqual(t, 1, h.CDF(50), 0.0001)
}

func TestHistogramSample(t *testing.T) {
	h := newTestHistogram()
	for range 20 {
		v := h.Sample()
		testx.True(t, v >= h.Min() && v <= h.Max())
	}
}

func TestHistogramSampleEmpty(t *testing.T) {
	h := dsx.NewHistogram(4, 0, 10)
	testx.AlmostEqual(t, 0, h.Sample(), 0.0001)
}

func TestHistogramSampleN(t *testing.T) {
	h := newTestHistogram()
	samples := h.SampleN(20)
	testx.Equal(t, 20, len(samples))
	for _, v := range samples {
		testx.False(t, math.IsNaN(v))
	}
}

func TestHistogramTotalCount(t *testing.T) {
	testx.AlmostEqual(t, 0, dsx.NewHistogram(4, 0, 10).TotalCount(), 0.0001)
	testx.AlmostEqual(t, 4, newTestHistogram().TotalCount(), 0.0001)
}

func TestHistogramNormalized(t *testing.T) {
	h := newTestHistogram()
	n := h.Normalized()
	testx.AlmostEqual(t, 0.25, n.BinCount(0), 0.0001)
	testx.AlmostEqual(t, 1, n.TotalCount(), 0.0001)
	// original histogram is not mutated
	testx.AlmostEqual(t, 1, h.BinCount(0), 0.0001)
}

func TestHistogramNormalizedEmpty(t *testing.T) {
	h := dsx.NewHistogram(4, 0, 10)
	n := h.Normalized()
	testx.AlmostEqual(t, 0, n.TotalCount(), 0.0001)
}
