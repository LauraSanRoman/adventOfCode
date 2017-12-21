package main

import (
	"fmt"
	"math"
)

func main (){
	calcSteps(1)
	calcSteps(12)
	calcSteps(23)
	calcSteps(1024)
	calcSteps(35)
	calcSteps(20)
	calcSteps(277678)

}
func calcSteps(startNum int) {
	ceilingOfTheSqrt := int(math.Ceil(math.Sqrt(float64(startNum))))
	endOfSquare := 0
	if( ceilingOfTheSqrt % 2 == 0){
		endOfSquare = int(math.Pow(float64(ceilingOfTheSqrt+1), 2))
	}else{
		endOfSquare = int(math.Pow(float64(ceilingOfTheSqrt), 2))
	}

	fmt.Println("end of square " , endOfSquare)
	minNumOfSteps := ceilingOfTheSqrt/2
	x:=endOfSquare
	for x > startNum{
		x = x - minNumOfSteps
	}

	numberOfSteps := (startNum - x) + minNumOfSteps
	fmt.Println("steps : ", numberOfSteps)

}
