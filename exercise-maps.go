package main
import(
    "fmt"
    "strings"
)
func WordCount(s string) map[string]int{
    arr:=strings.Fields(s)
    result:=make(map[string]int)
    for _,w:=range(arr){
        result[w]++
    }
    return result
}

func main(){
    testString:="how are you fine thank you and you"
    fmt.Println(WordCount(testString))
}
