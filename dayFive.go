package main

import (
  "fmt"
  "os"
  "bufio"
  "math"
  "strconv"
)

func main(){
  f, _ := os.Open("/Users/hpw355/Documents/AdventOfCode/input/day5")
  scanner := bufio.NewScanner(f)
  var input []int
  for scanner.Scan(){
    number, _ := strconv.Atoi(scanner.Text())
    input = append(input, number)
  }
    fmt.Println(countJumps(input))
}

func countJumps(list []int) int{
  currentPosition :=0
  nextPostion := 0
  jumps := 0
  for i := 0.0; i < math.Inf(1); i++ {
    nextPostion = currentPosition + list[currentPosition]
    if(nextPostion < len(list)){
      if(list[currentPosition] >= 3){
        list[currentPosition]--
        }else{
          list[currentPosition]++
        }
      jumps++
      currentPosition = nextPostion
    }else{
      return jumps+1
    }
  }
  return jumps
}
