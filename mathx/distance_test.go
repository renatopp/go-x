package mathx_test

import (
	"testing"

	"github.com/renatopp/go-x/mathx"
	"github.com/renatopp/go-x/testx"
)

func TestEuclideanDistance(t *testing.T) {
	testx.AlmostEqual(t, 5.0, mathx.EuclideanDistance(0, 0, 3, 4), 0.0001)
	testx.AlmostEqual(t, 0.0, mathx.EuclideanDistance(1, 1, 1, 1), 0.0001)
	testx.AlmostEqual(t, 11.40175425099138, mathx.EuclideanDistance(-1, -1, 6, 8), 0.0001)
}

func TestManhattanDistance(t *testing.T) {
	testx.Equal(t, 7, mathx.ManhattanDistance(0, 0, 3, 4))
	testx.Equal(t, 0, mathx.ManhattanDistance(1, 1, 1, 1))
	testx.Equal(t, 16, mathx.ManhattanDistance(-1, -1, 6, 8))
}

func TestMahallanobisDistance(t *testing.T) {
	testx.AlmostEqual(t, 2.0, mathx.MahalanobisDistance(5, 3, 1), 0.0001)
	testx.AlmostEqual(t, 0.0, mathx.MahalanobisDistance(1, 1, 1), 0.0001)
	testx.AlmostEqual(t, 4.0, mathx.MahalanobisDistance(10, 6, 1), 0.0001)
}

func TestChebyshevDistance(t *testing.T) {
	testx.Equal(t, 4, mathx.ChebyshevDistance(0, 0, 3, 4))
	testx.Equal(t, 0, mathx.ChebyshevDistance(1, 1, 1, 1))
	testx.Equal(t, 9, mathx.ChebyshevDistance(-1, -1, 6, 8))
}

func TestCosineSimilarity(t *testing.T) {
	testx.AlmostEqual(t, 1.0, mathx.CosineSimilarity(1, 0, 1, 0), 0.0001)
	testx.AlmostEqual(t, 0.0, mathx.CosineSimilarity(1, 0, 0, 1), 0.0001)
	testx.AlmostEqual(t, -1.0, mathx.CosineSimilarity(1, 0, -1, 0), 0.0001)
}

func TestHammingDistance(t *testing.T) {
	testx.Equal(t, 2, mathx.HammingDistance(0, 0, 3, 4))
	testx.Equal(t, 0, mathx.HammingDistance(1, 1, 1, 1))
	testx.Equal(t, 2, mathx.HammingDistance(-1, -1, 6, 8))
}

func TestMinkowskiDistance(t *testing.T) {
	testx.AlmostEqual(t, 4.497941445275414, mathx.MinkowskiDistance(0, 0, 3, 4, 3), 0.0001)
	testx.AlmostEqual(t, 0.0, mathx.MinkowskiDistance(1, 1, 1, 1, 3), 0.0001)
	testx.AlmostEqual(t, 10.2344598938, mathx.MinkowskiDistance(-1, -1, 6, 8, 3), 0.0001)
}

func TestCanberraDistance(t *testing.T) {
	testx.AlmostEqual(t, 1.0, mathx.CanberraDistance(0, 0, 3, 4), 0.0001)
	testx.AlmostEqual(t, 0.0, mathx.CanberraDistance(1, 1, 1, 1), 0.0001)
	testx.AlmostEqual(t, 1.0, mathx.CanberraDistance(-1, -1, 6, 8), 0.0001)
}
