package kmeans

/*
This module provides common distance functions for measuring distance
between observations.

Minkowski Distance is one the most inclusive among one as other distances are only
a specific case of Minkowski Distance(Chebyshev Distance is not straightforward, though).

when p=1 in MinkowskiDistance, it becomes ManhattanDistance,
when p=2 in MinkowskiDistance, it becomes EuclideanDIstance,
when p goes infinity, it becomes ChebyshevDistance.

Since the ManhattanDistance and EuclideanDistance are very frequently used, they are
implemented separately.
*/

import (
	"math"
)

// Lp Norm of an array, given p >= 1
func LPNorm(vector []float64, p float64) (float64, error) {
	distance := 0.
	for _, jj := range vector {
		distance += math.Pow(math.Abs(jj), p)
	}
	return math.Pow(distance, 1/p), nil
}

// 1-norm distance (l_1 distance)
func ManhattanDistance(firstVector, secondVector []float64) (float64, error) {
	distance := 0.
	for ii := range firstVector {
		distance += math.Abs(firstVector[ii] - secondVector[ii])
	}
	return distance, nil
}

// 2-norm distance (l_2 distance)
func EuclideanDistance(firstVector, secondVector []float64) (float64, error) {
	distance := 0.
	for ii := range firstVector {
		distance += (firstVector[ii] - secondVector[ii]) * (firstVector[ii] - secondVector[ii])
	}
	return math.Sqrt(distance), nil
}

// Higher weight for the points that are far apart
// Not a real metric as it does not obey triangle inequality
func SquaredEuclideanDistance(firstVector, secondVector []float64) (float64, error) {
	distance, err := EuclideanDistance(firstVector, secondVector)
	return distance * distance, err
}

// p-norm distance (l_p distance)
func MinkowskiDistance(firstVector, secondVector []float64, p float64) (float64, error) {
	distance := 0.
	for ii := range firstVector {
		distance += math.Pow(math.Abs(firstVector[ii]-secondVector[ii]), p)
	}
	return math.Pow(distance, 1/p), nil
}

// p-norm distance with weights (weighted l_p distance)
func WeightedMinkowskiDistance(firstVector, secondVector, weightVector []float64, p float64) (float64, error) {
	distance := 0.
	for ii := range firstVector {
		distance += weightVector[ii] * math.Pow(math.Abs(firstVector[ii]-secondVector[ii]), p)
	}
	return math.Pow(distance, 1/p), nil
}

// infinity norm distance (l_inf distance)
func ChebyshevDistance(firstVector, secondVector []float64) (float64, error) {
	distance := 0.
	for ii := range firstVector {
		if math.Abs(firstVector[ii]-secondVector[ii]) >= distance {
			distance = math.Abs(firstVector[ii] - secondVector[ii])
		}
	}
	return distance, nil
}

func HammingDistance(firstVector, secondVector []float64) (float64, error) {
	distance := 0.
	for ii := range firstVector {
		if firstVector[ii] != secondVector[ii] {
			distance++
		}
	}
	return distance, nil
}

func BrayCurtisDistance(firstVector, secondVector []float64) (float64, error) {
	numerator, denominator := 0., 0.
	for ii := range firstVector {
		numerator += math.Abs(firstVector[ii] - secondVector[ii])
		denominator += math.Abs(firstVector[ii] + secondVector[ii])
	}
	return numerator / denominator, nil
}

func CanberraDistance(firstVector, secondVector []float64) (float64, error) {
	distance := 0.
	for ii := range firstVector {
		distance += (math.Abs(firstVector[ii]-secondVector[ii]) / (math.Abs(firstVector[ii]) + math.Abs(secondVector[ii])))
	}
	return distance, nil
}
