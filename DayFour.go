package main

import (
	"bufio"
  "fmt"
  "os"
	"strings"
)
func main (){
	f, _ := os.Open("/Users/hpw355/Documents/AdventOfCode/input/day4")
	scanner := bufio.NewScanner(f)
	count1 := 0
	count2 := 0
	for scanner.Scan(){
		line := scanner.Text()
		if(checkPassPhrasePartOne(line)){
			count1++
		}
		if(checkPassPhrasePartTwo(line)){
			count2++
		}
	}

	fmt.Println("count 1: ", count1)
	fmt.Println("count 2: ", count2)
}

func checkPassPhrasePartOne(passPhrase string) bool{
		parts := strings.Fields(passPhrase)
		for i := 0;i< len(parts);i++ {
			for j := len(parts)-1; j > i; j-- {
				if(parts[i] == parts[j]){
					return false
				}
			}
		}
		return true
}
func checkPassPhrasePartTwo(passPhrase string) bool{
		parts := strings.Fields(passPhrase)
		for i := 0;i< len(parts);i++ {
			for j := len(parts)-1; j > i; j-- {
				if(len(parts[i]) == len(parts[j])){
					strOne := strings.SplitAfter(parts[i],"")
					strTwo := strings.SplitAfter(parts[j],"")
					if(matchingLetters(strOne, strTwo)){
						return false
					}
				}
			}
		}
		return true
}

func matchingLetters(one []string, two []string) bool{
	matching := 0
		for x := range one {
			for y := range two {
				if(one[x] == two[y]){
					matching++
					one[x] = "0"
					two[y] = "1"
				}
			}
		}
		return matching == len(one)
}
