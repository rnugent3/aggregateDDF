package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"

	"github.com/HydrologicEngineeringCenter/go-statistics/data"
	"github.com/HydrologicEngineeringCenter/go-statistics/statistics"

	"os"
)

// to-do:
// commercial structures
// finished/unfinished basements (uniform?)
func main() {
	commercialContentCurve()
	commericalStructureCurve()
	res1BigWave()
	res1Freshwater()
	res1MediumWave()
	res1Saltwater()
	res2BigWave()
	res2Freshwater()
	res2MediumWave()
	res2Saltwater()
}

func aggregateUniform(min float64, max float64) string {
	const seed = 54321
	src := rand.NewSource(seed)
	rnd := rand.New(src)

	var convergence bool = false
	var N int64 = 1000
	if max != 0 {
		if min != max {
			uniDist1 := statistics.UniformDistribution{Min: min, Max: max}
			uniDist2 := statistics.TriangularDistribution{Min: min, Max: max}
			// initialize histogram in which to store aggregated distributions
			histogram := data.Init(1, min, max)
			// randomly sample each distribution, store in histogram
			for convergence != true {
				var k int64 = 0
				for k < N {
					probability := rnd.Float64()
					val1 := uniDist1.InvCDF(probability)
					val2 := uniDist2.InvCDF(probability)
					histogram.AddObservation(val1)
					histogram.AddObservation(val2)
					k++
				}
				convergence, N = histogram.TestForConvergence(.05, .95, .95, .01) //upper confidence limit test, lower confidence limit test, confidenece, error tolerance
				fmt.Println(fmt.Sprintf("Computed some, estimated to need %d more iterations", N))
			}

			return histogram.String()
		} else {
			f := strconv.FormatFloat(min, 'f', -1, 64)
			return f + ","
		}

	} else {
		// zero-valued distributions
		return "0"
	}
}

func aggregateTriangular(min []float64, mostLikely []float64, max []float64) string {

	const seed = 54321
	src := rand.NewSource(seed)
	rnd := rand.New(src)

	var convergence bool = false
	var N int64 = 1000

	if len(max) == 2 {
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

			return histogram.String()

		} else {
			// zero-valued distributions
			return "0"
		}
	} else if len(max) == 4 {
		if max[1] != 0 {
			triDist1 := statistics.TriangularDistribution{Min: min[0], MostLikely: mostLikely[0], Max: max[0]}
			triDist2 := statistics.TriangularDistribution{Min: min[1], MostLikely: mostLikely[1], Max: max[1]}
			triDist3 := statistics.TriangularDistribution{Min: min[2], MostLikely: mostLikely[2], Max: max[2]}
			triDist4 := statistics.TriangularDistribution{Min: min[3], MostLikely: mostLikely[3], Max: max[3]}
			// initialize histogram in which to store aggregated distributions
			histogram := data.Init(1, min[0], max[1])
			// randomly sample each distribution, store in histogram
			for convergence != true {
				var k int64 = 0
				for k < N {
					probability := rnd.Float64()
					val1 := triDist1.InvCDF(probability)
					val2 := triDist2.InvCDF(probability)
					val3 := triDist3.InvCDF(probability)
					val4 := triDist4.InvCDF(probability)
					histogram.AddObservation(val1)
					histogram.AddObservation(val2)
					histogram.AddObservation(val3)
					histogram.AddObservation(val4)
					k++
				}
				convergence, N = histogram.TestForConvergence(.05, .95, .95, .01) //upper confidence limit test, lower confidence limit test, confidenece, error tolerance
				fmt.Println(fmt.Sprintf("Computed some, estimated to need %d more iterations", N))
			}

			return histogram.String()

		} else {
			// zero-valued distributions
			return "0"
		}
	} else {
		return "Slice length not handled"
	}

}

func commercialContentCurve() {
	filename := "/Users/rxjxnx3x/Dropbox/USACE Employment/HEC/Code inputs/commercialContents.csv"
	w, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer w.Close()
	// 2D array of most likely content damage engineeered structures
	contentDamageFunctionArrayMostLikely := [][]float64{
		{0, 0, 5, 18, 35, 39, 43, 47, 70, 75}, //non-perishable engineered
		{0, 0, 5, 18, 35, 39, 43, 47, 70, 75}, //perishable engineered
		{0, 0, 1, 8, 12, 18, 25, 39, 50, 60},  //non-perishable non-engineered
		{0, 0, 2, 15, 30, 42, 64, 71, 80, 87}} //perishable non-engineered
	mostLikely := transpose(contentDamageFunctionArrayMostLikely)
	// 2D array of minimum content damage engineered structures
	contentDamageFunctionArrayMin := [][]float64{
		{0, 0, 0, 4, 10, 22, 27, 33, 44, 48}, //non-perishable engineered
		{0, 0, 0, 5, 17, 28, 37, 43, 50, 50}, //perishable engineered
		{0, 0, 0, 3, 7, 13, 20, 30, 40, 45},  //non-perishable non-engineered
		{0, 0, 0, 5, 9, 15, 23, 30, 35, 41}}  //perishable non-engineered
	min := transpose(contentDamageFunctionArrayMin)
	// 2D array of maximum content damage engineered structures
	contentDamageFunctionArrayMax := [][]float64{
		{0, 0, 5, 15, 22, 35, 44, 50, 55, 70},   //non-perishable engineered
		{0, 0, 8, 28, 50, 58, 65, 65, 90, 90},   //perishable engineered
		{0, 0, 4, 18, 28, 38, 49, 64, 72, 90},   //non-perishable non-engineered
		{0, 0, 10, 35, 54, 65, 84, 95, 99, 100}} //perishable non-engineered
	max := transpose(contentDamageFunctionArrayMax)

	for i := 0; i < len(mostLikely); i++ {
		w.WriteString(aggregateTriangular(min[i], mostLikely[i], max[i]))
	}
}

func res1Freshwater() {
	filename := "/Users/rxjxnx3x/Dropbox/USACE Employment/HEC/Code inputs/res1Freshwater.csv"
	w, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer w.Close()
	unfinished := []float64{0, 0, 3, 5, 14, 27, 35, 39, 44, 48, 53, 56, 59, 62, 64, 66, 68, 69, 71, 73, 74}
	finished := []float64{0, 0, 13, 15, 24, 37, 45, 49, 54, 58, 63, 66, 69, 72, 74, 76, 78, 79, 81, 83, 84}
	for i := 0; i < len(finished); i++ {
		min := math.Min(unfinished[i], finished[i])
		max := math.Max(unfinished[i], finished[i])
		w.WriteString(aggregateUniform(min, max))
	}

}

func res1Saltwater() {
	filename := "/Users/rxjxnx3x/Dropbox/USACE Employment/HEC/Code inputs/res1Saltwater.csv"
	w, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer w.Close()
	unfinished := []float64{0, 0, 5, 7, 16, 34, 43, 49, 56, 61, 68, 71, 76, 80, 82, 84, 86, 89, 91, 93, 94}
	finished := []float64{0, 0, 15, 17, 26, 44, 53, 59, 66, 71, 78, 81, 86, 90, 92, 94, 96, 99, 100, 100, 100}
	for i := 0; i < len(finished); i++ {
		min := math.Min(unfinished[i], finished[i])
		max := math.Max(unfinished[i], finished[i])
		w.WriteString(aggregateUniform(min, max))
	}
}

func res1MediumWave() {
	filename := "/Users/rxjxnx3x/Dropbox/USACE Employment/HEC/Code inputs/res1MediumWave.csv"
	w, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer w.Close()
	unfinished := []float64{8, 9, 13, 16, 29, 44, 53, 60, 67, 74, 80, 84, 89, 92, 93, 95, 96, 97, 98, 99, 100}
	finished := []float64{8, 9, 18, 21, 34, 49, 58, 65, 72, 79, 85, 89, 94, 97, 98, 100, 100, 100, 100, 100, 100}
	for i := 0; i < len(finished); i++ {
		min := math.Min(unfinished[i], finished[i])
		max := math.Max(unfinished[i], finished[i])
		w.WriteString(aggregateUniform(min, max))
	}
}

func res1BigWave() {
	filename := "/Users/rxjxnx3x/Dropbox/USACE Employment/HEC/Code inputs/res1BigWave.csv"
	w, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer w.Close()
	unfinished := []float64{16, 18, 20, 25, 43, 55, 63, 71, 78, 87, 91, 97, 100, 100, 100, 100, 100, 100, 100, 100, 100}
	finished := []float64{16, 18, 20, 25, 43, 55, 63, 71, 78, 87, 91, 97, 100, 100, 100, 100, 100, 100, 100, 100, 100}
	for i := 0; i < len(finished); i++ {
		min := math.Min(unfinished[i], finished[i])
		max := math.Max(unfinished[i], finished[i])
		w.WriteString(aggregateUniform(min, max))
	}
}

func res2Freshwater() {
	filename := "/Users/rxjxnx3x/Dropbox/USACE Employment/HEC/Code inputs/res2Freshwater.csv"
	w, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer w.Close()
	unfinished := []float64{0, 0, 2, 3, 9, 20, 25, 29, 32, 35, 39, 41, 44, 46, 47, 49, 50, 51, 53, 54, 55}
	finished := []float64{0, 0, 12, 13, 19, 30, 35, 39, 42, 45, 49, 51, 54, 56, 57, 59, 60, 61, 63, 64, 65}
	for i := 0; i < len(finished); i++ {
		min := math.Min(unfinished[i], finished[i])
		max := math.Max(unfinished[i], finished[i])
		w.WriteString(aggregateUniform(min, max))
	}
}

func res2Saltwater() {
	filename := "/Users/rxjxnx3x/Dropbox/USACE Employment/HEC/Code inputs/res2Saltwater.csv"
	w, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer w.Close()
	unfinished := []float64{0, 0, 5, 7, 11, 25, 32, 36, 41, 45, 50, 53, 56, 59, 61, 63, 64, 66, 68, 69, 70}
	finished := []float64{0, 0, 15, 17, 21, 35, 42, 46, 51, 55, 60, 63, 66, 69, 71, 73, 74, 76, 78, 79, 80}
	for i := 0; i < len(finished); i++ {
		min := math.Min(unfinished[i], finished[i])
		max := math.Max(unfinished[i], finished[i])
		w.WriteString(aggregateUniform(min, max))
	}
}

func res2MediumWave() {
	filename := "/Users/rxjxnx3x/Dropbox/USACE Employment/HEC/Code inputs/res2MediumWave.csv"
	w, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer w.Close()
	unfinished := []float64{6, 7, 11, 14, 23, 38, 50, 60, 70, 77, 79, 81, 82, 84, 84, 85, 86, 87, 88, 89, 89}
	finished := []float64{6, 7, 16, 19, 28, 43, 55, 65, 75, 82, 84, 86, 87, 89, 89, 90, 91, 92, 93, 94, 94}
	for i := 0; i < len(finished); i++ {
		min := math.Min(unfinished[i], finished[i])
		max := math.Max(unfinished[i], finished[i])
		w.WriteString(aggregateUniform(min, max))
	}
}

func res2BigWave() {
	filename := "/Users/rxjxnx3x/Dropbox/USACE Employment/HEC/Code inputs/res2BigWave.csv"
	w, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer w.Close()
	unfinished := []float64{12, 14, 16, 21, 35, 51, 68, 84, 98, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100}
	finished := []float64{12, 14, 16, 21, 35, 51, 68, 84, 98, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100}
	for i := 0; i < len(finished); i++ {
		min := math.Min(unfinished[i], finished[i])
		max := math.Max(unfinished[i], finished[i])
		w.WriteString(aggregateUniform(min, max))
	}
}

func commericalStructureCurve() {
	filename := "/Users/rxjxnx3x/Dropbox/USACE Employment/HEC/Code inputs/commercialStructure.csv"
	w, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer w.Close()
	// 2D array of most likely content damage engineeered structures
	structureDamageFunctionArrayMostLikely := [][]float64{
		{0, 0, 5, 10, 20, 30, 35, 40, 53, 58}, // engineered
		{0, 0, 5, 12, 20, 28, 35, 45, 55, 60}} // non-engineered
	mostLikely := transpose(structureDamageFunctionArrayMostLikely)
	// 2D array of minimum content damage engineered structures
	structureDamageFunctionArrayMin := [][]float64{
		{0, 0, 0, 5, 12, 18, 28, 33, 43, 48}, // engineered
		{0, 0, 0, 5, 10, 15, 20, 28, 35, 40}} // non-engineered
	min := transpose(structureDamageFunctionArrayMin)
	// 2D array of maximum content damage engineered structures
	structureDamageFunctionArrayMax := [][]float64{
		{0, 0, 9, 17, 27, 36, 43, 48, 60, 69},   // engineered
		{0, 10, 15, 20, 30, 42, 55, 65, 75, 78}} // non-engineered
	max := transpose(structureDamageFunctionArrayMax)

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
