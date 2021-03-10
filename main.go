package main

import (
	"fmt"
	"math/rand"

	"github.com/HenryGeorgist/go-statistics/data"
	"github.com/HenryGeorgist/go-statistics/statistics"
)

// 1. Parameterize a triangular distribution

// 2. generate a random number u[0,1]
func main() {
	const seed = 54321
	src := rand.NewSource(seed)
	rnd := rand.New(src)

	const N = 1000
	var values [N]float64

	triDist := statistics.TriangularDistribution{Min: 10, MostLikely: 21, Max: 32}

	for i := 0; i < N; i++ {
		probability := rnd.Float64()
		val := triDist.InvCDF(probability)
		values[i] = val
	}

	histogram := data.Init(2, 10, 32)
	for idx := range values {
		histogram.AddObservation(values[idx])
	}

	min := histogram.InvCDF(0)
	max := histogram.InvCDF(1)
	median := histogram.InvCDF(.5)

	fmt.Println(fmt.Sprintf("The min is %f, the max is %f, and the most likely is %f", min, max, median))

}
