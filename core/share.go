package core

import "math"

func CalcInitialShare(reserveA, reserveB float64) float64 {
	initialShare := math.Sqrt(reserveA * reserveB)
	return initialShare
}

func CalcShare(deposited, reserveStarting, shareStarting float64) float64 {
	shareMinted := deposited / reserveStarting * shareStarting
	return shareMinted
}
