package main

import(
  "fmt"
  "os"
  "bufio"
)

func main(){
  distro := []int {5,1,10,0,1,7,13,14,3,12,8,10,7,12,0,6}
  f, _ := os.Create("/Users/hpw355/Documents/AdventOfCode/input/day6")
  w := bufio.NewWriter(f)
  listOfDistros := [][]int {distro}
  fmt.Println(distro)
  for i := 1; i < 100; i++ {
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
    fmt.Println("i: ",i,"distro:", distro)
    _,_ = fmt.Fprintf(w,"%v\n",distro)
    listOfDistros = Extend(listOfDistros,distro)
  }
  fmt.Println(listOfDistros)
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
func Extend(slice [][]int, element []int) [][]int {
    n := len(slice)
    if n == cap(slice) {
        // Slice is full; must grow.
        // We double its size and add 1, so if the size is zero we still grow.
        newSlice := make([][]int, len(slice), 2*len(slice)+1)
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0 : n+1]
    slice[n] = element
    return slice
}
