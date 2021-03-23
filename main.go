package main

import (
	"fmt"
	"math/rand"

	"github.com/HenryGeorgist/go-statistics/data"
	"github.com/HenryGeorgist/go-statistics/statistics"

	"os"
)

func main() {

	comEng()
	comNonEng()

}

func aggregateTriangular(min []float64, mostLikely []float64, max []float64) string {

	const seed = 54321
	src := rand.NewSource(seed)
	rnd := rand.New(src)

	var convergence bool = false
	var N int64 = 1000

	// aggregate non-zero distributions
	if max[1] != 0 {
		triDist1 := statistics.TriangularDistribution{Min: min[0], MostLikely: mostLikely[0], Max: max[0]}
		triDist2 := statistics.TriangularDistribution{Min: min[1], MostLikely: mostLikely[1], Max: max[1]}
		// initialize histogram in which to store aggregated distributions
		histogram := data.Init(1, min[0], max[1])
		// randomly sample each distribution, store in histogram
		for convergence != true {
			var k int64 = 0
			for k < N {
				probability := rnd.Float64()
				val1 := triDist1.InvCDF(probability)
				val2 := triDist2.InvCDF(probability)
				histogram.AddObservation(val1)
				histogram.AddObservation(val2)
				k++
			}
			convergence, N = histogram.TestForConvergence(.05, .95, .95, .01) //upper confidence limit test, lower confidence limit test, confidenece, error tolerance
			fmt.Println(fmt.Sprintf("Computed some, estimated to need %d more iterations", N))
		}

		// N is number of iterations per cycle (sometimes stay at 10 cycles)
		// pull summary stats of aggregated sample

		// return statistics.TriangularDistribution{Min: histogram.InvCDF(0), MostLikely: histogram.InvCDF(.5), Max: histogram.InvCDF(1)}

		// pert distribution
		// geometric mean or median rather than 50% - view data to choose something better
		// dump out each histogram to csv so we can view by bin for each depth

		//instead of return triangular distribution, return the histogram

		return histogram.String()

	} else {
		// zero-valued distributions
		return "0"
	}

}

func comEng() {

	filename := "C:\\Temp\\Richard\\HEC Research\\Go Auxiliary\\comEng.csv"
	w, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer w.Close()

	// 2D array of most likely content damage engineeered structures
	contentDamageFunctionArrayMostLikely := [][]float64{
		{0, 0, 5, 18, 35, 39, 43, 47, 70, 75}, //non-perishable
		{0, 0, 5, 18, 35, 39, 43, 47, 70, 75}} //perishable
	mostLikely := transpose(contentDamageFunctionArrayMostLikely)
	// 2D array of minimum content damage engineered structures
	contentDamageFunctionArrayMin := [][]float64{
		{0, 0, 0, 4, 10, 22, 27, 33, 44, 48}, //non-perishable
		{0, 0, 0, 5, 17, 28, 37, 43, 50, 50}} //perishable
	min := transpose(contentDamageFunctionArrayMin)
	// 2D array of maximum content damage engineered structures
	contentDamageFunctionArrayMax := [][]float64{
		{0, 0, 5, 15, 22, 35, 44, 50, 55, 70}, //non-perishable
		{0, 0, 8, 28, 50, 58, 65, 65, 90, 90}} //perishable
	max := transpose(contentDamageFunctionArrayMax)

	for i := 0; i < len(mostLikely); i++ {
		w.WriteString(aggregateTriangular(min[i], mostLikely[i], max[i]))
	}

}

func comNonEng() {

	filename := "C:\\Temp\\Richard\\HEC Research\\Go Auxiliary\\comNonEng.csv"
	w, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer w.Close()

	// 2D array of most likely content damage engineeered structures
	contentDamageFunctionArrayMostLikely := [][]float64{
		{0, 0, 1, 8, 12, 18, 25, 39, 50, 60},  //non-perishable
		{0, 0, 2, 15, 30, 42, 64, 71, 80, 87}} //perishable
	mostLikely := transpose(contentDamageFunctionArrayMostLikely)
	// 2D array of minimum content damage engineered structures
	contentDamageFunctionArrayMin := [][]float64{
		{0, 0, 0, 3, 7, 13, 20, 30, 40, 45}, //non-perishable
		{0, 0, 0, 5, 9, 15, 23, 30, 35, 41}} //perishable
	min := transpose(contentDamageFunctionArrayMin)
	// 2D array of maximum content damage engineered structures
	contentDamageFunctionArrayMax := [][]float64{
		{0, 0, 4, 18, 28, 38, 49, 64, 72, 90},   //non-perishable
		{0, 0, 10, 35, 54, 65, 84, 95, 99, 100}} //perishable
	max := transpose(contentDamageFunctionArrayMax)

	for i := 0; i < len(mostLikely); i++ {
		w.WriteString(aggregateTriangular(min[i], mostLikely[i], max[i]))
	}

}

func transpose(slice [][]float64) [][]float64 {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]float64, xl)
	for i := range result {
		result[i] = make([]float64, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}
