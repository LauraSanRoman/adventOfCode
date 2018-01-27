package main

import(
  "fmt"
  "os"
  "bufio"
)

func main(){
  distro := []int {5,1,10,0,1,7,13,14,3,12,8,10,7,12,0,6}
  listOfDistros := []string{""}
  for i := 1; i < 100000; i++ {
    max, maxIndex := maxOfList(distro)
    addingToIndex := maxIndex+1
    distro[maxIndex] = 0
    for i := max; i > 0; i-- {
      if(addingToIndex == len(distro)){
        addingToIndex = 0
      }
      distro[addingToIndex]++
      addingToIndex++
    }
    fmt.Println(distro)
  }
  	f, _ := os.Open("/Users/hpw355/Documents/AdventOfCode/listOfArrays.txt")
  	scanner := bufio.NewScanner(f)
  	for scanner.Scan(){
  		line := scanner.Text()
      if(checkIfRepeat(listOfDistros,line)){
        fmt.Println("repeated at length: ", len(listOfDistros))
      }
      listOfDistros = append(listOfDistros, line)
    }
    fmt.Println(listOfDistros)
    fmt.Println("done")
}

func maxOfList(numbers []int) (int, int){
  max := numbers[0]
  indexOfMax := 0
  for num := range numbers {
    if(max<numbers[num]){
      max = numbers[num]
      indexOfMax = num
    }
  }
  return max,indexOfMax
}

func checkIfRepeat(list []string, element string) bool {
  for item := range list {
    if (list[item] == element) {
      return true
    }
  }
  return false
}
