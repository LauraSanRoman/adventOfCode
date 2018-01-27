package main

import(
  "fmt"
  "strings"
  "os"
  "bufio"
)

func main(){
  f, _ := os.Open("/Users/Laura/Documents/adventOfCode/input/day7")
  scanner := bufio.NewScanner(f)
  var listOfStarters []string
  var listOfConnections [][]string
  for scanner.Scan(){
    line := scanner.Text()
    if(strings.ContainsAny(line,"->")){
      indexOfParen := strings.Index(line,"(")
      indexOfArrow := strings.Index(line, ">")
      listOfStarters = append(listOfStarters, strings.Trim(line[0:indexOfParen]," "))
      listOfConnections = append(listOfConnections, strings.Split(line[indexOfArrow+1:],", "))
    }
  }
  for start := range listOfStarters {
    for listOfConnects := range listOfConnections {
      for connect := range listOfConnections[listOfConnects] {
        if(listOfStarters[start] == strings.Trim(listOfConnections[listOfConnects][connect]," ")){
          listOfStarters[start] = "0"
        }
      }
    }
  }

  for element := range listOfStarters {
    if (listOfStarters[element] != "0") {
      fmt.Println(listOfStarters[element])
    }
  }
}
