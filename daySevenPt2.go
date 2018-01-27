package main
import(
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
)
type DISC struct {
  weight int
  connections []string
  balanced bool
}

func main (){
  f, _ := os.Open("/Users/Laura/Documents/adventOfCode/input/day7")
  scanner := bufio.NewScanner(f)
  m := make(map[string]DISC)
  var listOfConnections []string
  weightMap := make(map[string]int)

  for scanner.Scan(){
    line := scanner.Text()
    indexOfStartParen := strings.Index(line,"(")
    indexOfEndParen := strings.Index(line,")")
    indexOfArrow := strings.Index(line, ">")
    weight, _ := strconv.Atoi(strings.Trim(line[indexOfStartParen+1:indexOfEndParen]," "))
    if(strings.ContainsAny(line,"->")){
      listOfConnections = strings.Split(line[indexOfArrow+1:],", ")
      for str  := range listOfConnections {
        listOfConnections[str] = strings.Trim(listOfConnections[str]," ")
      }
      m[strings.Trim(line[0:indexOfStartParen]," ")] = DISC{weight,listOfConnections,true}
    }else{
      m[strings.Trim(line[0:indexOfStartParen]," ")] = DISC{weight, []string{},true}
    }
  }
  weightMap = convertWeights(m)
  for disc  := range m {
    for i := 0; i < len(m[disc].connections); i++ {
      if(weightMap[m[disc].connections[0]] != weightMap[m[disc].connections[i]]){
        m[disc] = DISC{m[disc].weight,m[disc].connections, false}
      }
    }
    if(!m[disc].balanced){
      fmt.Println("key: ", disc, "value", m[disc])
      for item  := range m[disc].connections {
        fmt.Println("key: ", m[disc].connections[item], "weight", weightMap[m[disc].connections[item]])
      }
    }
  }
}

func convertWeights(myMap map[string]DISC) map[string]int{
  weightMap := make(map[string]int)
  for item  := range myMap {
    sum := 0
    for i  := 0; i  < len(myMap[item].connections); i++ {
      listItem := myMap[item].connections[i]
      sum = sum + myMap[listItem].weight
    }
    weightMap[item] = sum + myMap[item].weight
  }
  return weightMap
}
