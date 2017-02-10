package main

import(
  "fmt"
  "io/ioutil"
  "Sequence-analysis/bio"
  "os"
)

func check(e error){
  if e != nil {
    panic(e)
    }
}

func main() {
  //f, err := os.Open("test.txt")
  args := os.Args[1:]
  dat, err := ioutil.ReadFile(args[1])
  check(err)

  switch method := args[0] ; method {
  case "GCContent":
    fmt.Println("GC-Gehalt:", bio.GCContent(string(dat)))
  case "CountNuc":
    bio.CountNuc(string(dat))
    fmt.Println("A:", bio.A)
    fmt.Println("T:", bio.T)
    fmt.Println("G:", bio.G)
    fmt.Println("C:", bio.C)
  default:
    fmt.Println("Keine Methode angegeben")

  }

  



}
