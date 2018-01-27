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
  for scanner.Scan(){
    line := scanner.Text()
    line = removeExclamations(line)
    line = removeGarbage(line)
    fmt.Println(getLevel(line))
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
func removeGarbage(stream string) string{
  returnStream := stream
  for i := 0; i < len(stream); i++ {
    startOfGarbage := strings.Index(returnStream,"<")
    endOfGarbage := strings.Index(returnStream,">")
    if(startOfGarbage != -1 && endOfGarbage != -1){
      returnStream = returnStream[0:startOfGarbage] + returnStream[endOfGarbage+1:]
    }
  }
  return returnStream
}

func getLevel(stream string)int{
  level := 1
  opening := strings.Index(stream,"{")
  closing := strings.Index(stream,"}")
  lastClosing := strings.LastIndex(stream,"}")
  if(closing == opening+1 || lastClosing != len(stream)-1){
    getGroupCount(stream)
    return level
  }else{
    return level + getLevel(stream[opening+1:lastClosing])
  }
}

func getGroupCount(stream string) {
  groupings := strings.Split(stream,",")
  for i := 0; i < len(groupings); i++ {
    if(groupings[i] != "{}"){
      fmt.Println("level", getLevel(groupings[i]))
    }
  }
  fmt.Println("Group count:", len(groupings))
}
