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

	var contentDamageFunctionArrayAvg [][]float64
	var contentDamageFunctionArrayMin [][]float64
	var contentDamageFunctionArrayMax [][]float64

	engineeredNonperishableAvg := []float64{0, 0, 5, 18, 35, 39, 43, 47, 70, 75}
	contentDamageFunctionArrayAvg = append(contentDamageFunctionArrayAvg, engineeredNonperishableAvg)
	engineeredNonperishableMin := []float64{0, 0, 0, 4, 10, 22, 27, 33, 44, 48}
	contentDamageFunctionArrayMin = append(contentDamageFunctionArrayMin, engineeredNonperishableMin)
	engineeredNonperishableMax := []float64{0, 0, 5, 15, 22, 35, 44, 50, 55, 70}
	contentDamageFunctionArrayMax = append(contentDamageFunctionArrayMax, engineeredNonperishableMax)
	engineeredPerishableAvg := []float64{0, 0, 5, 18, 35, 39, 43, 47, 70, 75}
	contentDamageFunctionArrayAvg = append(contentDamageFunctionArrayAvg, engineeredPerishableAvg)
	engineeredPerishableMin := []float64{0, 0, 0, 5, 17, 28, 37, 43, 50, 50}
	contentDamageFunctionArrayMin = append(contentDamageFunctionArrayMin, engineeredPerishableMin)
	engineeredPerishableMax := []float64{0, 0, 8, 28, 50, 58, 65, 65, 90, 90}
	contentDamageFunctionArrayMax = append(contentDamageFunctionArrayMax, engineeredPerishableMax)

	var aggregateArray [10][3]float64

	for j := 0; j < 10; j++ {

		triDist1 := statistics.TriangularDistribution{Min: contentDamageFunctionArrayMin[0][j], MostLikely: contentDamageFunctionArrayAvg[0][j], Max: contentDamageFunctionArrayMax[0][j]}
		triDist2 := statistics.TriangularDistribution{Min: contentDamageFunctionArrayMin[1][j], MostLikely: contentDamageFunctionArrayAvg[1][j], Max: contentDamageFunctionArrayMax[1][j]}

		/*
			maxOfMax := contentDamageFunctionArrayMax[1][j]
			minOfMin := contentDamageFunctionArrayMin[0][j]
			damageRange := maxOfMax - minOfMin + 1

			nBins := float64(int(damageRange / 20))
		*/

		histogram := data.Init(1, contentDamageFunctionArrayMin[0][j], contentDamageFunctionArrayMax[1][j])

		for i := 0; i < N; i++ {
			probability := rnd.Float64()
			val1 := triDist1.InvCDF(probability)
			val2 := triDist2.InvCDF(probability)
			histogram.AddObservation(val1)
			histogram.AddObservation(val2)
		}

		aggregateArray[j][0] = histogram.InvCDF(0)
		aggregateArray[j][1] = histogram.InvCDF(.5)
		aggregateArray[j][2] = histogram.InvCDF(1)

		fmt.Println(aggregateArray[j])

	}
}
