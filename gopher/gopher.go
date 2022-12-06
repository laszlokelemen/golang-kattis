package main

import (
	"fmt"
	"math"
)

func main() {
	var gopherX, gopherY, dogX, dogY, holeX, holeY float64
	_, _ = fmt.Scanf("%f %f %f %f\n", &gopherX, &gopherY, &dogX, &dogY)

	for {
		_, err := fmt.Scanf("%f %f\n", &holeX, &holeY)
		if err != nil {
			fmt.Println("The gopher cannot escape.")
			return
		}
		if math.Sqrt(math.Pow(gopherX-holeX, 2)+math.Pow(gopherY-holeY, 2)) <= 0.5*math.Sqrt(math.Pow(dogX-holeX, 2)+math.Pow(dogY-holeY, 2)) {
			fmt.Printf("The gopher can escape through the hole at (%.3f,%.3f).\n", holeX, holeY)
			return
		}
	}
}
