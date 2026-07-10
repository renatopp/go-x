package dsx

import "math/rand/v2"

// Histogram represents a histogram with a specified number of bins and a range
// defined by min and max values. It provides methods to add values, compute
// the probability density function (PDF), cumulative distribution function
// (CDF), and generate random samples based on the histogram data.
type Histogram struct {
	bins []float64
	min  float64
	max  float64
}

// NewHistogram creates a new Histogram with the specified number of bins and
// range defined by min and max.
func NewHistogram(bins int, min, max float64) *Histogram {
	return &Histogram{
		bins: make([]float64, bins),
		min:  min,
		max:  max,
	}
}

// Add adds a value to the histogram, updating the appropriate bin count and
// adjusting the min and max values if necessary.
func (h *Histogram) Add(value ...float64) {
	for _, v := range value {
		if v < h.min {
			h.min = v
		}
		if v > h.max {
			h.max = v
		}
		binIndex := int((v - h.min) / (h.max - h.min) * float64(len(h.bins)))
		if binIndex < 0 {
			binIndex = 0
		} else if binIndex >= len(h.bins) {
			binIndex = len(h.bins) - 1
		}
		h.bins[binIndex]++
	}
}

// Min returns the minimum value of the histogram's range.
func (h *Histogram) Min() float64 {
	return h.min
}

// Max returns the maximum value of the histogram's range.
func (h *Histogram) Max() float64 {
	return h.max
}

// BinWidth returns the width of each bin in the histogram.
func (h *Histogram) BinWidth() float64 {
	return (h.max - h.min) / float64(len(h.bins))
}

// BinCenter returns the center value of the specified bin.
func (h *Histogram) BinCenter(index int) float64 {
	return h.min + (float64(index)+0.5)*h.BinWidth()
}

// BinCount returns the count of values in the specified bin.
func (h *Histogram) BinCount(index int) float64 {
	if index < 0 || index >= len(h.bins) {
		return 0
	}
	return h.bins[index]
}

// BinRange returns the range of values that fall into the specified bin.
func (h *Histogram) BinRange(index int) (float64, float64) {
	if index < 0 || index >= len(h.bins) {
		return 0, 0
	}
	binWidth := h.BinWidth()
	return h.min + float64(index)*binWidth, h.min + float64(index+1)*binWidth
}

// PDF computes the probability density function of the histogram at a given value.
func (h *Histogram) PDF(value float64) float64 {
	if value < h.min || value > h.max {
		return 0
	}
	for i, count := range h.bins {
		binRangeStart, binRangeEnd := h.BinRange(i)
		if value >= binRangeStart && value < binRangeEnd {
			return count / (h.TotalCount() * h.BinWidth())
		}
	}
	return 0
}

// CDF computes the cumulative distribution function of the histogram at a given value.
func (h *Histogram) CDF(value float64) float64 {
	if value < h.min {
		return 0
	}
	if value > h.max {
		return 1
	}
	cdf := 0.0
	for i, count := range h.bins {
		binRangeStart, binRangeEnd := h.BinRange(i)
		if value >= binRangeEnd {
			cdf += count
		} else if value >= binRangeStart {
			cdf += count * (value - binRangeStart) / (binRangeEnd - binRangeStart)
			break
		}
	}
	return cdf / h.TotalCount()
}

// Sample generates a random sample from the histogram based on the distribution of values in the bins.
func (h *Histogram) Sample() float64 {
	total := h.TotalCount()
	if total == 0 {
		return 0
	}
	r := rand.Float64() * total
	cumulative := 0.0
	for i, count := range h.bins {
		cumulative += count
		if r < cumulative {
			binRangeStart, binRangeEnd := h.BinRange(i)
			return binRangeStart + rand.Float64()*(binRangeEnd-binRangeStart)
		}
	}
	return 0
}

// SampleN generates n random samples from the histogram based on the distribution of values in the bins.
func (h *Histogram) SampleN(n int) []float64 {
	samples := make([]float64, n)
	for i := range n {
		samples[i] = h.Sample()
	}
	return samples
}

// TotalCount returns the total count of values in the histogram by summing the counts of all bins.
func (h *Histogram) TotalCount() float64 {
	total := 0.0
	for _, count := range h.bins {
		total += count
	}
	return total
}

// Normalized returns a new Histogram where the bin counts are normalized to represent probabilities, summing to 1.
func (h *Histogram) Normalized() *Histogram {
	total := h.TotalCount()
	if total == 0 {
		return h
	}
	normalizedBins := make([]float64, len(h.bins))
	for i, count := range h.bins {
		normalizedBins[i] = count / total
	}
	return &Histogram{
		bins: normalizedBins,
		min:  h.min,
		max:  h.max,
	}
}
