package main
import (
    "fmt"
    "strings"
)
type MyReader struct{}

func (mr MyReader) Read(b []byte)(int ,error){
    r:=strings.NewReader("A")
    return r.Read(b)
}
