package main

import(
  "fmt"
  "os"
  "bufio"
)

func main(){
    var listOfDistros = []string{}
  	f, _ := os.Open("/Users/Laura/Documents/adventOfCode/listOfArrays.txt")
  	scanner := bufio.NewScanner(f)
  	for scanner.Scan(){
  		line := scanner.Text()
      if(checkIfRepeat(listOfDistros,line)){
        fmt.Println("repeated at length: ", len(listOfDistros))
        fmt.Println(line)
        fmt.Println(listOfDistros)
        break
      }
      listOfDistros = append(listOfDistros, line)
    }
}


func checkIfRepeat(list []string, element string) bool {
  for item := range list {
    if (list[item] == element) {
      return true
    }
  }
  return false
}
