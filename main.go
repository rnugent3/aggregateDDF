package main

import (
	"fmt"
	"math/rand"

	"github.com/HenryGeorgist/go-statistics/data"
	"github.com/HenryGeorgist/go-statistics/statistics"
)

func main() {

	aggregateTriangular()

}

func aggregateTriangular() {

	const seed = 54321
	src := rand.NewSource(seed)
	rnd := rand.New(src)

	const N = 1000
	// initialize 2D arrays to hold most likely, min, and max damage functions
	var contentDamageFunctionArrayAvg [][]float64
	var contentDamageFunctionArrayMin [][]float64
	var contentDamageFunctionArrayMax [][]float64

	// 2D array of most likely content damage engineeered structures
	engineeredNonperishableAvg := []float64{0, 0, 5, 18, 35, 39, 43, 47, 70, 75}
	contentDamageFunctionArrayAvg = append(contentDamageFunctionArrayAvg, engineeredNonperishableAvg)
	engineeredPerishableAvg := []float64{0, 0, 5, 18, 35, 39, 43, 47, 70, 75}
	contentDamageFunctionArrayAvg = append(contentDamageFunctionArrayAvg, engineeredPerishableAvg)
	// 2D array of most likely content damage non-engineered structures
	nonEngineeredNonperishableAvg := []float64{0, 0, 1, 8, 12, 18, 25, 39, 50, 60}
	contentDamageFunctionArrayAvg = append(contentDamageFunctionArrayAvg, nonEngineeredNonperishableAvg)
	nonEngineeredPerishableAvg := []float64{0, 0, 2, 15, 30, 42, 64, 71, 80, 87}
	contentDamageFunctionArrayAvg = append(contentDamageFunctionArrayAvg, nonEngineeredPerishableAvg)

	// 2D array of minimum content damage engineered structures
	engineeredNonperishableMin := []float64{0, 0, 0, 4, 10, 22, 27, 33, 44, 48}
	contentDamageFunctionArrayMin = append(contentDamageFunctionArrayMin, engineeredNonperishableMin)
	engineeredPerishableMin := []float64{0, 0, 0, 5, 17, 28, 37, 43, 50, 50}
	contentDamageFunctionArrayMin = append(contentDamageFunctionArrayMin, engineeredPerishableMin)
	// 2D array of minimum content damage non-engineered structures
	nonEngineeredNonperishableMin := []float64{0, 0, 0, 3, 7, 13, 20, 30, 40, 45}
	contentDamageFunctionArrayMin = append(contentDamageFunctionArrayMin, nonEngineeredNonperishableMin)
	nonEngineeredPerishableMin := []float64{0, 0, 0, 5, 9, 15, 23, 30, 35, 41}
	contentDamageFunctionArrayMin = append(contentDamageFunctionArrayMin, nonEngineeredPerishableMin)

	// 2D array of maximum content damage engineered structures
	engineeredNonperishableMax := []float64{0, 0, 5, 15, 22, 35, 44, 50, 55, 70}
	contentDamageFunctionArrayMax = append(contentDamageFunctionArrayMax, engineeredNonperishableMax)
	engineeredPerishableMax := []float64{0, 0, 8, 28, 50, 58, 65, 65, 90, 90}
	contentDamageFunctionArrayMax = append(contentDamageFunctionArrayMax, engineeredPerishableMax)
	// 2D array of maximum content damage non-engineered structures
	nonEngineeredNonperishableMax := []float64{0, 0, 4, 18, 28, 38, 49, 64, 72, 90}
	contentDamageFunctionArrayMax = append(contentDamageFunctionArrayMax, nonEngineeredNonperishableMax)
	nonEngineeredPerishableMax := []float64{0, 0, 10, 35, 54, 65, 84, 95, 99, 100}
	contentDamageFunctionArrayMax = append(contentDamageFunctionArrayMax, nonEngineeredPerishableMax)

	// initialize array to hold aggregated damage functions
	var aggregateArray [3][10][3]float64

	// loop through building type (engineered then non-engineered)
	for i := 0; i < len(contentDamageFunctionArrayAvg)-1; i = i + 2 {
		// loop through depths, initialize triangular dist for each depth by building type

		for j := 0; j < 10; j++ {
			// aggregate non-zero distributions
			if contentDamageFunctionArrayMax[i+1][j] != 0 {
				triDist1 := statistics.TriangularDistribution{Min: contentDamageFunctionArrayMin[i][j], MostLikely: contentDamageFunctionArrayAvg[i][j], Max: contentDamageFunctionArrayMax[i][j]}
				triDist2 := statistics.TriangularDistribution{Min: contentDamageFunctionArrayMin[i+1][j], MostLikely: contentDamageFunctionArrayAvg[i+1][j], Max: contentDamageFunctionArrayMax[i+1][j]}
				// initialize histogram in which to store aggregated distributions
				histogram := data.Init(1, contentDamageFunctionArrayMin[i][j], contentDamageFunctionArrayMax[i+1][j])
				// randomly sample each distribution, store in histogram
				for k := 0; k < N; k++ {
					probability := rnd.Float64()
					val1 := triDist1.InvCDF(probability)
					val2 := triDist2.InvCDF(probability)
					histogram.AddObservation(val1)
					histogram.AddObservation(val2)
				}
				// pull summary stats of aggregated sample
				// this might be implemented using a struct, depends on transfer of data to go-consequences
				aggregateArray[i][j][0] = histogram.InvCDF(0)
				aggregateArray[i][j][1] = histogram.InvCDF(.5)
				aggregateArray[i][j][2] = histogram.InvCDF(1)
			} else {
				// zero-valued distributions
				aggregateArray[i][j][0] = 0
				aggregateArray[i][j][1] = 0
				aggregateArray[i][j][2] = 0
			}
			// print for verification of logic
			fmt.Println(aggregateArray[i][j])
		}

	}

}
