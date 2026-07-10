package dsx

import (
	"math"
	"math/rand/v2"
)

// Gaussian represents a normal distribution with mean (Mu) and standard deviation (Sigma).
type Gaussian struct {
	Mu    float64
	Sigma float64
}

// NewGaussian creates a new Gaussian distribution with the specified mean and standard deviation.
func NewGaussian(mu, sigma float64) *Gaussian {
	return &Gaussian{
		Mu:    mu,
		Sigma: sigma,
	}
}

func NewStandardGaussian() *Gaussian {
	return &Gaussian{
		Mu:    0,
		Sigma: 1,
	}
}

// PDF computes the probability density function of the Gaussian distribution at a given point x.
func (g *Gaussian) PDF(x float64) float64 {
	exponent := -0.5 * ((x - g.Mu) / g.Sigma) * ((x - g.Mu) / g.Sigma)
	return (1 / (g.Sigma * math.Sqrt(2*math.Pi))) * math.Exp(exponent)
}

// CDF computes the cumulative distribution function of the Gaussian distribution at a given point x.
func (g *Gaussian) CDF(x float64) float64 {
	return 0.5 * (1 + math.Erf((x-g.Mu)/(g.Sigma*math.Sqrt2)))
}

// Sample generates a random sample from the Gaussian distribution using the Box-Muller transform.
func (g *Gaussian) Sample() float64 {
	u1 := rand.Float64()
	u2 := rand.Float64()
	z0 := math.Sqrt(-2.0*math.Log(u1)) * math.Cos(2.0*math.Pi*u2)
	return g.Mu + z0*g.Sigma
}

// SampleN generates n random samples from the Gaussian distribution.
func (g *Gaussian) SampleN(n int) []float64 {
	samples := make([]float64, n)
	for i := range n {
		samples[i] = g.Sample()
	}
	return samples
}

// Mean returns the mean (Mu) of the Gaussian distribution.
func (g *Gaussian) Mean() float64 {
	return g.Mu
}

// Variance returns the variance (Sigma^2) of the Gaussian distribution.
func (g *Gaussian) Variance() float64 {
	return g.Sigma * g.Sigma
}

// StdDev returns the standard deviation (Sigma) of the Gaussian distribution.
func (g *Gaussian) StdDev() float64 {
	return g.Sigma
}

// Quantile computes the quantile (inverse CDF) of the Gaussian distribution for a given probability p.
func (g *Gaussian) Quantile(p float64) float64 {
	return g.Mu + g.Sigma*math.Sqrt2*math.Erfinv(2*p-1)
}

// LogPDF computes the logarithm of the probability density function of the Gaussian distribution at a given point x.
func (g *Gaussian) LogPDF(x float64) float64 {
	return -0.5*math.Log(2*math.Pi) - math.Log(g.Sigma) - 0.5*((x-g.Mu)/g.Sigma)*((x-g.Mu)/g.Sigma)
}

// LogCDF computes the logarithm of the cumulative distribution function of the Gaussian distribution at a given point x.
func (g *Gaussian) LogCDF(x float64) float64 {
	return math.Log(0.5 * (1 + math.Erf((x-g.Mu)/(g.Sigma*math.Sqrt2))))
}

// Entropy computes the entropy of the Gaussian distribution.
func (g *Gaussian) Entropy() float64 {
	return 0.5 * math.Log(2*math.Pi*math.E*g.Sigma*g.Sigma)
}

// KLDivergence computes the Kullback-Leibler divergence between this Gaussian distribution and another Gaussian distribution.
func (g *Gaussian) KLDivergence(other *Gaussian) float64 {
	variance := g.Sigma * g.Sigma
	otherVariance := other.Sigma * other.Sigma
	return math.Log(other.Sigma/g.Sigma) + (variance+(g.Mu-other.Mu)*(g.Mu-other.Mu))/(2*otherVariance) - 0.5
}

// MahalanobisDistance computes the Mahalanobis distance from a point x to the mean of the Gaussian distribution.
func (g *Gaussian) MahalanobisDistance(x float64) float64 {
	return math.Sqrt((x - g.Mu) * (x - g.Mu) / (g.Sigma * g.Sigma))
}

// BhattacharyyaDistance computes the Bhattacharyya distance between this Gaussian distribution and another Gaussian distribution.
func (g *Gaussian) BhattacharyyaDistance(other *Gaussian) float64 {
	meanDiff := g.Mu - other.Mu
	varianceSum := g.Sigma*g.Sigma + other.Sigma*other.Sigma
	return 0.25*meanDiff*meanDiff/varianceSum + 0.5*math.Log(varianceSum/(2*g.Sigma*g.Sigma*other.Sigma*other.Sigma))
}
