
package main

import (
	"fmt"
	"math/rand"

	"github.com/HenryGeorgist/go-statistics/statistics"
	

)

/*
func main() {
	fmt.Println(quote.Go())
}
*/



// 1. Parameterize a triangular distribution


// 2. generate a random number u[0,1]
func randomNumber() {
	const seed = 54321
	src := rand.NewSource(seed)
	rnd := rand.New(src)

	const N = 1
	var values [N]float64
	
	triDist := make([]statistics.ContinuousDistribution, N)
	triDist[0] = statistics.TriangularDistribution{Min: 10; MostLikely: 20; Max: 30}
	triDist[1] = statistics.TriangularDistribution{Min: 12; MostLikely: 22; Max: 32}

	for i := 0; i < N; i++ {
		probability := rnd.Float64()
		val := statistics.TriangularDistribution.invCDF(probability)
		values[i] := val
	}
}





func
// pull the value from dist based on a random number put into InvCDF


*/
