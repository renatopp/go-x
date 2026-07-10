package mathx

// EuclideanDistance calculates the straight-line distance between two points
// (x1, y1) and (x2, y2).
func EuclideanDistance[T Number](x1, y1, x2, y2 T) float64 {
	dx := float64(x2 - x1)
	dy := float64(y2 - y1)
	return Sqrt(dx*dx + dy*dy)
}

// ManhattanDistance calculates the distance between two points (x1, y1) and
// (x2, y2) by summing the absolute differences of their coordinates.
func ManhattanDistance[T SignedNumber](x1, y1, x2, y2 T) T {
	dx := Abs(x2 - x1)
	dy := Abs(y2 - y1)
	return dx + dy
}

// MahalanobisDistance calculates the distance between a point x and the mean mu
// of a distribution with standard deviation sigma. It accounts for the scale of
// the data and is useful for measuring distances in multivariate space.
func MahalanobisDistance[T Number](x T, mu, sigma T) float64 {
	return Sqrt(float64((x - mu) * (x - mu) / (sigma * sigma)))
}

// ChebyshevDistance calculates the distance between two points (x1, y1) and
// (x2, y2) by taking the maximum of the absolute differences of their
// coordinates.
func ChebyshevDistance[T SignedNumber](x1, y1, x2, y2 T) T {
	dx := Abs(x2 - x1)
	dy := Abs(y2 - y1)
	if dx > dy {
		return dx
	}
	return dy
}

// CosineSimilarity calculates the cosine similarity between two points (x1, y1)
// and (x2, y2) by treating them as vectors and measuring the cosine of the angle
// between them. It returns a value between -1 and 1, where 1 means the vectors are
// identical, 0 means they are orthogonal, and -1 means they are opposite.
func CosineSimilarity[T Number](x1, y1, x2, y2 T) float64 {
	dot := float64(x1*x2 + y1*y2)
	mag1 := Sqrt(float64(x1*x1 + y1*y1))
	mag2 := Sqrt(float64(x2*x2 + y2*y2))
	if mag1 == 0 || mag2 == 0 {
		return 0
	}
	return dot / (mag1 * mag2)
}

// HammingDistance calculates the Hamming distance between two points (x1, y1) and
// (x2, y2) by counting the number of coordinates that are different. It returns an
// integer value representing the distance.
func HammingDistance[T Number](x1, y1, x2, y2 T) int {
	var distance int
	if x1 != x2 {
		distance++
	}
	if y1 != y2 {
		distance++
	}
	return distance
}

// MinkowskiDistance calculates the Minkowski distance between two points (x1, y1) and
// (x2, y2) for a given order p. It generalizes the Euclidean and Manhattan distances,
// where p=2 corresponds to Euclidean distance and p=1 corresponds to Manhattan
// distance. It returns a non-negative value representing the distance.
func MinkowskiDistance[T Number](x1, y1, x2, y2 T, p int) float64 {
	pf := float64(p)
	dx := Abs(x2 - x1)
	dy := Abs(y2 - y1)
	return Pow(Pow(dx, pf)+Pow(dy, pf), 1/pf)
}

// CanberraDistance calculates the Canberra distance between two points (x1, y1) and
// (x2, y2) by summing the absolute differences of their coordinates divided by the
// sum of the absolute values of their coordinates. It returns a non-negative value
// representing the distance, where 0 means the points are identical and larger
// values indicate greater distance.
func CanberraDistance[T Number](x1, y1, x2, y2 T) float64 {
	dx := Abs(x2 - x1)
	dy := Abs(y2 - y1)
	sum := Abs(x1) + Abs(y1) + Abs(x2) + Abs(y2)
	if sum == 0 {
		return 0
	}
	return float64((dx + dy) / sum)
}
