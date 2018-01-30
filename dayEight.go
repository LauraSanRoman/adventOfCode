package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
)

func main () {
  f, _ := os.Open("/Users/Laura/Documents/adventOfCode/input/day8")
  scanner := bufio.NewScanner(f)
  var listOfInstructions [][]string
  for scanner.Scan(){
    line := scanner.Text()
    listOfInstructions = append(listOfInstructions,strings.Split(line," "))
  }
  registers := make(map[string]int)

  for instruction := range listOfInstructions {
    registers[listOfInstructions[instruction][0]] = 0
  }
  max := -9999999
  temp := 0
  for instruction := range listOfInstructions {
     // 0 = register
     // 1 = inc / dec
     // 2 = amount
     // 3 = if
     // 4 = register
     // 5 = operator
     // 6 = value
     registerValue := registers[listOfInstructions[instruction][4]]
     givenValue, _ := strconv.Atoi(listOfInstructions[instruction][6])
     amount, _ := strconv.Atoi(listOfInstructions[instruction][2])
    if(compare(registerValue, givenValue, listOfInstructions[instruction][5])){
      if(listOfInstructions[instruction][1] == "inc"){
        registers[listOfInstructions[instruction][0]] += amount
      }
      if(listOfInstructions[instruction][1] == "dec"){
        registers[listOfInstructions[instruction][0]] -= amount
      }
    }
    temp = getGreatest(registers)
    if(temp>max){
      max = temp
    }
  }
  fmt.Println(getGreatest(registers))
  fmt.Println(max)

}

func compare (value1 int, value2 int, op string) bool{
  switch op {
  case "==":
    return value1 == value2
  case "<=":
    return value1 <= value2
  case ">=":
    return value1 >= value2
  case ">":
    return value1 > value2
  case "<":
    return value1 < value2
  case "!=":
    return value1 != value2
  default:
    return false

  }
}

func getGreatest(regs map[string]int) int {
  max := -9999999
  for entry  := range regs {
    if(regs[entry] > max){
      max = regs[entry]
    }
  }
  return max
}
