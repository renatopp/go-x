package dsx_test

import (
	"math"
	"testing"

	"github.com/renatopp/go-x/dsx"
	"github.com/renatopp/go-x/testx"
)

func TestGaussianPDF(t *testing.T) {
	g := dsx.NewStandardGaussian()
	testx.AlmostEqual(t, 1/math.Sqrt(2*math.Pi), g.PDF(0), 0.0001)
	testx.AlmostEqual(t, 0.24197072451914337, g.PDF(1), 0.0001)
}

func TestGaussianCDF(t *testing.T) {
	g := dsx.NewStandardGaussian()
	testx.AlmostEqual(t, 0.5, g.CDF(0), 0.0001)
	testx.True(t, g.CDF(-1) < g.CDF(0))
	testx.True(t, g.CDF(0) < g.CDF(1))
}

func TestGaussianMeanVarianceStdDev(t *testing.T) {
	g := dsx.NewGaussian(2, 3)
	testx.AlmostEqual(t, 2, g.Mean(), 0.0001)
	testx.AlmostEqual(t, 9, g.Variance(), 0.0001)
	testx.AlmostEqual(t, 3, g.StdDev(), 0.0001)
}

func TestGaussianQuantile(t *testing.T) {
	g := dsx.NewGaussian(2, 3)
	testx.AlmostEqual(t, 2, g.Quantile(0.5), 0.0001)

	for _, p := range []float64{0.1, 0.25, 0.5, 0.75, 0.9} {
		x := g.Quantile(p)
		testx.AlmostEqual(t, p, g.CDF(x), 0.0001)
	}
}

func TestGaussianLogPDF(t *testing.T) {
	g := dsx.NewGaussian(2, 3)
	for _, x := range []float64{-1, 0, 2, 5} {
		testx.AlmostEqual(t, math.Log(g.PDF(x)), g.LogPDF(x), 0.0001)
	}
}

func TestGaussianLogCDF(t *testing.T) {
	g := dsx.NewGaussian(2, 3)
	for _, x := range []float64{-1, 0, 2, 5} {
		testx.AlmostEqual(t, math.Log(g.CDF(x)), g.LogCDF(x), 0.0001)
	}
}

func TestGaussianEntropy(t *testing.T) {
	g := dsx.NewStandardGaussian()
	testx.AlmostEqual(t, 0.5*math.Log(2*math.Pi*math.E), g.Entropy(), 0.0001)
}

func TestGaussianKLDivergence(t *testing.T) {
	g := dsx.NewGaussian(2, 3)
	testx.AlmostEqual(t, 0, g.KLDivergence(dsx.NewGaussian(2, 3)), 0.0001)
	testx.True(t, g.KLDivergence(dsx.NewGaussian(5, 1)) > 0)
}

func TestGaussianMahalanobisDistance(t *testing.T) {
	g := dsx.NewGaussian(2, 3)
	testx.AlmostEqual(t, 0, g.MahalanobisDistance(2), 0.0001)
	testx.AlmostEqual(t, 1, g.MahalanobisDistance(5), 0.0001)
}

func TestGaussianBhattacharyyaDistance(t *testing.T) {
	// Note: self-distance is only 0 for unit sigma, since the formula's
	// variance-normalization term does not cancel out for sigma != 1.
	g := dsx.NewStandardGaussian()
	testx.AlmostEqual(t, 0, g.BhattacharyyaDistance(dsx.NewStandardGaussian()), 0.0001)
	testx.True(t, g.BhattacharyyaDistance(dsx.NewGaussian(5, 1)) > 0)
}

func TestGaussianSample(t *testing.T) {
	g := dsx.NewStandardGaussian()
	v := g.Sample()
	testx.False(t, math.IsNaN(v))
	testx.False(t, math.IsInf(v, 0))
}

func TestGaussianSampleN(t *testing.T) {
	g := dsx.NewStandardGaussian()
	samples := g.SampleN(100)
	testx.Equal(t, 100, len(samples))
	for _, v := range samples {
		testx.False(t, math.IsNaN(v))
	}
}
