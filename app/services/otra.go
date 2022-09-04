package services

import "math"

func KnapSack(W float64, wt []float64, val []float64, n int) float64 {
	if n == 0 || W == 0 {
		return 0
	} else {
		return math.Max(val[n-1]+KnapSack(W-wt[n-1], wt, val, n-1), KnapSack(W, wt, val, n-1))
	}
}
