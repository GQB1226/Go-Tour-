package main
import (
    "log"
    "net/http"
)

type String string

type Struct struct{
    Greeting string
    Punct    string
    Who      string
}

func (s String) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request){
        w.Write([]byte(s))
    }
func (s Struct) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request){
        w.Write([]byte(s.Greeting+s.Punct+s.Who))
    }


func main(){
    http.Handle("/string",String("I'm a dfa"))
    http.Handle("/struct",&Struct{"Hello",":","GQB1226!"})
    log.Fatal(http.ListenAndServe("localhost:4000",nil))
}
