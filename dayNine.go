package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"

)

func main () {
  f, _ := os.Open("/Users/Laura/Documents/adventOfCode/input/day9")
  scanner := bufio.NewScanner(f)
  garbageCount := 0
  for scanner.Scan(){
    line := scanner.Text()
    line = removeExclamations(line)
    line,garbageCount = removeGarbage(line)
    //fmt.Println(getScore(line))
    fmt.Println(garbageCount)
  }

}
func removeExclamations(stream string) string{
  returnStream := stream
  for i := 0; i < len(stream); i++ {
    exclamation := strings.Index(returnStream,"!")
    if(exclamation != -1){
      returnStream = returnStream[0:exclamation] + returnStream[exclamation+2:]
    }
  }
  return returnStream
}
func removeGarbage(stream string) (string, int){
  returnStream := stream
  totalGarbageCount := 0
  for i := 0; i < len(stream); i++ {
    startOfGarbage := strings.Index(returnStream,"<")
    endOfGarbage := strings.Index(returnStream,">")
    if(startOfGarbage != -1 && endOfGarbage != -1){
      returnStream = returnStream[0:startOfGarbage] + returnStream[endOfGarbage+1:]
      totalGarbageCount+=endOfGarbage - (startOfGarbage +1)
    }
  }
  return returnStream, totalGarbageCount
}

func getScore(stream string)int{
  score := 0
  level := 0
  for i := 0; i < len(stream); i++ {
    if(stream[i] == '{'){
      level++
    }
    if stream[i] == '}' {
      score+=level
      level--
    }
  }
  return score
}
